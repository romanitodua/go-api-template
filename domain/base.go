package domain

import "time"

type BaseDB struct {
	Id        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	CreatedBy int       `json:"-"`
}

type BaseUpdated struct {
	UpdatedBy int       `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Pagination struct {
	Page    int `query:"p" validate:"gte=0"`
	PerPage int `query:"pp" validate:"gt=0"`
}

func DefaultPagination() *Pagination {
	return &Pagination{PerPage: 50}
}
