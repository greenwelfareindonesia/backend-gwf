package ecopedia

import "errors"

type Service interface {
	GetAllEcopedia(input int) ([]Ecopedia, error)
	GetEcopediaByID(id int) (Ecopedia, error)
	CreateEcopedia(ecopedia EcopediaInput, FileName string) (Ecopedia, error)
	DeleteEcopedia(ID int) (Ecopedia, error)
	UpdateEcopedia (getIdEcopedia EcopediaID, input EcopediaInput, FileName string) (Ecopedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteEcopedia(ID int) (Ecopedia, error) {
	ecopedias, err := s.repository.FindById(ID)
	if err != nil {
		return ecopedias, err
	}

	ecopedia, err := s.repository.DeleteEcopedia(ecopedias)
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}

func (s *service) UpdateEcopedia (getIdEcopedia EcopediaID, input EcopediaInput, FileName string) (Ecopedia, error) {
	ecopedia, err := s.repository.FindById(getIdEcopedia.ID)
	if err != nil { 
		return ecopedia, nil
	}

	ecopedia.Judul = input.Judul
	ecopedia.Subjudul = input.Subjudul
	ecopedia.Deskripsi = input.Deskripsi
	ecopedia.Gambar = FileName
	ecopedia.Srcgambar = input.Srcgambar
	ecopedia.Referensi = input.Referensi

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

func (s *service) GetEcopediaByID(id int) (Ecopedia, error) {
	ecopedias, err := s.repository.FindById(id)
	if err != nil {
		return ecopedias, errors.New("found nothing")
	}
	return ecopedias, nil
}

func (s *service) CreateEcopedia(ecopedia EcopediaInput, FileName string) (Ecopedia, error) {
	newEcopedia := Ecopedia{}

	newEcopedia.Judul = ecopedia.Judul
	newEcopedia.Subjudul = ecopedia.Subjudul
	newEcopedia.Deskripsi = ecopedia.Deskripsi
	// newEcopedia.Gambar = ecopedia.Gambar
	newEcopedia.Gambar = FileName
	newEcopedia.Srcgambar = ecopedia.Srcgambar
	newEcopedia.Referensi = ecopedia.Referensi

	saveEcopedia, err := s.repository.Create(newEcopedia)

	if err != nil {
		return saveEcopedia, err
	}
	return saveEcopedia, nil
}
