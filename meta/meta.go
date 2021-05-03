package meta

import "time"

type Meta struct {
	CreatedAt time.Time  `json:"created_at" sql:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" sql:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"deleted_at"`
}

func New() Meta {
	return Meta{
		CreatedAt: time.Now(),
	}
}
