package service

import "go_front/model"

type Service interface {
	GetTodoById(id string) (todo model.Todo,err error)
}
