/*
Copyright 2020 The KubeSphere Authors.

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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubesphere.io/api/virtualization/v1alpha1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// DiskVolumesGetter has a method to return a DiskVolumeInterface.
// A group's client should implement this interface.
type DiskVolumesGetter interface {
	DiskVolumes(namespace string) DiskVolumeInterface
}

// DiskVolumeInterface has methods to work with DiskVolume resources.
type DiskVolumeInterface interface {
	Create(ctx context.Context, diskVolume *v1alpha1.DiskVolume, opts v1.CreateOptions) (*v1alpha1.DiskVolume, error)
	Update(ctx context.Context, diskVolume *v1alpha1.DiskVolume, opts v1.UpdateOptions) (*v1alpha1.DiskVolume, error)
	UpdateStatus(ctx context.Context, diskVolume *v1alpha1.DiskVolume, opts v1.UpdateOptions) (*v1alpha1.DiskVolume, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DiskVolume, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DiskVolumeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DiskVolume, err error)
	DiskVolumeExpansion
}

// diskVolumes implements DiskVolumeInterface
type diskVolumes struct {
	client rest.Interface
	ns     string
}

// newDiskVolumes returns a DiskVolumes
func newDiskVolumes(c *VirtualizationV1alpha1Client, namespace string) *diskVolumes {
	return &diskVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the diskVolume, and returns the corresponding diskVolume object, and an error if there is any.
func (c *diskVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DiskVolume, err error) {
	result = &v1alpha1.DiskVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("diskvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DiskVolumes that match those selectors.
func (c *diskVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DiskVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DiskVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("diskvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested diskVolumes.
func (c *diskVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("diskvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a diskVolume and creates it.  Returns the server's representation of the diskVolume, and an error, if there is any.
func (c *diskVolumes) Create(ctx context.Context, diskVolume *v1alpha1.DiskVolume, opts v1.CreateOptions) (result *v1alpha1.DiskVolume, err error) {
	result = &v1alpha1.DiskVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("diskvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(diskVolume).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a diskVolume and updates it. Returns the server's representation of the diskVolume, and an error, if there is any.
func (c *diskVolumes) Update(ctx context.Context, diskVolume *v1alpha1.DiskVolume, opts v1.UpdateOptions) (result *v1alpha1.DiskVolume, err error) {
	result = &v1alpha1.DiskVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("diskvolumes").
		Name(diskVolume.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(diskVolume).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *diskVolumes) UpdateStatus(ctx context.Context, diskVolume *v1alpha1.DiskVolume, opts v1.UpdateOptions) (result *v1alpha1.DiskVolume, err error) {
	result = &v1alpha1.DiskVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("diskvolumes").
		Name(diskVolume.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(diskVolume).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the diskVolume and deletes it. Returns an error if one occurs.
func (c *diskVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("diskvolumes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *diskVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("diskvolumes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched diskVolume.
func (c *diskVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DiskVolume, err error) {
	result = &v1alpha1.DiskVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("diskvolumes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
