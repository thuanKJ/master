package data

import (
	"go_songs/models"

	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/mock"
)

type MockConnectdata struct {
	mock.Mock
}

func (m *MockConnectdata) FindAll() ([]models.Song, error) {
	args := m.Mock.Called()
	return args.Get(0).([]models.Song), args.Error(1)
}

func (m *MockConnectdata) FindById(id string) (models.Song, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(models.Song), args.Error(1)
}

func (m *MockConnectdata) CountId(id string) int {
	args := m.Mock.Called(id)
	return args.Int(0)
}

func (m *MockConnectdata) MaxId() int {
	args := m.Mock.Called()
	return args.Int(0)
}

func (m *MockConnectdata) Insert(data models.Song) error {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockConnectdata) Delete(id string) error {
	args := m.Mock.Called(id)
	return args.Error(0)
}

func (m *MockConnectdata) Update(id string, k bson.M) error {
	args := m.Mock.Called(id, k)
	return args.Error(0)
}

func (m *MockConnectdata) FindByName(name string) (models.Account, error) {
	args := m.Mock.Called(name)
	return args.Get(0).(models.Account), args.Error(1)
}

func (m *MockConnectdata) AddNewAccount(account models.Account) error {
	args := m.Mock.Called(account)
	return args.Error(0)
}

func (m *MockConnectdata) CheckByName(name string) (bool, error) {
	args := m.Mock.Called(name)
	return args.Bool(0), args.Error(1)
}
