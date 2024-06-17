package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Student struct {
	Name    string  `json:"name"`
	Marks   int     `json:"marks"`
	Address Address `json:"address" gorm:"embedded"`
}

type Class struct {
	ClassName string    `json:"class_name"`
	Students  []Student `json:"students"`
}

type Classes struct {
	Class6 Class `json:"class_6" gorm:"embedded;embeddedPrefix:class6_"`
	Class7 Class `json:"class_7" gorm:"embedded;embeddedPrefix:class7_"`
}

type School struct {
	gorm.Model
	Name        string  `json:"name"`
	SchoolID    string  `gorm:"unique;not null" json:"school_id"`
	Classes     Classes `gorm:"-"` // Ignore this field in GORM
	ClassesJSON string  `json:"-"` // Store serialized JSON
}

func (s *School) BeforeSave(tx *gorm.DB) (err error) {
	return s.serializeClasses()
}

func (s *School) AfterFind(tx *gorm.DB) (err error) {
	return s.deserializeClasses()
}

func (s *School) serializeClasses() error {
	classesJSON, err := json.Marshal(s.Classes)
	if err != nil {
		return err
	}
	s.ClassesJSON = string(classesJSON)
	return nil
}

func (s *School) deserializeClasses() error {
	return json.Unmarshal([]byte(s.ClassesJSON), &s.Classes)
}
