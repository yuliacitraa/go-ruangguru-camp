package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	t.db.Create(&data)
	return nil // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	teacher := []model.Teacher{}
	err := t.db.Unscoped().Find(&teacher).Error
	return teacher, err // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	err := t.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name).Error
	return err // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error { 
	err := t.db.Where("id = ?", id).Delete(&model.Teacher{}).Error
	return err
}
