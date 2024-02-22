package user

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmaillAvailabilty(input string) (bool, error)
	GetUserByid(ID int) (User, error)
	DeleteUser(slug string) (User, error)
	// SaveAvatar(ID int, fileLocation string) (User, error)
	UpdateUser(slugs string, input UpdateUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

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

func (s *service) DeleteUser(slug string) (User, error) {
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

func (s *service) IsEmaillAvailabilty(input string) (bool, error) {
	//karena hanya email maka di inisiasi hanya email
	emailUser := User{}
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

func (s *service) UpdateUser(slugs string, input UpdateUserInput) (User, error) {
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
