// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

type Reader_Expecter struct {
	mock *mock.Mock
}

func (_m *Reader) EXPECT() *Reader_Expecter {
	return &Reader_Expecter{mock: &_m.Mock}
}

// ContentFile provides a mock function with given fields: index
func (_m *Reader) ContentFile(index int) ([]byte, error) {
	ret := _m.Called(index)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]byte, error)); ok {
		return rf(index)
	}
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reader_ContentFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ContentFile'
type Reader_ContentFile_Call struct {
	*mock.Call
}

// ContentFile is a helper method to define mock.On call
//   - index int
func (_e *Reader_Expecter) ContentFile(index interface{}) *Reader_ContentFile_Call {
	return &Reader_ContentFile_Call{Call: _e.mock.On("ContentFile", index)}
}

func (_c *Reader_ContentFile_Call) Run(run func(index int)) *Reader_ContentFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Reader_ContentFile_Call) Return(_a0 []byte, _a1 error) *Reader_ContentFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Reader_ContentFile_Call) RunAndReturn(run func(int) ([]byte, error)) *Reader_ContentFile_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: filepath
func (_m *Reader) Create(filepath string) error {
	ret := _m.Called(filepath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filepath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reader_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Reader_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - filepath string
func (_e *Reader_Expecter) Create(filepath interface{}) *Reader_Create_Call {
	return &Reader_Create_Call{Call: _e.mock.On("Create", filepath)}
}

func (_c *Reader_Create_Call) Run(run func(filepath string)) *Reader_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Reader_Create_Call) Return(_a0 error) *Reader_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reader_Create_Call) RunAndReturn(run func(string) error) *Reader_Create_Call {
	_c.Call.Return(run)
	return _c
}

// InfoFile provides a mock function with given fields: index
func (_m *Reader) InfoFile(index int) (fs.FileInfo, error) {
	ret := _m.Called(index)

	var r0 fs.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (fs.FileInfo, error)); ok {
		return rf(index)
	}
	if rf, ok := ret.Get(0).(func(int) fs.FileInfo); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reader_InfoFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InfoFile'
type Reader_InfoFile_Call struct {
	*mock.Call
}

// InfoFile is a helper method to define mock.On call
//   - index int
func (_e *Reader_Expecter) InfoFile(index interface{}) *Reader_InfoFile_Call {
	return &Reader_InfoFile_Call{Call: _e.mock.On("InfoFile", index)}
}

func (_c *Reader_InfoFile_Call) Run(run func(index int)) *Reader_InfoFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Reader_InfoFile_Call) Return(_a0 fs.FileInfo, _a1 error) *Reader_InfoFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Reader_InfoFile_Call) RunAndReturn(run func(int) (fs.FileInfo, error)) *Reader_InfoFile_Call {
	_c.Call.Return(run)
	return _c
}

// NFiles provides a mock function with given fields:
func (_m *Reader) NFiles() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Reader_NFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NFiles'
type Reader_NFiles_Call struct {
	*mock.Call
}

// NFiles is a helper method to define mock.On call
func (_e *Reader_Expecter) NFiles() *Reader_NFiles_Call {
	return &Reader_NFiles_Call{Call: _e.mock.On("NFiles")}
}

func (_c *Reader_NFiles_Call) Run(run func()) *Reader_NFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Reader_NFiles_Call) Return(_a0 int) *Reader_NFiles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reader_NFiles_Call) RunAndReturn(run func() int) *Reader_NFiles_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFile provides a mock function with given fields: index, filepath
func (_m *Reader) WriteFile(index int, filepath string) error {
	ret := _m.Called(index, filepath)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(index, filepath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reader_WriteFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFile'
type Reader_WriteFile_Call struct {
	*mock.Call
}

// WriteFile is a helper method to define mock.On call
//   - index int
//   - filepath string
func (_e *Reader_Expecter) WriteFile(index interface{}, filepath interface{}) *Reader_WriteFile_Call {
	return &Reader_WriteFile_Call{Call: _e.mock.On("WriteFile", index, filepath)}
}

func (_c *Reader_WriteFile_Call) Run(run func(index int, filepath string)) *Reader_WriteFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string))
	})
	return _c
}

func (_c *Reader_WriteFile_Call) Return(_a0 error) *Reader_WriteFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reader_WriteFile_Call) RunAndReturn(run func(int, string) error) *Reader_WriteFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
