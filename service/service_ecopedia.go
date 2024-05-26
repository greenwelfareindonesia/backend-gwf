package service

import (
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type ServiceEcopedia interface {
	GetAllEcopedia(input int) ([]*entity.Ecopedia, error)
	GetEcopediaByID(slugs string) (*entity.Ecopedia, error)
	CreateEcopedia(ecopedia dto.EcopediaInput) (*entity.Ecopedia, error)
	CreateEcopediaImage(ecopediaID int, FileName string) error
	DeleteEcopedia(ID int) (*entity.Ecopedia, error)
	UpdateEcopedia(slugs string, input dto.EcopediaInput) (*entity.Ecopedia, error)
}

type service_ecopedia struct {
	repository repository.RepositoryEcopedia
}

func NewServiceEcopedia(repository repository.RepositoryEcopedia) *service_ecopedia {
	return &service_ecopedia{repository}
}

func (s *service_ecopedia) CreateEcopediaImage(ecopediaID int, FileName string) error {
	createBerita := &entity.EcopediaImage{}

	createBerita.FileName = FileName
	createBerita.EcopediaID = ecopediaID

	err := s.repository.CreateImage(createBerita)
	if err != nil {
		return err
	}
	return nil
}



func (s *service_ecopedia) DeleteEcopedia(ID int) (*entity.Ecopedia, error) {
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


func (s *service_ecopedia) UpdateEcopedia(slugs string, input dto.EcopediaInput) (*entity.Ecopedia, error) {
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

func (s *service_ecopedia) GetAllEcopedia(input int) ([]*entity.Ecopedia, error) {
	ecopedias, err := s.repository.FindAll()
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
}

func (s *service_ecopedia) GetEcopediaByID(slugs string) (*entity.Ecopedia, error) {
	ecopedias, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return ecopedias, err
	}
	if ecopedias.ID == 0 {
		return ecopedias, err
	}
	return ecopedias, nil
}

func (s *service_ecopedia) CreateEcopedia(ecopedia dto.EcopediaInput) (*entity.Ecopedia, error) {
	newEcopedia := &entity.Ecopedia{}

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
