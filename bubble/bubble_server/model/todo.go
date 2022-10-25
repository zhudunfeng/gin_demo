package model

import "go_front/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
//CreateAToDo 创建todo
func CreateATodo(todo *Todo) (err error) {
	if err = dao.DB.Create(todo).Error; err != nil {
		return err
	}
	return
}

func GetAllTodo() (todoList []Todo,err error)  {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetTodoById(id string)(todo Todo,err error){
	err = dao.DB.Where("id=?", id).Find(&todo).Error
	return
}
func UpdateATodo(todo *Todo)(err error)  {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteATodo(id string)(err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
