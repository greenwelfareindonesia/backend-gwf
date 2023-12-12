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
	DeleteEcopedia(ID int) (Ecopedia, error)
	UpdateEcopedia(slugs string, input EcopediaInput) (Ecopedia, error)
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



func (s *service) DeleteEcopedia(ID int) (Ecopedia, error) {
	ecopedia, err := s.repository.FindById(ID)
	if err != nil {
		return ecopedia, err
	}

	// Hapus entri terkait di tabel comments terlebih dahulu
	// err = s.repository.DeleteComment(ecopedia.ID)
	// if err != nil {
	// 	return ecopedia, err
	// }
	err = s.repository.DeleteImages(ecopedia.ID)
	if err != nil {
		return ecopedia, err
	}

	// Hapus entri dari tabel ecopedia setelah entri terkait di comments dihapus
	err = s.repository.DeleteEcopedia(ecopedia)
	if err != nil {
		return ecopedia, err
	}

	// Hapus entri terkait di tabel images setelah entri utama dihapus


	return ecopedia, nil
}


func (s *service) UpdateEcopedia(slugs string, input EcopediaInput) (Ecopedia, error) {
	ecopedia, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return ecopedia, nil
	}

	ecopedia.Title = input.Title
	ecopedia.SubTitle = input.SubTitle
	ecopedia.Description = input.Description
	ecopedia.SrcFile = input.SrcFile
	ecopedia.Reference = input.Reference

	oldSlug := ecopedia.Slug
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    slugTitle := strings.ToLower(input.Title)
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

	newEcopedia.Title = ecopedia.Title
	newEcopedia.SubTitle = ecopedia.SubTitle
	newEcopedia.Description = ecopedia.Description
	newEcopedia.SrcFile = ecopedia.SrcFile
	newEcopedia.Reference = ecopedia.Reference
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(ecopedia.Title)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	newEcopedia.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	saveEcopedia, err := s.repository.Create(newEcopedia)

	if err != nil {
		return saveEcopedia, err
	}
	return saveEcopedia, nil
}
