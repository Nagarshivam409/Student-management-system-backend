package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Teacher struct {
	gorm.Model

	Name string `json:"name"`
	TID  int    `gorm:"unique_index primary key ;autoIncrement" json:"tid"`
}
type AttendanceTeacher struct {
	gorm.Model

	Month    time.Month `json:"month"`
	Year     int        `json:"year"`
	Day      int        `json:"day"`
	TID      int        `gorm:"unique_index" json:"tid"`
	PunchIn  string     `json:"punchin"`
	PunchOut string     `json:"punchout"`
}
