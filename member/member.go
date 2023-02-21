package member

import "time"

type Member struct {
	ID        int
	Nama      string
	Email     string
	Password  string
	Gambar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (member Member) TableName() string {
	return "member"
}
