package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string
	Image string
}

type CategoryResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Image     string    `json:"string"`
}
