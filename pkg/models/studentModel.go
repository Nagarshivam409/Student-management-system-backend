package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model

	Name  string `json:"name"`
	Class string `json:"class"`
	SID   int    `gorm:"unique_index primary key;autoIncrement" json:"sid" `
}
type AttendanceStudent struct {
	gorm.Model

	SID      int        `gorm:"unique_index " json:"sid" `
	Month    time.Month `json:"month"`
	Year     int        `json:"year"`
	Day      int        `json:"day"`
	Class    string     `json:"class"`
	PunchIn  string     `json:"punchin"`
	PunchOut string     `json:"punchout"`
}
