// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	config "github.com/johnfercher/maroto/pkg/v2/config"
	domain "github.com/johnfercher/maroto/pkg/v2/domain"

	internal "github.com/johnfercher/maroto/internal"

	mock "github.com/stretchr/testify/mock"

	tree "github.com/johnfercher/go-tree/tree"
)

// Component is an autogenerated mock type for the Component type
type Component struct {
	mock.Mock
}

type Component_Expecter struct {
	mock *mock.Mock
}

func (_m *Component) EXPECT() *Component_Expecter {
	return &Component_Expecter{mock: &_m.Mock}
}

// GetStructure provides a mock function with given fields:
func (_m *Component) GetStructure() *tree.Node[domain.Structure] {
	ret := _m.Called()

	var r0 *tree.Node[domain.Structure]
	if rf, ok := ret.Get(0).(func() *tree.Node[domain.Structure]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tree.Node[domain.Structure])
		}
	}

	return r0
}

// Component_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type Component_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *Component_Expecter) GetStructure() *Component_GetStructure_Call {
	return &Component_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *Component_GetStructure_Call) Run(run func()) *Component_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Component_GetStructure_Call) Return(_a0 *tree.Node[domain.Structure]) *Component_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Component_GetStructure_Call) RunAndReturn(run func() *tree.Node[domain.Structure]) *Component_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// Render provides a mock function with given fields: provider, cell
func (_m *Component) Render(provider domain.Provider, cell internal.Cell) {
	_m.Called(provider, cell)
}

// Component_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type Component_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - provider domain.Provider
//   - cell internal.Cell
func (_e *Component_Expecter) Render(provider interface{}, cell interface{}) *Component_Render_Call {
	return &Component_Render_Call{Call: _e.mock.On("Render", provider, cell)}
}

func (_c *Component_Render_Call) Run(run func(provider domain.Provider, cell internal.Cell)) *Component_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.Provider), args[1].(internal.Cell))
	})
	return _c
}

func (_c *Component_Render_Call) Return() *Component_Render_Call {
	_c.Call.Return()
	return _c
}

func (_c *Component_Render_Call) RunAndReturn(run func(domain.Provider, internal.Cell)) *Component_Render_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfig provides a mock function with given fields: _a0
func (_m *Component) SetConfig(_a0 *config.Maroto) {
	_m.Called(_a0)
}

// Component_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type Component_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//   - _a0 *config.Maroto
func (_e *Component_Expecter) SetConfig(_a0 interface{}) *Component_SetConfig_Call {
	return &Component_SetConfig_Call{Call: _e.mock.On("SetConfig", _a0)}
}

func (_c *Component_SetConfig_Call) Run(run func(_a0 *config.Maroto)) *Component_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config.Maroto))
	})
	return _c
}

func (_c *Component_SetConfig_Call) Return() *Component_SetConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *Component_SetConfig_Call) RunAndReturn(run func(*config.Maroto)) *Component_SetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewComponent creates a new instance of Component. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComponent(t interface {
	mock.TestingT
	Cleanup(func())
}) *Component {
	mock := &Component{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}