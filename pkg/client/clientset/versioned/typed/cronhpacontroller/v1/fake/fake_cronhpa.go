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
	cronhpacontrollerv1 "tkestack.io/cron-hpa/pkg/apis/cronhpacontroller/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCronHPAs implements CronHPAInterface
type FakeCronHPAs struct {
	Fake *FakeCronhpacontrollerV1
	ns   string
}

var cronhpasResource = schema.GroupVersionResource{Group: "cronhpacontroller.extensions.tkestack.io", Version: "v1", Resource: "cronhpas"}

var cronhpasKind = schema.GroupVersionKind{Group: "cronhpacontroller.extensions.tkestack.io", Version: "v1", Kind: "CronHPA"}

// Get takes name of the cronHPA, and returns the corresponding cronHPA object, and an error if there is any.
func (c *FakeCronHPAs) Get(name string, options v1.GetOptions) (result *cronhpacontrollerv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cronhpasResource, c.ns, name), &cronhpacontrollerv1.CronHPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cronhpacontrollerv1.CronHPA), err
}

// List takes label and field selectors, and returns the list of CronHPAs that match those selectors.
func (c *FakeCronHPAs) List(opts v1.ListOptions) (result *cronhpacontrollerv1.CronHPAList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cronhpasResource, cronhpasKind, c.ns, opts), &cronhpacontrollerv1.CronHPAList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cronhpacontrollerv1.CronHPAList{ListMeta: obj.(*cronhpacontrollerv1.CronHPAList).ListMeta}
	for _, item := range obj.(*cronhpacontrollerv1.CronHPAList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronHPAs.
func (c *FakeCronHPAs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cronhpasResource, c.ns, opts))

}

// Create takes the representation of a cronHPA and creates it.  Returns the server's representation of the cronHPA, and an error, if there is any.
func (c *FakeCronHPAs) Create(cronHPA *cronhpacontrollerv1.CronHPA) (result *cronhpacontrollerv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cronhpasResource, c.ns, cronHPA), &cronhpacontrollerv1.CronHPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cronhpacontrollerv1.CronHPA), err
}

// Update takes the representation of a cronHPA and updates it. Returns the server's representation of the cronHPA, and an error, if there is any.
func (c *FakeCronHPAs) Update(cronHPA *cronhpacontrollerv1.CronHPA) (result *cronhpacontrollerv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cronhpasResource, c.ns, cronHPA), &cronhpacontrollerv1.CronHPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cronhpacontrollerv1.CronHPA), err
}

// Delete takes name of the cronHPA and deletes it. Returns an error if one occurs.
func (c *FakeCronHPAs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cronhpasResource, c.ns, name), &cronhpacontrollerv1.CronHPA{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCronHPAs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cronhpasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cronhpacontrollerv1.CronHPAList{})
	return err
}

// Patch applies the patch and returns the patched cronHPA.
func (c *FakeCronHPAs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cronhpacontrollerv1.CronHPA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cronhpasResource, c.ns, name, pt, data, subresources...), &cronhpacontrollerv1.CronHPA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cronhpacontrollerv1.CronHPA), err
}