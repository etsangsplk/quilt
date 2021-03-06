package google

import "github.com/stretchr/testify/mock"

import compute "google.golang.org/api/compute/v1"

// mockClient is an autogenerated mock type for the client type
type mockClient struct {
	mock.Mock
}

// AddAccessConfig provides a mock function with given fields: zone, instance, networkInterface, accessConfig
func (_m *mockClient) AddAccessConfig(zone string, instance string, networkInterface string, accessConfig *compute.AccessConfig) (*compute.Operation, error) {
	ret := _m.Called(zone, instance, networkInterface, accessConfig)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string, string, string, *compute.AccessConfig) *compute.Operation); ok {
		r0 = rf(zone, instance, networkInterface, accessConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, *compute.AccessConfig) error); ok {
		r1 = rf(zone, instance, networkInterface, accessConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccessConfig provides a mock function with given fields: zone, instance, accessConfig, networkInterface
func (_m *mockClient) DeleteAccessConfig(zone string, instance string, accessConfig string, networkInterface string) (*compute.Operation, error) {
	ret := _m.Called(zone, instance, accessConfig, networkInterface)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string, string, string, string) *compute.Operation); ok {
		r0 = rf(zone, instance, accessConfig, networkInterface)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(zone, instance, accessConfig, networkInterface)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFirewall provides a mock function with given fields: firewall
func (_m *mockClient) DeleteFirewall(firewall string) (*compute.Operation, error) {
	ret := _m.Called(firewall)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string) *compute.Operation); ok {
		r0 = rf(firewall)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(firewall)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInstance provides a mock function with given fields: zone, operation
func (_m *mockClient) DeleteInstance(zone string, operation string) (*compute.Operation, error) {
	ret := _m.Called(zone, operation)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string, string) *compute.Operation); ok {
		r0 = rf(zone, operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(zone, operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGlobalOperation provides a mock function with given fields: operation
func (_m *mockClient) GetGlobalOperation(operation string) (*compute.Operation, error) {
	ret := _m.Called(operation)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string) *compute.Operation); ok {
		r0 = rf(operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstance provides a mock function with given fields: zone, id
func (_m *mockClient) GetInstance(zone string, id string) (*compute.Instance, error) {
	ret := _m.Called(zone, id)

	var r0 *compute.Instance
	if rf, ok := ret.Get(0).(func(string, string) *compute.Instance); ok {
		r0 = rf(zone, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(zone, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetZoneOperation provides a mock function with given fields: zone, operation
func (_m *mockClient) GetZoneOperation(zone string, operation string) (*compute.Operation, error) {
	ret := _m.Called(zone, operation)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string, string) *compute.Operation); ok {
		r0 = rf(zone, operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(zone, operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertFirewall provides a mock function with given fields: firewall
func (_m *mockClient) InsertFirewall(firewall *compute.Firewall) (*compute.Operation, error) {
	ret := _m.Called(firewall)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(*compute.Firewall) *compute.Operation); ok {
		r0 = rf(firewall)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*compute.Firewall) error); ok {
		r1 = rf(firewall)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertInstance provides a mock function with given fields: zone, instance
func (_m *mockClient) InsertInstance(zone string, instance *compute.Instance) (*compute.Operation, error) {
	ret := _m.Called(zone, instance)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string, *compute.Instance) *compute.Operation); ok {
		r0 = rf(zone, instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *compute.Instance) error); ok {
		r1 = rf(zone, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertNetwork provides a mock function with given fields: network
func (_m *mockClient) InsertNetwork(network *compute.Network) (*compute.Operation, error) {
	ret := _m.Called(network)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(*compute.Network) *compute.Operation); ok {
		r0 = rf(network)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*compute.Network) error); ok {
		r1 = rf(network)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFirewalls provides a mock function with given fields:
func (_m *mockClient) ListFirewalls() (*compute.FirewallList, error) {
	ret := _m.Called()

	var r0 *compute.FirewallList
	if rf, ok := ret.Get(0).(func() *compute.FirewallList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.FirewallList)
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

// ListInstances provides a mock function with given fields: zone, opts
func (_m *mockClient) ListInstances(zone string, opts apiOptions) (*compute.InstanceList, error) {
	ret := _m.Called(zone, opts)

	var r0 *compute.InstanceList
	if rf, ok := ret.Get(0).(func(string, apiOptions) *compute.InstanceList); ok {
		r0 = rf(zone, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.InstanceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, apiOptions) error); ok {
		r1 = rf(zone, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNetworks provides a mock function with given fields:
func (_m *mockClient) ListNetworks() (*compute.NetworkList, error) {
	ret := _m.Called()

	var r0 *compute.NetworkList
	if rf, ok := ret.Get(0).(func() *compute.NetworkList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.NetworkList)
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

// PatchFirewall provides a mock function with given fields: name, firewall
func (_m *mockClient) PatchFirewall(name string, firewall *compute.Firewall) (*compute.Operation, error) {
	ret := _m.Called(name, firewall)

	var r0 *compute.Operation
	if rf, ok := ret.Get(0).(func(string, *compute.Firewall) *compute.Operation); ok {
		r0 = rf(name, firewall)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *compute.Firewall) error); ok {
		r1 = rf(name, firewall)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
