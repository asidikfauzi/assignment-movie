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
		Image       string     `gorm:"type:text;not null;" json:"image"`
		CreatedAt   time.Time  `gorm:"default:null" json:"created_at"`
		UpdatedAt   *time.Time `gorm:"default:null" json:"updated_at"`
		DeletedAt   *time.Time `gorm:"default:null" json:"deleted_at"`
	}

	GetMovies struct {
		ID          int        `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Rating      float32    `json:"rating"`
		Image       string     `json:"image"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at"`
	}

	ReqMovie struct {
		Title       string  `form:"title" bindingType:"form" validate:"required"`
		Description string  `form:"description" bindingType:"form" validate:"required"`
		Rating      float32 `form:"rating" bindingType:"form" validate:"required,number"`
		Image       string  `form:"image,omitempty" bindingType:"form"`
	}
)
