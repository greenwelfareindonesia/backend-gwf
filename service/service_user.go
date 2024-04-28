package service

import (
	"errors"
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	RegisterUser(input dto.RegisterUserInput) (*entity.User, error)
	Login(input dto.LoginInput) (*entity.User, error)
	IsEmaillAvailabilty(input string) (bool, error)
	GetUserByid(ID int) (*entity.User, error)
	DeleteUser(slug string) (*entity.User, error)
	// SaveAvatar(ID int, fileLocation string) (User, error)
	UpdateUser(slugs string, input dto.UpdateUserInput) (*entity.User, error)
}

type service_user struct {
	repository repository.RepositoryUser
}

func NewServiceUser(repository repository.RepositoryUser) *service_user {
	return &service_user{repository}
}

func (s *service_user) RegisterUser(input dto.RegisterUserInput) (*entity.User, error) {
	user := &entity.User{}

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Username)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	user.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Role = 0
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

func (s *service_user) Login(input dto.LoginInput) (*entity.User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service_user) DeleteUser(slug string) (*entity.User, error) {
	user, err := s.repository.FindBySlug(slug)
	if err != nil {
		return user, err
	}
	userDel, err := s.repository.Delete(user)

	if err != nil {
		return userDel, err
	}
	return userDel, nil
}

func (s *service_user) IsEmaillAvailabilty(input string) (bool, error) {
	//karena hanya email maka di inisiasi hanya email
	emailUser := &entity.User{}
	emailUser.Email = input

	//pengambilan algoritmanya repository yaitu findbyemail
	user, err := s.repository.FindByEmail(input)
	if err != nil {
		return false, err
	}

	// ini nilainya true karena misal kita input email ini sama ga dengan email yang terdaftar dg id sekian
	//kalau g ada maka balikkanya 0 sehingga bisa di daftrakan atau availabilty
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service_user) GetUserByid(ID int) (*entity.User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user Not Found With That ID")
	}

	return user, nil

}

func (s *service_user) UpdateUser(slugs string, input dto.UpdateUserInput) (*entity.User, error) {
	user, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return user, err
	}

	oldSlug := user.Slug

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	slugTitle := strings.ToLower(input.Username)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
	user.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	// Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
	user.Slug = oldSlug

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
