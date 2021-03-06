// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	todo_app "github.com/Hanqur/todo_app"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user todo_app.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockTodoList is a mock of TodoList interface.
type MockTodoList struct {
	ctrl     *gomock.Controller
	recorder *MockTodoListMockRecorder
}

// MockTodoListMockRecorder is the mock recorder for MockTodoList.
type MockTodoListMockRecorder struct {
	mock *MockTodoList
}

// NewMockTodoList creates a new mock instance.
func NewMockTodoList(ctrl *gomock.Controller) *MockTodoList {
	mock := &MockTodoList{ctrl: ctrl}
	mock.recorder = &MockTodoListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoList) EXPECT() *MockTodoListMockRecorder {
	return m.recorder
}

// CreateList mocks base method.
func (m *MockTodoList) CreateList(userId int, list todo_app.TodoList) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", userId, list)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockTodoListMockRecorder) CreateList(userId, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockTodoList)(nil).CreateList), userId, list)
}

// DeleteList mocks base method.
func (m *MockTodoList) DeleteList(userId, listId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", userId, listId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockTodoListMockRecorder) DeleteList(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockTodoList)(nil).DeleteList), userId, listId)
}

// GetAllLists mocks base method.
func (m *MockTodoList) GetAllLists(userId int) ([]todo_app.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLists", userId)
	ret0, _ := ret[0].([]todo_app.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLists indicates an expected call of GetAllLists.
func (mr *MockTodoListMockRecorder) GetAllLists(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLists", reflect.TypeOf((*MockTodoList)(nil).GetAllLists), userId)
}

// GetListById mocks base method.
func (m *MockTodoList) GetListById(userId, listId int) (todo_app.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListById", userId, listId)
	ret0, _ := ret[0].(todo_app.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListById indicates an expected call of GetListById.
func (mr *MockTodoListMockRecorder) GetListById(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListById", reflect.TypeOf((*MockTodoList)(nil).GetListById), userId, listId)
}

// UpdateList mocks base method.
func (m *MockTodoList) UpdateList(userId, listId int, input todo_app.UpdateListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", userId, listId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockTodoListMockRecorder) UpdateList(userId, listId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockTodoList)(nil).UpdateList), userId, listId, input)
}

// MockTodoItem is a mock of TodoItem interface.
type MockTodoItem struct {
	ctrl     *gomock.Controller
	recorder *MockTodoItemMockRecorder
}

// MockTodoItemMockRecorder is the mock recorder for MockTodoItem.
type MockTodoItemMockRecorder struct {
	mock *MockTodoItem
}

// NewMockTodoItem creates a new mock instance.
func NewMockTodoItem(ctrl *gomock.Controller) *MockTodoItem {
	mock := &MockTodoItem{ctrl: ctrl}
	mock.recorder = &MockTodoItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoItem) EXPECT() *MockTodoItemMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockTodoItem) CreateItem(userId, listId int, input todo_app.Item) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", userId, listId, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockTodoItemMockRecorder) CreateItem(userId, listId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockTodoItem)(nil).CreateItem), userId, listId, input)
}

// DeleteItem mocks base method.
func (m *MockTodoItem) DeleteItem(userId, itemId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", userId, itemId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockTodoItemMockRecorder) DeleteItem(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockTodoItem)(nil).DeleteItem), userId, itemId)
}

// GetAllItems mocks base method.
func (m *MockTodoItem) GetAllItems(userId, listId int) ([]todo_app.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllItems", userId, listId)
	ret0, _ := ret[0].([]todo_app.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllItems indicates an expected call of GetAllItems.
func (mr *MockTodoItemMockRecorder) GetAllItems(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllItems", reflect.TypeOf((*MockTodoItem)(nil).GetAllItems), userId, listId)
}

// GetItemById mocks base method.
func (m *MockTodoItem) GetItemById(userId, itemId int) (todo_app.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemById", userId, itemId)
	ret0, _ := ret[0].(todo_app.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemById indicates an expected call of GetItemById.
func (mr *MockTodoItemMockRecorder) GetItemById(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemById", reflect.TypeOf((*MockTodoItem)(nil).GetItemById), userId, itemId)
}

// UpdateItem mocks base method.
func (m *MockTodoItem) UpdateItem(userId, itemId int, input todo_app.UpdateItemInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", userId, itemId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockTodoItemMockRecorder) UpdateItem(userId, itemId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockTodoItem)(nil).UpdateItem), userId, itemId, input)
}

// MockTag is a mock of Tag interface.
type MockTag struct {
	ctrl     *gomock.Controller
	recorder *MockTagMockRecorder
}

// MockTagMockRecorder is the mock recorder for MockTag.
type MockTagMockRecorder struct {
	mock *MockTag
}

// NewMockTag creates a new mock instance.
func NewMockTag(ctrl *gomock.Controller) *MockTag {
	mock := &MockTag{ctrl: ctrl}
	mock.recorder = &MockTagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTag) EXPECT() *MockTagMockRecorder {
	return m.recorder
}

// CreateTag mocks base method.
func (m *MockTag) CreateTag(userId, itemId int, input todo_app.Tag) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", userId, itemId, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockTagMockRecorder) CreateTag(userId, itemId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockTag)(nil).CreateTag), userId, itemId, input)
}

// GetAllTags mocks base method.
func (m *MockTag) GetAllTags(userId, itemId int) ([]todo_app.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTags", userId, itemId)
	ret0, _ := ret[0].([]todo_app.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTags indicates an expected call of GetAllTags.
func (mr *MockTagMockRecorder) GetAllTags(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTags", reflect.TypeOf((*MockTag)(nil).GetAllTags), userId, itemId)
}
