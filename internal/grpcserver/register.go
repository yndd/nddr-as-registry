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
	"github.com/yndd/nddo-runtime/pkg/odr"
	asregv1alpha1 "github.com/yndd/nddr-as-registry/apis/registry/v1alpha1"
	"github.com/yndd/nddr-as-registry/internal/handler"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func (r *server) ResourceGet(ctx context.Context, req *resourcepb.Request) (*resourcepb.Reply, error) {
	log := r.log.WithValues("Request", req)
	log.Debug("ResourceGet...")

	return &resourcepb.Reply{Ready: true}, nil
}

func (r *server) ResourceAlloc(ctx context.Context, req *resourcepb.Request) (*resourcepb.Reply, error) {
	log := r.log.WithValues("Request", req)

	namespace := req.GetNamespace()
	// the resourceName contains org.poolname.nodename to be consistent with the k8s alloc approach
	// where the name needs to be unique. As such we have to cut the last element of the string seperated
	// by dots to get the asPoolName
	// this approach is consistent when you use as pool per deployment or per organization
	//registryName := strings.Join(strings.Split(req.GetResourceName(), ".")[:len(strings.Split(req.GetResourceName(), "."))-1], ".")
	//registryName := strings.Join([]string{getOrganizationName(req.ResourceName), getRegistryName(req.ResourceName)}, ".")
	odr, err := odr.GetOdrRegisterInfo(req.ResourceName)
	if err != nil {
		return nil, err
	}

	registerInfo := &handler.RegisterInfo{
		Namespace:    req.GetNamespace(),
		RegistryName: odr.FullRegistryName,
		CrName:       strings.Join([]string{namespace, odr.FullRegistryName}, "."),
		Selector:     req.Alloc.Selector,
		SourceTag:    req.Alloc.SourceTag,
	}

	log.Debug("resource alloc", "registerInfo", registerInfo)

	index, err := r.handler.Register(ctx, registerInfo)
	if err != nil {
		return &resourcepb.Reply{Ready: false}, err
	}

	// send a generic event to trigger a registry reconciliation based on a new allocation
	r.eventChs[asregv1alpha1.RegistryGroupKind] <- event.GenericEvent{
		Object: &asregv1alpha1.Register{
			ObjectMeta: metav1.ObjectMeta{Name: req.ResourceName, Namespace: namespace},
		},
	}

	return &resourcepb.Reply{
		Ready:      true,
		Timestamp:  time.Now().UnixNano(),
		ExpiryTime: time.Now().UnixNano(),
		Data: map[string]*resourcepb.TypedValue{
			"index": {Value: &resourcepb.TypedValue_StringVal{StringVal: strconv.Itoa(int(*index))}},
		},
	}, nil
}

func (r *server) ResourceDeAlloc(ctx context.Context, req *resourcepb.Request) (*resourcepb.Reply, error) {
	log := r.log.WithValues("Request", req)
	log.Debug("ResourceDeAlloc...")

	namespace := req.GetNamespace()
	// the resourceName contains org.poolname.nodename to be consistent with the k8s alloc approach
	// where the name needs to be unique. As such we have to cut the last element of the string seperated
	// by dots to get the asPoolName
	// this approach is consistent when you use as pool per deployment or per organization
	//registryName := strings.Join(strings.Split(req.GetResourceName(), ".")[:len(strings.Split(req.GetResourceName(), "."))-1], ".")
	//registryName := strings.Join([]string{getOrganizationName(req.ResourceName), getRegistryName(req.ResourceName)}, ".")
	odr, err := odr.GetOdrRegisterInfo(req.ResourceName)
	if err != nil {
		return nil, err
	}

	registerInfo := &handler.RegisterInfo{
		Namespace:    req.GetNamespace(),
		RegistryName: odr.FullRegistryName,
		CrName:       strings.Join([]string{namespace, odr.FullRegistryName}, "."),
		Selector:     req.Alloc.Selector,
		SourceTag:    req.Alloc.SourceTag,
	}

	log.Debug("resource dealloc", "registerInfo", registerInfo)

	if err := r.handler.DeRegister(ctx, registerInfo); err != nil {
		return &resourcepb.Reply{Ready: false}, err
	}

	// send a generic event to trigger a registry reconciliation based on a new DeAllocation
	r.eventChs[asregv1alpha1.RegistryGroupKind] <- event.GenericEvent{
		Object: &asregv1alpha1.Register{
			ObjectMeta: metav1.ObjectMeta{Name: req.ResourceName, Namespace: namespace},
		},
	}

	return &resourcepb.Reply{Ready: true}, nil
}
