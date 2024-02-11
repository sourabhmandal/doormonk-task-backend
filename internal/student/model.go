package student

import (
	"time"
)

type Student struct {
	ID              int64   `json:"id"`
	FirstName       string  `json:"first_name" binding:"required,max=20"`
	LastName        string  `json:"last_name" binding:"omitempty,max=20"`
	Email           string  `json:"email" binding:"required,email"`
	PhoneNumber     string  `json:"phone_number" binding:"omitempty,max=15,phone"`
	InstitutionName string  `json:"institution_name" binding:"required,max=50"`
	PassOutYear     int32   `json:"pass_out_year" binding:"required,gte=1900,lte=now"`
	CgpiScore       float32 `json:"cgpi_score" binding:"required,gte=0,lte=10"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
