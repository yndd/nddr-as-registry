/*
Copyright 2021 NDD.

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

package registry

import (
	"context"

	//ndddvrv1 "github.com/yndd/ndd-core/apis/dvr/v1"
	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/nddo-runtime/pkg/odns"
	asv1alpha1 "github.com/yndd/nddr-as-registry/apis/as/v1alpha1"
	"github.com/yndd/nddr-as-registry/internal/handler"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type adder interface {
	Add(item interface{})
}

type EnqueueRequestForAllRegisters struct {
	client client.Client
	log    logging.Logger
	ctx    context.Context

	handler handler.Handler
	//mutex  sync.Mutex

	newRegistryList func() asv1alpha1.RgList
}

// Create enqueues a request for all infrastructures which pertains to the topology.
func (e *EnqueueRequestForAllRegisters) Create(evt event.CreateEvent, q workqueue.RateLimitingInterface) {
	e.add(evt.Object, q)
}

// Create enqueues a request for all infrastructures which pertains to the topology.
func (e *EnqueueRequestForAllRegisters) Update(evt event.UpdateEvent, q workqueue.RateLimitingInterface) {
	e.add(evt.ObjectOld, q)
	e.add(evt.ObjectNew, q)
}

// Create enqueues a request for all infrastructures which pertains to the topology.
func (e *EnqueueRequestForAllRegisters) Delete(evt event.DeleteEvent, q workqueue.RateLimitingInterface) {
	e.add(evt.Object, q)
}

// Create enqueues a request for all infrastructures which pertains to the topology.
func (e *EnqueueRequestForAllRegisters) Generic(evt event.GenericEvent, q workqueue.RateLimitingInterface) {
	e.add(evt.Object, q)
}

func (e *EnqueueRequestForAllRegisters) add(obj runtime.Object, queue adder) {
	dd, ok := obj.(*asv1alpha1.Register)
	if !ok {
		return
	}
	log := e.log.WithValues("function", "watch register", "name", dd.GetName())
	log.Debug("register handleEvent")

	d := e.newRegistryList()
	if err := e.client.List(e.ctx, d); err != nil {
		return
	}

	registryName := odns.Name2OdnsRegistry(dd.GetName()).GetRegistryName()

	for _, registry := range d.GetRegistries() {
		// only enqueue if the org and/or deployment name match
		if registry.GetName() == registryName {
			crName := getCrName(registry)
			e.handler.ResetSpeedy(crName)

			queue.Add(reconcile.Request{NamespacedName: types.NamespacedName{
				Namespace: registry.GetNamespace(),
				Name:      registry.GetName()}})
		}
	}
}
