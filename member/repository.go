package member

import "gorm.io/gorm"

type Repository interface {
	SaveRepository(member Member) (Member, error)
	LoginRepository(email string) (Member, error)
	FindByIDRepository(ID int) (Member, error)
	UpdateRepository(member Member) (Member, error)
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

func (r *repository) FindByIDRepository(ID int) (Member, error) {
	var member Member

	err := r.db.Where("id = ?", ID).Find(&member).Error
	if err != nil {
		return member, err
	} else {
		return member, nil
	}
}

func (r *repository) UpdateRepository(member Member) (Member, error) {
	err := r.db.Save(&member).Error
	if err != nil {
		return member, err
	} else {
		return member, nil
	}
}
