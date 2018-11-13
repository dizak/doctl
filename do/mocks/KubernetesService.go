// Code generated by mockery v1.0.0. DO NOT EDIT.

// Generated: please do not edit by hand

package mocks

import do "github.com/digitalocean/doctl/do"
import godo "github.com/digitalocean/godo"
import mock "github.com/stretchr/testify/mock"

// KubernetesService is an autogenerated mock type for the KubernetesService type
type KubernetesService struct {
	mock.Mock
}

// Create provides a mock function with given fields: create
func (_m *KubernetesService) Create(create *godo.KubernetesClusterCreateRequest) (*do.KubernetesCluster, error) {
	ret := _m.Called(create)

	var r0 *do.KubernetesCluster
	if rf, ok := ret.Get(0).(func(*godo.KubernetesClusterCreateRequest) *do.KubernetesCluster); ok {
		r0 = rf(create)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.KubernetesCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.KubernetesClusterCreateRequest) error); ok {
		r1 = rf(create)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNodePool provides a mock function with given fields: clusterID, req
func (_m *KubernetesService) CreateNodePool(clusterID string, req *godo.KubernetesNodePoolCreateRequest) (*do.KubernetesNodePool, error) {
	ret := _m.Called(clusterID, req)

	var r0 *do.KubernetesNodePool
	if rf, ok := ret.Get(0).(func(string, *godo.KubernetesNodePoolCreateRequest) *do.KubernetesNodePool); ok {
		r0 = rf(clusterID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.KubernetesNodePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.KubernetesNodePoolCreateRequest) error); ok {
		r1 = rf(clusterID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: clusterID
func (_m *KubernetesService) Delete(clusterID string) error {
	ret := _m.Called(clusterID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(clusterID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNodePool provides a mock function with given fields: clusterID, poolID
func (_m *KubernetesService) DeleteNodePool(clusterID string, poolID string) error {
	ret := _m.Called(clusterID, poolID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(clusterID, poolID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: clusterID
func (_m *KubernetesService) Get(clusterID string) (*do.KubernetesCluster, error) {
	ret := _m.Called(clusterID)

	var r0 *do.KubernetesCluster
	if rf, ok := ret.Get(0).(func(string) *do.KubernetesCluster); ok {
		r0 = rf(clusterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.KubernetesCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKubeConfig provides a mock function with given fields: clusterID
func (_m *KubernetesService) GetKubeConfig(clusterID string) ([]byte, error) {
	ret := _m.Called(clusterID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(clusterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodePool provides a mock function with given fields: clusterID, poolID
func (_m *KubernetesService) GetNodePool(clusterID string, poolID string) (*do.KubernetesNodePool, error) {
	ret := _m.Called(clusterID, poolID)

	var r0 *do.KubernetesNodePool
	if rf, ok := ret.Get(0).(func(string, string) *do.KubernetesNodePool); ok {
		r0 = rf(clusterID, poolID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.KubernetesNodePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(clusterID, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOptions provides a mock function with given fields:
func (_m *KubernetesService) GetOptions() (*godo.KubernetesOptions, error) {
	ret := _m.Called()

	var r0 *godo.KubernetesOptions
	if rf, ok := ret.Get(0).(func() *godo.KubernetesOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*godo.KubernetesOptions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *KubernetesService) List() (do.KubernetesClusters, error) {
	ret := _m.Called()

	var r0 do.KubernetesClusters
	if rf, ok := ret.Get(0).(func() do.KubernetesClusters); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.KubernetesClusters)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNodePools provides a mock function with given fields: clusterID, opts
func (_m *KubernetesService) ListNodePools(clusterID string, opts *godo.ListOptions) (do.KubernetesNodePools, error) {
	ret := _m.Called(clusterID, opts)

	var r0 do.KubernetesNodePools
	if rf, ok := ret.Get(0).(func(string, *godo.ListOptions) do.KubernetesNodePools); ok {
		r0 = rf(clusterID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.KubernetesNodePools)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.ListOptions) error); ok {
		r1 = rf(clusterID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecycleNodePoolNodes provides a mock function with given fields: clusterID, poolID, req
func (_m *KubernetesService) RecycleNodePoolNodes(clusterID string, poolID string, req *godo.KubernetesNodePoolRecycleNodesRequest) error {
	ret := _m.Called(clusterID, poolID, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *godo.KubernetesNodePoolRecycleNodesRequest) error); ok {
		r0 = rf(clusterID, poolID, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: clusterID, update
func (_m *KubernetesService) Update(clusterID string, update *godo.KubernetesClusterUpdateRequest) (*do.KubernetesCluster, error) {
	ret := _m.Called(clusterID, update)

	var r0 *do.KubernetesCluster
	if rf, ok := ret.Get(0).(func(string, *godo.KubernetesClusterUpdateRequest) *do.KubernetesCluster); ok {
		r0 = rf(clusterID, update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.KubernetesCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.KubernetesClusterUpdateRequest) error); ok {
		r1 = rf(clusterID, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNodePool provides a mock function with given fields: clusterID, poolID, req
func (_m *KubernetesService) UpdateNodePool(clusterID string, poolID string, req *godo.KubernetesNodePoolUpdateRequest) (*do.KubernetesNodePool, error) {
	ret := _m.Called(clusterID, poolID, req)

	var r0 *do.KubernetesNodePool
	if rf, ok := ret.Get(0).(func(string, string, *godo.KubernetesNodePoolUpdateRequest) *do.KubernetesNodePool); ok {
		r0 = rf(clusterID, poolID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.KubernetesNodePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *godo.KubernetesNodePoolUpdateRequest) error); ok {
		r1 = rf(clusterID, poolID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
