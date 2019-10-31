// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/kyma-project/helm-broker/pkg/apis/addons/v1alpha1"

// clusterAddonsCfgMutations is an autogenerated mock type for the clusterAddonsCfgMutations type
type clusterAddonsCfgMutations struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, repository, labels
func (_m *clusterAddonsCfgMutations) Create(name string, repository []gqlschema.AddonsConfigurationRepositoryInput, labels *gqlschema.Labels) (*v1alpha1.ClusterAddonsConfiguration, error) {
	ret := _m.Called(name, repository, labels)

	var r0 *v1alpha1.ClusterAddonsConfiguration
	if rf, ok := ret.Get(0).(func(string, []gqlschema.AddonsConfigurationRepositoryInput, *gqlschema.Labels) *v1alpha1.ClusterAddonsConfiguration); ok {
		r0 = rf(name, repository, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ClusterAddonsConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []gqlschema.AddonsConfigurationRepositoryInput, *gqlschema.Labels) error); ok {
		r1 = rf(name, repository, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *clusterAddonsCfgMutations) Delete(name string) (*v1alpha1.ClusterAddonsConfiguration, error) {
	ret := _m.Called(name)

	var r0 *v1alpha1.ClusterAddonsConfiguration
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.ClusterAddonsConfiguration); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ClusterAddonsConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: name, repository, labels
func (_m *clusterAddonsCfgMutations) Update(name string, repository []gqlschema.AddonsConfigurationRepositoryInput, labels *gqlschema.Labels) (*v1alpha1.ClusterAddonsConfiguration, error) {
	ret := _m.Called(name, repository, labels)

	var r0 *v1alpha1.ClusterAddonsConfiguration
	if rf, ok := ret.Get(0).(func(string, []gqlschema.AddonsConfigurationRepositoryInput, *gqlschema.Labels) *v1alpha1.ClusterAddonsConfiguration); ok {
		r0 = rf(name, repository, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ClusterAddonsConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []gqlschema.AddonsConfigurationRepositoryInput, *gqlschema.Labels) error); ok {
		r1 = rf(name, repository, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
