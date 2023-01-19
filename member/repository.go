package member

import "gorm.io/gorm"

type Repository interface {
	SaveRepository(member Member) (Member, error)
	LoginRepository(email string) (Member, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveRepository(member Member) (Member, error) {
	err := r.db.Create(&member).Error
	if err != nil {
		return member, err
	} else {
		return member, nil
	}
}

func (r *repository) LoginRepository(email string) (Member, error) {
	var keyMember Member
	err := r.db.Where("email", email).Find(&keyMember).Error
	if err != nil {
		return keyMember, err
	} else {
		return keyMember, nil
	}
}
