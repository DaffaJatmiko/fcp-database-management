package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	if result := s.db.Create(&session); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	result := s.db.Table("sessions").Where("token = ?", token).Delete(&model.Session{})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Table("sessions").Where("username = ?", session.Username).Updates(model.Session{
		Token:    session.Token,
		Username: session.Username,
		Expiry:   session.Expiry,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	result := s.db.Table("sessions").Where("username = ?", name).Find(&model.Session{})
	if result.RowsAffected == 0 {
		return errors.New("session not found")
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := s.db.Table("sessions").Select("*").Where("token = ?", token).First(&session).Error
	if err != nil {
		return model.Session{}, err
	}
	return session, nil
	
	// row, err := s.db.Table("sessions").Select("*").Where("token = ?", token).Rows()
	// if err != nil {
	// 	return model.Session{}, err
	// }

	// var result model.Session
	// for row.Next() {
	// 	err := s.db.ScanRows(row, &result)
	// 	if err != nil {
	// 		return model.Session{}, err
	// 	}
	// }

	// return result, nil // TODO: replace this
}
