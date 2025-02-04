// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	devicemonitor "github.com/mendersoftware/inventory/client/devicemonitor"
	inv "github.com/mendersoftware/inventory/inv"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mendersoftware/inventory/model"

	store "github.com/mendersoftware/inventory/store"

	workflows "github.com/mendersoftware/inventory/client/workflows"
)

// InventoryApp is an autogenerated mock type for the InventoryApp type
type InventoryApp struct {
	mock.Mock
}

// AddDevice provides a mock function with given fields: ctx, d
func (_m *InventoryApp) AddDevice(ctx context.Context, d *model.Device) error {
	ret := _m.Called(ctx, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Device) error); ok {
		r0 = rf(ctx, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckAlerts provides a mock function with given fields: ctx, deviceId
func (_m *InventoryApp) CheckAlerts(ctx context.Context, deviceId string) (int, error) {
	ret := _m.Called(ctx, deviceId)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, deviceId)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deviceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTenant provides a mock function with given fields: ctx, tenant
func (_m *InventoryApp) CreateTenant(ctx context.Context, tenant model.NewTenant) error {
	ret := _m.Called(ctx, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.NewTenant) error); ok {
		r0 = rf(ctx, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevice provides a mock function with given fields: ctx, id
func (_m *InventoryApp) DeleteDevice(ctx context.Context, id model.DeviceID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevices provides a mock function with given fields: ctx, ids
func (_m *InventoryApp) DeleteDevices(ctx context.Context, ids []model.DeviceID) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, ids)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID) *model.UpdateResult); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroup provides a mock function with given fields: ctx, groupName
func (_m *InventoryApp) DeleteGroup(ctx context.Context, groupName model.GroupName) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, groupName)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, model.GroupName) *model.UpdateResult); ok {
		r0 = rf(ctx, groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.GroupName) error); ok {
		r1 = rf(ctx, groupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, id
func (_m *InventoryApp) GetDevice(ctx context.Context, id model.DeviceID) (*model.Device, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID) *model.Device); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.DeviceID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceGroup provides a mock function with given fields: ctx, id
func (_m *InventoryApp) GetDeviceGroup(ctx context.Context, id model.DeviceID) (model.GroupName, error) {
	ret := _m.Called(ctx, id)

	var r0 model.GroupName
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID) model.GroupName); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.GroupName)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.DeviceID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFiltersAttributes provides a mock function with given fields: ctx
func (_m *InventoryApp) GetFiltersAttributes(ctx context.Context) ([]model.FilterAttribute, error) {
	ret := _m.Called(ctx)

	var r0 []model.FilterAttribute
	if rf, ok := ret.Get(0).(func(context.Context) []model.FilterAttribute); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.FilterAttribute)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *InventoryApp) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListDevices provides a mock function with given fields: ctx, q
func (_m *InventoryApp) ListDevices(ctx context.Context, q store.ListQuery) ([]model.Device, int, error) {
	ret := _m.Called(ctx, q)

	var r0 []model.Device
	if rf, ok := ret.Get(0).(func(context.Context, store.ListQuery) []model.Device); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Device)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, store.ListQuery) int); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, store.ListQuery) error); ok {
		r2 = rf(ctx, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListDevicesByGroup provides a mock function with given fields: ctx, group, skip, limit
func (_m *InventoryApp) ListDevicesByGroup(ctx context.Context, group model.GroupName, skip int, limit int) ([]model.DeviceID, int, error) {
	ret := _m.Called(ctx, group, skip, limit)

	var r0 []model.DeviceID
	if rf, ok := ret.Get(0).(func(context.Context, model.GroupName, int, int) []model.DeviceID); ok {
		r0 = rf(ctx, group, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DeviceID)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, model.GroupName, int, int) int); ok {
		r1 = rf(ctx, group, skip, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.GroupName, int, int) error); ok {
		r2 = rf(ctx, group, skip, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListGroups provides a mock function with given fields: ctx, filters
func (_m *InventoryApp) ListGroups(ctx context.Context, filters []model.FilterPredicate) ([]model.GroupName, error) {
	ret := _m.Called(ctx, filters)

	var r0 []model.GroupName
	if rf, ok := ret.Get(0).(func(context.Context, []model.FilterPredicate) []model.GroupName); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GroupName)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.FilterPredicate) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceAttributes provides a mock function with given fields: ctx, id, upsertAttrs, scope, etag
func (_m *InventoryApp) ReplaceAttributes(ctx context.Context, id model.DeviceID, upsertAttrs model.DeviceAttributes, scope string, etag string) error {
	ret := _m.Called(ctx, id, upsertAttrs, scope, etag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID, model.DeviceAttributes, string, string) error); ok {
		r0 = rf(ctx, id, upsertAttrs, scope, etag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchDevices provides a mock function with given fields: ctx, searchParams
func (_m *InventoryApp) SearchDevices(ctx context.Context, searchParams model.SearchParams) ([]model.Device, int, error) {
	ret := _m.Called(ctx, searchParams)

	var r0 []model.Device
	if rf, ok := ret.Get(0).(func(context.Context, model.SearchParams) []model.Device); ok {
		r0 = rf(ctx, searchParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Device)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, model.SearchParams) int); ok {
		r1 = rf(ctx, searchParams)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.SearchParams) error); ok {
		r2 = rf(ctx, searchParams)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UnsetDeviceGroup provides a mock function with given fields: ctx, id, groupName
func (_m *InventoryApp) UnsetDeviceGroup(ctx context.Context, id model.DeviceID, groupName model.GroupName) error {
	ret := _m.Called(ctx, id, groupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID, model.GroupName) error); ok {
		r0 = rf(ctx, id, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnsetDevicesGroup provides a mock function with given fields: ctx, deviceIDs, groupName
func (_m *InventoryApp) UnsetDevicesGroup(ctx context.Context, deviceIDs []model.DeviceID, groupName model.GroupName) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, deviceIDs, groupName)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID, model.GroupName) *model.UpdateResult); ok {
		r0 = rf(ctx, deviceIDs, groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID, model.GroupName) error); ok {
		r1 = rf(ctx, deviceIDs, groupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeviceGroup provides a mock function with given fields: ctx, id, group
func (_m *InventoryApp) UpdateDeviceGroup(ctx context.Context, id model.DeviceID, group model.GroupName) error {
	ret := _m.Called(ctx, id, group)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID, model.GroupName) error); ok {
		r0 = rf(ctx, id, group)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDevicesGroup provides a mock function with given fields: ctx, ids, group
func (_m *InventoryApp) UpdateDevicesGroup(ctx context.Context, ids []model.DeviceID, group model.GroupName) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, ids, group)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceID, model.GroupName) *model.UpdateResult); ok {
		r0 = rf(ctx, ids, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceID, model.GroupName) error); ok {
		r1 = rf(ctx, ids, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertAttributes provides a mock function with given fields: ctx, id, attrs
func (_m *InventoryApp) UpsertAttributes(ctx context.Context, id model.DeviceID, attrs model.DeviceAttributes) error {
	ret := _m.Called(ctx, id, attrs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID, model.DeviceAttributes) error); ok {
		r0 = rf(ctx, id, attrs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertAttributesWithUpdated provides a mock function with given fields: ctx, id, attrs, scope, etag
func (_m *InventoryApp) UpsertAttributesWithUpdated(ctx context.Context, id model.DeviceID, attrs model.DeviceAttributes, scope string, etag string) error {
	ret := _m.Called(ctx, id, attrs, scope, etag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceID, model.DeviceAttributes, string, string) error); ok {
		r0 = rf(ctx, id, attrs, scope, etag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertDevicesStatuses provides a mock function with given fields: ctx, devices, attrs
func (_m *InventoryApp) UpsertDevicesStatuses(ctx context.Context, devices []model.DeviceUpdate, attrs model.DeviceAttributes) (*model.UpdateResult, error) {
	ret := _m.Called(ctx, devices, attrs)

	var r0 *model.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, []model.DeviceUpdate, model.DeviceAttributes) *model.UpdateResult); ok {
		r0 = rf(ctx, devices, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.DeviceUpdate, model.DeviceAttributes) error); ok {
		r1 = rf(ctx, devices, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithDevicemonitor provides a mock function with given fields: client
func (_m *InventoryApp) WithDevicemonitor(client devicemonitor.Client) inv.InventoryApp {
	ret := _m.Called(client)

	var r0 inv.InventoryApp
	if rf, ok := ret.Get(0).(func(devicemonitor.Client) inv.InventoryApp); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(inv.InventoryApp)
		}
	}

	return r0
}

// WithLimits provides a mock function with given fields: attributes, tags
func (_m *InventoryApp) WithLimits(attributes int, tags int) inv.InventoryApp {
	ret := _m.Called(attributes, tags)

	var r0 inv.InventoryApp
	if rf, ok := ret.Get(0).(func(int, int) inv.InventoryApp); ok {
		r0 = rf(attributes, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(inv.InventoryApp)
		}
	}

	return r0
}

// WithReporting provides a mock function with given fields: c
func (_m *InventoryApp) WithReporting(c workflows.Client) inv.InventoryApp {
	ret := _m.Called(c)

	var r0 inv.InventoryApp
	if rf, ok := ret.Get(0).(func(workflows.Client) inv.InventoryApp); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(inv.InventoryApp)
		}
	}

	return r0
}
