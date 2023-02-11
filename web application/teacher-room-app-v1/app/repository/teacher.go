package handler

import (
	"a21hc3NpZ25tZW50/app/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (u *TeacherRepo) AddTeacher(teacher model.Teacher) error {
	err := u.db.Create(&teacher).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *TeacherRepo) ReadTeacher() ([]model.ViewTeacher, error) {
	teacher := []model.ViewTeacher{}
	err := u.db.Table("teachers").Select("name","field_of_study","age").Where("deleted_at is NULL").Find(&teacher).Error
	if err != nil {
		return []model.ViewTeacher{}, err
	}
	return teacher, nil
}

func (u *TeacherRepo) UpdateName(id uint, name string) error {
	err := u.db.Model(&model.Teacher{}).Where("id=?", id).Update("name", name).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *TeacherRepo) DeleteTeacher(id uint) error {
	teacher := model.Teacher{}
	err := u.db.Where("id = ?", id).Delete(&teacher).Error
	if err != nil {
		return err
	}
	return nil
}