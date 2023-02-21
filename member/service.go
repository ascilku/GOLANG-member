package member

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SaveService(input InputMember) (Member, error)
	LoginService(login LoginMember) (Member, error)
	CheckEmailIsAvailable(cheack CheckEmailIsAvailable) (bool, error)
	SaveAvatarService(ID int, fileLocation string) (Member, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveService(input InputMember) (Member, error) {
	key_member := Member{}
	key_member.Nama = input.Nama
	key_member.Email = input.Email
	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return key_member, err
	} else {
		key_member.Password = string(pass)
		new_repository, err := s.repository.SaveRepository(key_member)
		if err != nil {
			return new_repository, err
		} else {
			return new_repository, nil
		}
	}
}

func (s *service) LoginService(login LoginMember) (Member, error) {
	email := login.Email
	password := login.Password
	loginRepository, err := s.repository.LoginRepository(email)
	if err != nil {
		return loginRepository, err
	} else {
		if loginRepository.ID == 0 {
			return loginRepository, errors.New("Tidak ada data")
		}
		err := bcrypt.CompareHashAndPassword([]byte(loginRepository.Password), []byte(password))
		if err != nil {
			return loginRepository, err
		} else {
			return loginRepository, nil
		}
	}
}

func (s *service) CheckEmailIsAvailable(cheack CheckEmailIsAvailable) (bool, error) {
	email := cheack.Email
	LoginRepository, err := s.repository.LoginRepository(email)
	if err != nil {
		return false, err
	} else {
		if LoginRepository.ID == 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func (s *service) SaveAvatarService(ID int, fileLocation string) (Member, error) {
	findByIDRepository, err := s.repository.FindByIDRepository(ID)
	if err != nil {
		return findByIDRepository, err
	} else {
		findByIDRepository.Gambar = fileLocation

		updateRepository, err := s.repository.UpdateRepository(findByIDRepository)
		if err != nil {
			return updateRepository, err
		} else {
			return updateRepository, nil
		}
	}
}
