package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {
	rows, err := s.db.Table("classes").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var students []model.Class
	for rows.Next() {
		var student model.Class
		err := s.db.ScanRows(rows, &student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil // TODO: replace this
}
