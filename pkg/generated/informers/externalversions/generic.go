//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// This file was automatically generated by informer-gen

package externalversions

import (
	"fmt"

	v1alpha "github.com/arangodb/k8s-operator/pkg/apis/arangodb/v1alpha"
	storage_v1alpha "github.com/arangodb/k8s-operator/pkg/apis/storage/v1alpha"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=database.arangodb.com, Version=v1alpha
	case v1alpha.SchemeGroupVersion.WithResource("arangodeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Database().V1alpha().ArangoDeployments().Informer()}, nil

		// Group=storage.arangodb.com, Version=v1alpha
	case storage_v1alpha.SchemeGroupVersion.WithResource("arangolocalstorages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha().ArangoLocalStorages().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
