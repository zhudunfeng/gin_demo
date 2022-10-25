package service

import (
	"go_front/model"
)

type ServiceImpl struct {

}

func (s *ServiceImpl) GetTodoById(id string) (todo model.Todo,err error)  {
	todo, err = model.GetTodoById(id)
	return
}
