package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmaillAvailabilty(input CheckEmailInput) (bool, error)
	GetUserByid(ID int) (User, error)
	DeleteUser(ID int) (User, error)
	// SaveAvatar(ID int, fileLocation string) (User, error)
	UpdateUser(InputID IdUser, input UpdateUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User not found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) DeleteUser(userID int) (User, error) {
	user, err := s.repository.FindById(userID)
	if err != nil {
		return user, err
	}
	userDel, err := s.repository.Delete(user)

	if err != nil {
		return userDel, err
	}
	return userDel, nil
}

func (s *service) IsEmaillAvailabilty(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByid(ID int) (User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found With That ID")
	}

	return user, nil

}

// func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
// 	//mengambil repository findbyid karena id mana ni yang mau upload avatar
// 	user, err := s.repository.FindById(ID)

// 	if err != nil {
// 		return user, err
// 	}

// 	//lalu ini adalah nama filenya apa disimpan dalam parameter
// 	user.Avatar_file_name = fileLocation

// 	//ini ambil dalam algonya repository kalau mau diupdate
// 	updatedUser, err := s.repository.Update(user)
// 	if err != nil {
// 		return updatedUser, err
// 	}

// 	return updatedUser, nil
// }

func (s *service) UpdateUser(InputID IdUser, input UpdateUserInput) (User, error) {
	user, err := s.repository.FindById(InputID.ID)
	if err != nil {
		return user, err
	}
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Update(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
