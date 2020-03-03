package service

import (
	`context`
	`fmt`

	`github.com/LuLStackCoder/test-service/pkg/models`
)

type Service interface {
	GetUser(ctx context.Context, request *models.Request) (response models.Response, err error)
	PostOrder(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetCount(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetOrder(ctx context.Context) (response models.Response, err error)
}

type service struct {
}

func (s *service) GetUser(ctx context.Context, request *models.Request) (response models.Response, err error) {
	fmt.Printf("id=%d", request.Id)
	if request.Id > 0 {
		response.Data = &models.DataStruct{Res: true}
		return
	}
	errStr := "id < 0"
	response = models.Response{Error: true, ErrorText: errStr}
	err = fmt.Errorf(errStr)
	return
}

func (s *service) PostOrder(ctx context.Context, request *models.Request) (response models.Response, err error) {
	if request.Id > 0 {
		response.Data = &models.DataStruct{Res: true}
		return
	}
	errStr := "id < 0"
	response = models.Response{Error: true, ErrorText: errStr}
	err = fmt.Errorf(errStr)
	return
}

func (s *service) GetCount(ctx context.Context, request *models.Request) (response models.Response, err error) {
	if request.Id > 0 {
		response.Data = &models.DataStruct{Res: true}
		return
	}
	errStr := "id < 0"
	response = models.Response{Error: true, ErrorText: errStr}
	err = fmt.Errorf(errStr)
	return
}

func (s *service) GetOrder(ctx context.Context) (response models.Response, err error) {
	return
}

func NewService() Service {
	return &service{}
}
