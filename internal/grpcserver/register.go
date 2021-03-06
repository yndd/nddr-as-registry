/*
Copyright 2021 NDDO.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpcserver

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/yndd/nddo-grpc/resource/resourcepb"
	asv1alpha1 "github.com/yndd/nddr-as-registry/apis/as/v1alpha1"
	"github.com/yndd/nddr-as-registry/internal/handler"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func (r *server) ResourceGet(ctx context.Context, req *resourcepb.Request) (*resourcepb.Reply, error) {
	log := r.log.WithValues("Request", req)
	log.Debug("ResourceGet...")

	return &resourcepb.Reply{Ready: true}, nil
}

func (r *server) ResourceRequest(ctx context.Context, req *resourcepb.Request) (*resourcepb.Reply, error) {
	log := r.log.WithValues("Request", req)

	registerInfo := &handler.RegisterInfo{
		Namespace:    req.GetNamespace(),
		RegistryName: req.GetRegistryName(),
		Name:         req.GetName(),
		CrName:       strings.Join([]string{req.GetNamespace(), req.GetRegistryName()}, "."),
		Selector:     req.Request.Selector,
		SourceTag:    req.Request.SourceTag,
	}

	log.Debug("resource alloc", "registerInfo", registerInfo)

	index, err := r.handler.Register(ctx, registerInfo)
	if err != nil {
		return &resourcepb.Reply{Ready: false}, err
	}

	// send a generic event to trigger a registry reconciliation based on a new allocation
	// to update the status
	r.eventChs[asv1alpha1.RegistryGroupKind] <- event.GenericEvent{
		Object: &asv1alpha1.Register{
			ObjectMeta: metav1.ObjectMeta{Name: req.GetName(), Namespace: req.GetNamespace()},
		},
	}

	return &resourcepb.Reply{
		Ready:      true,
		Timestamp:  time.Now().UnixNano(),
		ExpiryTime: time.Now().UnixNano(),
		Data: map[string]*resourcepb.TypedValue{
			"as": {Value: &resourcepb.TypedValue_StringVal{StringVal: strconv.Itoa(int(*index))}},
		},
	}, nil
}

func (r *server) ResourceRelease(ctx context.Context, req *resourcepb.Request) (*resourcepb.Reply, error) {
	log := r.log.WithValues("Request", req)
	log.Debug("ResourceDeAlloc...")

	registerInfo := &handler.RegisterInfo{
		Namespace:    req.GetNamespace(),
		RegistryName: req.GetRegistryName(),
		CrName:       strings.Join([]string{req.GetNamespace(), req.GetRegistryName()}, "."),
		Name:         req.GetName(),
		Selector:     req.Request.Selector,
		SourceTag:    req.Request.SourceTag,
	}

	log.Debug("resource dealloc", "registerInfo", registerInfo)

	if err := r.handler.DeRegister(ctx, registerInfo); err != nil {
		return &resourcepb.Reply{Ready: false}, err
	}

	// send a generic event to trigger a registry reconciliation based on a new DeAllocation
	r.eventChs[asv1alpha1.RegistryGroupKind] <- event.GenericEvent{
		Object: &asv1alpha1.Register{
			ObjectMeta: metav1.ObjectMeta{Name: req.GetName(), Namespace: req.GetNamespace()},
		},
	}

	return &resourcepb.Reply{Ready: true}, nil
}
