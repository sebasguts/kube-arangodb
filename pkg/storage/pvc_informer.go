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
// Author Ewout Prangsma
//

package storage

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

// listenForPvcEvents keep listening for changes in PVC's until the given channel is closed.
func (ls *LocalStorage) listenForPvcEvents() {
	source := cache.NewListWatchFromClient(
		ls.deps.KubeCli.CoreV1().RESTClient(),
		"persistentvolumeclaims",
		ls.apiObject.GetNamespace(),
		fields.Everything())

	getPvc := func(obj interface{}) (*v1.PersistentVolumeClaim, bool) {
		pvc, ok := obj.(*v1.PersistentVolumeClaim)
		if !ok {
			tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
			if !ok {
				return nil, false
			}
			pvc, ok = tombstone.Obj.(*v1.PersistentVolumeClaim)
			return pvc, ok
		}
		return pvc, true
	}

	_, informer := cache.NewIndexerInformer(source, &v1.PersistentVolumeClaim{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if pvc, ok := getPvc(obj); ok {
				ls.send(&localStorageEvent{
					Type: eventPVCAdded,
					PersistentVolumeClaim: pvc,
				})
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			if pvc, ok := getPvc(newObj); ok {
				ls.send(&localStorageEvent{
					Type: eventPVCUpdated,
					PersistentVolumeClaim: pvc,
				})
			}
		},
		DeleteFunc: func(obj interface{}) {
			// Ignore
		},
	}, cache.Indexers{})

	informer.Run(ls.stopCh)
}
