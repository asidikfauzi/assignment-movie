package models

import (
	"time"
)

type (
	Movies struct {
		ID          int        `gorm:"type:int;unique;primary_key;auto_increment;column:id" json:"id"`
		Title       string     `gorm:"type:varchar(150);not null;" json:"title"`
		Description string     `gorm:"type:text;not null;" json:"description"`
		Rating      float32    `gorm:"type:float" json:"rating"`
		Image       string     `gorm:"type:varchar(50);not null;" json:"image"`
		CreatedAt   time.Time  `gorm:"default:null" json:"created_at"`
		UpdatedAt   *time.Time `gorm:"default:null" json:"updated_at"`
		DeletedAt   *time.Time `gorm:"default:null" json:"deleted_at"`
	}
)
