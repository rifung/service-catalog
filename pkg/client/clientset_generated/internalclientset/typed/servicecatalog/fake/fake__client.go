/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/typed/servicecatalog"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeServicecatalog struct {
	*core.Fake
}

func (c *FakeServicecatalog) Bindings(namespace string) servicecatalog.BindingInterface {
	return &FakeBindings{c, namespace}
}

func (c *FakeServicecatalog) Brokers() servicecatalog.BrokerInterface {
	return &FakeBrokers{c}
}

func (c *FakeServicecatalog) Instances(namespace string) servicecatalog.InstanceInterface {
	return &FakeInstances{c, namespace}
}

func (c *FakeServicecatalog) ServiceClasses(namespace string) servicecatalog.ServiceClassInterface {
	return &FakeServiceClasses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServicecatalog) RESTClient() restclient.Interface {
	var ret *restclient.RESTClient
	return ret
}