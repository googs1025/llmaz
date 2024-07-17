/*
Copyright 2024.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "inftyai.com/llmaz/api/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ModelLister helps list Models.
// All objects returned here must be treated as read-only.
type ModelLister interface {
	// List lists all Models in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Model, err error)
	// Models returns an object that can list and get Models.
	Models(namespace string) ModelNamespaceLister
	ModelListerExpansion
}

// modelLister implements the ModelLister interface.
type modelLister struct {
	indexer cache.Indexer
}

// NewModelLister returns a new ModelLister.
func NewModelLister(indexer cache.Indexer) ModelLister {
	return &modelLister{indexer: indexer}
}

// List lists all Models in the indexer.
func (s *modelLister) List(selector labels.Selector) (ret []*v1alpha1.Model, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Model))
	})
	return ret, err
}

// Models returns an object that can list and get Models.
func (s *modelLister) Models(namespace string) ModelNamespaceLister {
	return modelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ModelNamespaceLister helps list and get Models.
// All objects returned here must be treated as read-only.
type ModelNamespaceLister interface {
	// List lists all Models in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Model, err error)
	// Get retrieves the Model from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Model, error)
	ModelNamespaceListerExpansion
}

// modelNamespaceLister implements the ModelNamespaceLister
// interface.
type modelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Models in the indexer for a given namespace.
func (s modelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Model, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Model))
	})
	return ret, err
}

// Get retrieves the Model from the indexer for a given namespace and name.
func (s modelNamespaceLister) Get(name string) (*v1alpha1.Model, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("model"), name)
	}
	return obj.(*v1alpha1.Model), nil
}
