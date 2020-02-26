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

package fake

import (
	rpccontrolleristioiov1 "istio.io/istio/pkg/api/rpccontroller.istio.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRpcServices implements RpcServiceInterface
type FakeRpcServices struct {
	Fake *FakeRpccontrollerV1
	ns   string
}

var rpcservicesResource = schema.GroupVersionResource{Group: "rpccontroller.istio.io", Version: "v1", Resource: "rpcservices"}

var rpcservicesKind = schema.GroupVersionKind{Group: "rpccontroller.istio.io", Version: "v1", Kind: "RpcService"}

// Get takes name of the rpcService, and returns the corresponding rpcService object, and an error if there is any.
func (c *FakeRpcServices) Get(name string, options v1.GetOptions) (result *rpccontrolleristioiov1.RpcService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rpcservicesResource, c.ns, name), &rpccontrolleristioiov1.RpcService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rpccontrolleristioiov1.RpcService), err
}

// List takes label and field selectors, and returns the list of RpcServices that match those selectors.
func (c *FakeRpcServices) List(opts v1.ListOptions) (result *rpccontrolleristioiov1.RpcServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rpcservicesResource, rpcservicesKind, c.ns, opts), &rpccontrolleristioiov1.RpcServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rpccontrolleristioiov1.RpcServiceList{ListMeta: obj.(*rpccontrolleristioiov1.RpcServiceList).ListMeta}
	for _, item := range obj.(*rpccontrolleristioiov1.RpcServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rpcServices.
func (c *FakeRpcServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rpcservicesResource, c.ns, opts))

}

// Create takes the representation of a rpcService and creates it.  Returns the server's representation of the rpcService, and an error, if there is any.
func (c *FakeRpcServices) Create(rpcService *rpccontrolleristioiov1.RpcService) (result *rpccontrolleristioiov1.RpcService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rpcservicesResource, c.ns, rpcService), &rpccontrolleristioiov1.RpcService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rpccontrolleristioiov1.RpcService), err
}

// Update takes the representation of a rpcService and updates it. Returns the server's representation of the rpcService, and an error, if there is any.
func (c *FakeRpcServices) Update(rpcService *rpccontrolleristioiov1.RpcService) (result *rpccontrolleristioiov1.RpcService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rpcservicesResource, c.ns, rpcService), &rpccontrolleristioiov1.RpcService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rpccontrolleristioiov1.RpcService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRpcServices) UpdateStatus(rpcService *rpccontrolleristioiov1.RpcService) (*rpccontrolleristioiov1.RpcService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rpcservicesResource, "status", c.ns, rpcService), &rpccontrolleristioiov1.RpcService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rpccontrolleristioiov1.RpcService), err
}

// Delete takes name of the rpcService and deletes it. Returns an error if one occurs.
func (c *FakeRpcServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rpcservicesResource, c.ns, name), &rpccontrolleristioiov1.RpcService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRpcServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rpcservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &rpccontrolleristioiov1.RpcServiceList{})
	return err
}

// Patch applies the patch and returns the patched rpcService.
func (c *FakeRpcServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rpccontrolleristioiov1.RpcService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rpcservicesResource, c.ns, name, data, subresources...), &rpccontrolleristioiov1.RpcService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rpccontrolleristioiov1.RpcService), err
}