// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import (
	core "github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	mock "github.com/stretchr/testify/mock"
)

// MockCoreCentralSystemHandler is an autogenerated mock type for the CentralSystemHandler type
type MockCoreCentralSystemHandler struct {
	mock.Mock
}

type MockCoreCentralSystemHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCoreCentralSystemHandler) EXPECT() *MockCoreCentralSystemHandler_Expecter {
	return &MockCoreCentralSystemHandler_Expecter{mock: &_m.Mock}
}

// OnAuthorize provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnAuthorize(chargePointId string, request *core.AuthorizeRequest) (*core.AuthorizeConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnAuthorize")
	}

	var r0 *core.AuthorizeConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.AuthorizeRequest) (*core.AuthorizeConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.AuthorizeRequest) *core.AuthorizeConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.AuthorizeConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.AuthorizeRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnAuthorize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnAuthorize'
type MockCoreCentralSystemHandler_OnAuthorize_Call struct {
	*mock.Call
}

// OnAuthorize is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.AuthorizeRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnAuthorize(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnAuthorize_Call {
	return &MockCoreCentralSystemHandler_OnAuthorize_Call{Call: _e.mock.On("OnAuthorize", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnAuthorize_Call) Run(run func(chargePointId string, request *core.AuthorizeRequest)) *MockCoreCentralSystemHandler_OnAuthorize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.AuthorizeRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnAuthorize_Call) Return(confirmation *core.AuthorizeConfirmation, err error) *MockCoreCentralSystemHandler_OnAuthorize_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnAuthorize_Call) RunAndReturn(run func(string, *core.AuthorizeRequest) (*core.AuthorizeConfirmation, error)) *MockCoreCentralSystemHandler_OnAuthorize_Call {
	_c.Call.Return(run)
	return _c
}

// OnBootNotification provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnBootNotification(chargePointId string, request *core.BootNotificationRequest) (*core.BootNotificationConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnBootNotification")
	}

	var r0 *core.BootNotificationConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.BootNotificationRequest) (*core.BootNotificationConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.BootNotificationRequest) *core.BootNotificationConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BootNotificationConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.BootNotificationRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnBootNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnBootNotification'
type MockCoreCentralSystemHandler_OnBootNotification_Call struct {
	*mock.Call
}

// OnBootNotification is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.BootNotificationRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnBootNotification(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnBootNotification_Call {
	return &MockCoreCentralSystemHandler_OnBootNotification_Call{Call: _e.mock.On("OnBootNotification", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnBootNotification_Call) Run(run func(chargePointId string, request *core.BootNotificationRequest)) *MockCoreCentralSystemHandler_OnBootNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.BootNotificationRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnBootNotification_Call) Return(confirmation *core.BootNotificationConfirmation, err error) *MockCoreCentralSystemHandler_OnBootNotification_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnBootNotification_Call) RunAndReturn(run func(string, *core.BootNotificationRequest) (*core.BootNotificationConfirmation, error)) *MockCoreCentralSystemHandler_OnBootNotification_Call {
	_c.Call.Return(run)
	return _c
}

// OnDataTransfer provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnDataTransfer(chargePointId string, request *core.DataTransferRequest) (*core.DataTransferConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnDataTransfer")
	}

	var r0 *core.DataTransferConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.DataTransferRequest) (*core.DataTransferConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.DataTransferRequest) *core.DataTransferConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.DataTransferConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.DataTransferRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnDataTransfer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnDataTransfer'
type MockCoreCentralSystemHandler_OnDataTransfer_Call struct {
	*mock.Call
}

// OnDataTransfer is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.DataTransferRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnDataTransfer(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnDataTransfer_Call {
	return &MockCoreCentralSystemHandler_OnDataTransfer_Call{Call: _e.mock.On("OnDataTransfer", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnDataTransfer_Call) Run(run func(chargePointId string, request *core.DataTransferRequest)) *MockCoreCentralSystemHandler_OnDataTransfer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.DataTransferRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnDataTransfer_Call) Return(confirmation *core.DataTransferConfirmation, err error) *MockCoreCentralSystemHandler_OnDataTransfer_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnDataTransfer_Call) RunAndReturn(run func(string, *core.DataTransferRequest) (*core.DataTransferConfirmation, error)) *MockCoreCentralSystemHandler_OnDataTransfer_Call {
	_c.Call.Return(run)
	return _c
}

// OnHeartbeat provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnHeartbeat(chargePointId string, request *core.HeartbeatRequest) (*core.HeartbeatConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnHeartbeat")
	}

	var r0 *core.HeartbeatConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.HeartbeatRequest) (*core.HeartbeatConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.HeartbeatRequest) *core.HeartbeatConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.HeartbeatConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.HeartbeatRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnHeartbeat'
type MockCoreCentralSystemHandler_OnHeartbeat_Call struct {
	*mock.Call
}

// OnHeartbeat is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.HeartbeatRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnHeartbeat(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnHeartbeat_Call {
	return &MockCoreCentralSystemHandler_OnHeartbeat_Call{Call: _e.mock.On("OnHeartbeat", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnHeartbeat_Call) Run(run func(chargePointId string, request *core.HeartbeatRequest)) *MockCoreCentralSystemHandler_OnHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.HeartbeatRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnHeartbeat_Call) Return(confirmation *core.HeartbeatConfirmation, err error) *MockCoreCentralSystemHandler_OnHeartbeat_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnHeartbeat_Call) RunAndReturn(run func(string, *core.HeartbeatRequest) (*core.HeartbeatConfirmation, error)) *MockCoreCentralSystemHandler_OnHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// OnMeterValues provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnMeterValues(chargePointId string, request *core.MeterValuesRequest) (*core.MeterValuesConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnMeterValues")
	}

	var r0 *core.MeterValuesConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.MeterValuesRequest) (*core.MeterValuesConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.MeterValuesRequest) *core.MeterValuesConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.MeterValuesConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.MeterValuesRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnMeterValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnMeterValues'
type MockCoreCentralSystemHandler_OnMeterValues_Call struct {
	*mock.Call
}

// OnMeterValues is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.MeterValuesRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnMeterValues(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnMeterValues_Call {
	return &MockCoreCentralSystemHandler_OnMeterValues_Call{Call: _e.mock.On("OnMeterValues", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnMeterValues_Call) Run(run func(chargePointId string, request *core.MeterValuesRequest)) *MockCoreCentralSystemHandler_OnMeterValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.MeterValuesRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnMeterValues_Call) Return(confirmation *core.MeterValuesConfirmation, err error) *MockCoreCentralSystemHandler_OnMeterValues_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnMeterValues_Call) RunAndReturn(run func(string, *core.MeterValuesRequest) (*core.MeterValuesConfirmation, error)) *MockCoreCentralSystemHandler_OnMeterValues_Call {
	_c.Call.Return(run)
	return _c
}

// OnStartTransaction provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnStartTransaction(chargePointId string, request *core.StartTransactionRequest) (*core.StartTransactionConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnStartTransaction")
	}

	var r0 *core.StartTransactionConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.StartTransactionRequest) (*core.StartTransactionConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.StartTransactionRequest) *core.StartTransactionConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.StartTransactionConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.StartTransactionRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnStartTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnStartTransaction'
type MockCoreCentralSystemHandler_OnStartTransaction_Call struct {
	*mock.Call
}

// OnStartTransaction is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.StartTransactionRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnStartTransaction(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnStartTransaction_Call {
	return &MockCoreCentralSystemHandler_OnStartTransaction_Call{Call: _e.mock.On("OnStartTransaction", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnStartTransaction_Call) Run(run func(chargePointId string, request *core.StartTransactionRequest)) *MockCoreCentralSystemHandler_OnStartTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.StartTransactionRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnStartTransaction_Call) Return(confirmation *core.StartTransactionConfirmation, err error) *MockCoreCentralSystemHandler_OnStartTransaction_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnStartTransaction_Call) RunAndReturn(run func(string, *core.StartTransactionRequest) (*core.StartTransactionConfirmation, error)) *MockCoreCentralSystemHandler_OnStartTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// OnStatusNotification provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnStatusNotification(chargePointId string, request *core.StatusNotificationRequest) (*core.StatusNotificationConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnStatusNotification")
	}

	var r0 *core.StatusNotificationConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.StatusNotificationRequest) (*core.StatusNotificationConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.StatusNotificationRequest) *core.StatusNotificationConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.StatusNotificationConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.StatusNotificationRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnStatusNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnStatusNotification'
type MockCoreCentralSystemHandler_OnStatusNotification_Call struct {
	*mock.Call
}

// OnStatusNotification is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.StatusNotificationRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnStatusNotification(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnStatusNotification_Call {
	return &MockCoreCentralSystemHandler_OnStatusNotification_Call{Call: _e.mock.On("OnStatusNotification", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnStatusNotification_Call) Run(run func(chargePointId string, request *core.StatusNotificationRequest)) *MockCoreCentralSystemHandler_OnStatusNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.StatusNotificationRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnStatusNotification_Call) Return(confirmation *core.StatusNotificationConfirmation, err error) *MockCoreCentralSystemHandler_OnStatusNotification_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnStatusNotification_Call) RunAndReturn(run func(string, *core.StatusNotificationRequest) (*core.StatusNotificationConfirmation, error)) *MockCoreCentralSystemHandler_OnStatusNotification_Call {
	_c.Call.Return(run)
	return _c
}

// OnStopTransaction provides a mock function with given fields: chargePointId, request
func (_m *MockCoreCentralSystemHandler) OnStopTransaction(chargePointId string, request *core.StopTransactionRequest) (*core.StopTransactionConfirmation, error) {
	ret := _m.Called(chargePointId, request)

	if len(ret) == 0 {
		panic("no return value specified for OnStopTransaction")
	}

	var r0 *core.StopTransactionConfirmation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *core.StopTransactionRequest) (*core.StopTransactionConfirmation, error)); ok {
		return rf(chargePointId, request)
	}
	if rf, ok := ret.Get(0).(func(string, *core.StopTransactionRequest) *core.StopTransactionConfirmation); ok {
		r0 = rf(chargePointId, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.StopTransactionConfirmation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *core.StopTransactionRequest) error); ok {
		r1 = rf(chargePointId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCoreCentralSystemHandler_OnStopTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnStopTransaction'
type MockCoreCentralSystemHandler_OnStopTransaction_Call struct {
	*mock.Call
}

// OnStopTransaction is a helper method to define mock.On call
//   - chargePointId string
//   - request *core.StopTransactionRequest
func (_e *MockCoreCentralSystemHandler_Expecter) OnStopTransaction(chargePointId interface{}, request interface{}) *MockCoreCentralSystemHandler_OnStopTransaction_Call {
	return &MockCoreCentralSystemHandler_OnStopTransaction_Call{Call: _e.mock.On("OnStopTransaction", chargePointId, request)}
}

func (_c *MockCoreCentralSystemHandler_OnStopTransaction_Call) Run(run func(chargePointId string, request *core.StopTransactionRequest)) *MockCoreCentralSystemHandler_OnStopTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*core.StopTransactionRequest))
	})
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnStopTransaction_Call) Return(confirmation *core.StopTransactionConfirmation, err error) *MockCoreCentralSystemHandler_OnStopTransaction_Call {
	_c.Call.Return(confirmation, err)
	return _c
}

func (_c *MockCoreCentralSystemHandler_OnStopTransaction_Call) RunAndReturn(run func(string, *core.StopTransactionRequest) (*core.StopTransactionConfirmation, error)) *MockCoreCentralSystemHandler_OnStopTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCoreCentralSystemHandler creates a new instance of MockCoreCentralSystemHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCoreCentralSystemHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCoreCentralSystemHandler {
	mock := &MockCoreCentralSystemHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
