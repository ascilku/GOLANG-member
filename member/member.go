package member

import "time"

type Member struct {
	ID        int
	Nama      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (member Member) TableName() string {
	return "member"
}
