/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CephFilesystemMirrorsGetter has a method to return a CephFilesystemMirrorInterface.
// A group's client should implement this interface.
type CephFilesystemMirrorsGetter interface {
	CephFilesystemMirrors(namespace string) CephFilesystemMirrorInterface
}

// CephFilesystemMirrorInterface has methods to work with CephFilesystemMirror resources.
type CephFilesystemMirrorInterface interface {
	Create(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.CreateOptions) (*v1.CephFilesystemMirror, error)
	Update(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.UpdateOptions) (*v1.CephFilesystemMirror, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CephFilesystemMirror, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CephFilesystemMirrorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephFilesystemMirror, err error)
	CephFilesystemMirrorExpansion
}

// cephFilesystemMirrors implements CephFilesystemMirrorInterface
type cephFilesystemMirrors struct {
	client rest.Interface
	ns     string
}

// newCephFilesystemMirrors returns a CephFilesystemMirrors
func newCephFilesystemMirrors(c *CephV1Client, namespace string) *cephFilesystemMirrors {
	return &cephFilesystemMirrors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cephFilesystemMirror, and returns the corresponding cephFilesystemMirror object, and an error if there is any.
func (c *cephFilesystemMirrors) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephFilesystemMirror, err error) {
	result = &v1.CephFilesystemMirror{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CephFilesystemMirrors that match those selectors.
func (c *cephFilesystemMirrors) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephFilesystemMirrorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CephFilesystemMirrorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cephFilesystemMirrors.
func (c *cephFilesystemMirrors) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cephFilesystemMirror and creates it.  Returns the server's representation of the cephFilesystemMirror, and an error, if there is any.
func (c *cephFilesystemMirrors) Create(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.CreateOptions) (result *v1.CephFilesystemMirror, err error) {
	result = &v1.CephFilesystemMirror{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephFilesystemMirror).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cephFilesystemMirror and updates it. Returns the server's representation of the cephFilesystemMirror, and an error, if there is any.
func (c *cephFilesystemMirrors) Update(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.UpdateOptions) (result *v1.CephFilesystemMirror, err error) {
	result = &v1.CephFilesystemMirror{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		Name(cephFilesystemMirror.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephFilesystemMirror).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cephFilesystemMirror and deletes it. Returns an error if one occurs.
func (c *cephFilesystemMirrors) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cephFilesystemMirrors) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cephFilesystemMirror.
func (c *cephFilesystemMirrors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephFilesystemMirror, err error) {
	result = &v1.CephFilesystemMirror{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cephfilesystemmirrors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
