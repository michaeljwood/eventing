/*
Copyright 2018 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	scheme "github.com/knative/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TargetsGetter has a method to return a TargetInterface.
// A group's client should implement this interface.
type TargetsGetter interface {
	Targets(namespace string) TargetInterface
}

// TargetInterface has methods to work with Target resources.
type TargetInterface interface {
	Create(*v1alpha1.Target) (*v1alpha1.Target, error)
	Update(*v1alpha1.Target) (*v1alpha1.Target, error)
	UpdateStatus(*v1alpha1.Target) (*v1alpha1.Target, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Target, error)
	List(opts v1.ListOptions) (*v1alpha1.TargetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Target, err error)
	TargetExpansion
}

// targets implements TargetInterface
type targets struct {
	client rest.Interface
	ns     string
}

// newTargets returns a Targets
func newTargets(c *DuckV1alpha1Client, namespace string) *targets {
	return &targets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the target, and returns the corresponding target object, and an error if there is any.
func (c *targets) Get(name string, options v1.GetOptions) (result *v1alpha1.Target, err error) {
	result = &v1alpha1.Target{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("targets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Targets that match those selectors.
func (c *targets) List(opts v1.ListOptions) (result *v1alpha1.TargetList, err error) {
	result = &v1alpha1.TargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("targets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested targets.
func (c *targets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("targets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a target and creates it.  Returns the server's representation of the target, and an error, if there is any.
func (c *targets) Create(target *v1alpha1.Target) (result *v1alpha1.Target, err error) {
	result = &v1alpha1.Target{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("targets").
		Body(target).
		Do().
		Into(result)
	return
}

// Update takes the representation of a target and updates it. Returns the server's representation of the target, and an error, if there is any.
func (c *targets) Update(target *v1alpha1.Target) (result *v1alpha1.Target, err error) {
	result = &v1alpha1.Target{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("targets").
		Name(target.Name).
		Body(target).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *targets) UpdateStatus(target *v1alpha1.Target) (result *v1alpha1.Target, err error) {
	result = &v1alpha1.Target{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("targets").
		Name(target.Name).
		SubResource("status").
		Body(target).
		Do().
		Into(result)
	return
}

// Delete takes name of the target and deletes it. Returns an error if one occurs.
func (c *targets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("targets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *targets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("targets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched target.
func (c *targets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Target, err error) {
	result = &v1alpha1.Target{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("targets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
