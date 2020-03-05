package service

import (
	`context`

	"github.com/stretchr/testify/mock"

	`github.com/LuLStackCoder/test-service/pkg/models`
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetUser(ctx context.Context, request *models.Request) (response models.Response, err error) {
	args := m.Called(context.Background(), request)
	if a, ok := args.Get(0).(models.Response); ok {
		return a, args.Error(1)
	}
	return response, args.Error(0)
}

func (m *Mock) PostOrder(ctx context.Context, request *models.Request) (response models.Response, err error) {
	args := m.Called(context.Background(), request)
	if a, ok := args.Get(0).(models.Response); ok {
		return a, args.Error(1)
	}
	return response, args.Error(0)
}

func (m *Mock) GetCount(ctx context.Context, request *models.Request) (response models.Response, err error) {
	args := m.Called(context.Background(), request)
	if a, ok := args.Get(0).(models.Response); ok {
		return a, args.Error(1)
	}
	return response, args.Error(0)
}

func (m *Mock) GetOrder(ctx context.Context) (response models.Response, err error) {
	args := m.Called(context.Background())
	if a, ok := args.Get(0).(models.Response); ok {
		return a, args.Error(1)
	}
	return response, args.Error(0)
}
