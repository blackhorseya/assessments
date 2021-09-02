package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ============= DO NOT EDIT =============

// var db gorm.DB         // connected db client
// var cache redis.Client // connected redis client
// var logger zap.Logger  // some useful logger for all developers in company

type MessageName string

const (
	TeacherCreated MessageName = "teacher_created"
)

type Teacher struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type MessageQueue interface {
	Publish(MessageName, map[string]interface{}) error
}

type TeacherDataAccessObject interface {
	CreateTeacher(string) (*Teacher, error)
}

type TeacherDAO struct {
	mq MessageQueue
	db *gorm.DB
}

func NewTeacherDAO(mqClient MessageQueue, dbClient *gorm.DB) TeacherDataAccessObject {
	return TeacherDAO{
		mq: mqClient,
		db: dbClient,
	}
}

func (dao TeacherDAO) CreateTeacher(name string) (*Teacher, error) {

	teacher := Teacher{Name: name}

	r := dao.db.Create(teacher)

	if r.Error != nil {
		return nil, r.Error
	}

	payload := map[string]interface{}{
		"teacherId": teacher.ID,
	}

	err := dao.mq.Publish(TeacherCreated, payload)

	if err != nil {
		// use logger here
	}

	return &teacher, nil
}
