package ecopedia

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetAllEcopedia(input int) ([]Ecopedia, error)
	GetEcopediaByID(slugs string) (Ecopedia, error)
	CreateEcopedia(ecopedia EcopediaInput) (Ecopedia, error)
	CreateEcopediaImage(ecopediaID int, FileName string) error
	DeleteEcopedia(slugs string) (Ecopedia, error)
	UpdateEcopedia(slugs string, input EcopediaInput) (Ecopedia, error)
	UserActionToEcopedia(getIdEcopedia EcopediaID, inputUser UserActionToEcopedia) (Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateEcopediaImage(ecopediaID int, FileName string) error {
	createBerita := EcopediaImage{}

	createBerita.FileName = FileName
	createBerita.EcopediaID = ecopediaID

	err := s.repository.CreateImage(createBerita)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UserActionToEcopedia(getIdEcopedia EcopediaID, inputUser UserActionToEcopedia) (Comment, error) {

	FindEcoId, err := s.repository.FindEcopediaCommentID(getIdEcopedia.ID)
	if err != nil {
		return FindEcoId, err
	}

	createComment := Comment{}

	createComment.Comment = inputUser.Comment
	createComment.UserId = inputUser.User.ID
	createComment.EcopediaId = getIdEcopedia.ID
	// FindUserId, err := s.repository.FindByUserId(FindEcoId.User.ID)
	// if err != nil {
	// 	return FindUserId, err
	// }

	create, err := s.repository.CreateComment(createComment)
	if err != nil {
		return create, err
	}
	return create, nil

}

func (s *service) DeleteEcopedia(slugs string) (Ecopedia, error) {
	ecopedias, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return ecopedias, err
	}

	ecopedia, err := s.repository.DeleteEcopedia(ecopedias)
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}

func (s *service) UpdateEcopedia(slugs string, input EcopediaInput) (Ecopedia, error) {
	ecopedia, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return ecopedia, nil
	}

	ecopedia.Judul = input.Judul
	ecopedia.Subjudul = input.Subjudul
	ecopedia.Deskripsi = input.Deskripsi
	ecopedia.Srcgambar = input.Srcgambar
	ecopedia.Referensi = input.Referensi

	oldSlug := ecopedia.Slug
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    slugTitle := strings.ToLower(input.Judul)
    mySlug := slug.Make(slugTitle)
    randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
    ecopedia.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

    // Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
    ecopedia.Slug = oldSlug

	newEcopedia, err := s.repository.Update(ecopedia)
	if err != nil {
		return newEcopedia, err
	}
	return newEcopedia, nil
}

func (s *service) GetAllEcopedia(input int) ([]Ecopedia, error) {
	ecopedias, err := s.repository.FindAll()
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
}

func (s *service) GetEcopediaByID(slugs string) (Ecopedia, error) {
	ecopedias, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return ecopedias, err
	}
	if ecopedias.ID == 0 {
		return ecopedias, err
	}
	return ecopedias, nil
}

func (s *service) CreateEcopedia(ecopedia EcopediaInput) (Ecopedia, error) {
	newEcopedia := Ecopedia{}

	newEcopedia.Judul = ecopedia.Judul
	newEcopedia.Subjudul = ecopedia.Subjudul
	newEcopedia.Deskripsi = ecopedia.Deskripsi
	// newEcopedia.Gambar = ecopedia.Gambar
	newEcopedia.Srcgambar = ecopedia.Srcgambar
	newEcopedia.Referensi = ecopedia.Referensi
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(ecopedia.Judul)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	newEcopedia.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	saveEcopedia, err := s.repository.Create(newEcopedia)

	if err != nil {
		return saveEcopedia, err
	}
	return saveEcopedia, nil
}
