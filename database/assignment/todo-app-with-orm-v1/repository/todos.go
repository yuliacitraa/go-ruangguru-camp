package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (u *TodoRepository) AddTodo(todo model.Todo) error {
	err := u.db.Create(&todo).Error
	return err
}

func (u *TodoRepository) ReadTodo() ([]model.Todo, error) {
	todo := []model.Todo{}
	err := u.db.Unscoped().Find(&todo).Error
	return todo, err
}

func (u *TodoRepository) UpdateDone(id uint, status bool) error {
	err := u.db.Model(&model.Todo{}).Where("id = ?", id).Update("done", status).Error
	return err
}

func (u *TodoRepository) DeleteTodo(id uint) error {
	err := u.db.Where("id = ?", id).Delete(&model.Todo{}).Error
	return err
}
