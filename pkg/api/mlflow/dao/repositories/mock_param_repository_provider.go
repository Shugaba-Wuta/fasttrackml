// Code generated by mockery v2.34.0. DO NOT EDIT.

package repositories

import (
	context "context"

	models "github.com/G-Research/fasttrackml/pkg/api/mlflow/dao/models"
	mock "github.com/stretchr/testify/mock"
)

// MockParamRepositoryProvider is an autogenerated mock type for the ParamRepositoryProvider type
type MockParamRepositoryProvider struct {
	mock.Mock
}

// CreateBatch provides a mock function with given fields: ctx, batchSize, params
func (_m *MockParamRepositoryProvider) CreateBatch(ctx context.Context, batchSize int, params []models.Param) error {
	ret := _m.Called(ctx, batchSize, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []models.Param) error); ok {
		r0 = rf(ctx, batchSize, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockParamRepositoryProvider creates a new instance of MockParamRepositoryProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockParamRepositoryProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockParamRepositoryProvider {
	mock := &MockParamRepositoryProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}