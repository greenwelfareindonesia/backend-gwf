package service

import (
	"errors"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"

	"github.com/gosimple/slug"
)

type ServiceKajian interface {
	Create(input dto.InputKajian) (*entity.Kajian, error)
	CreateImage(kajianID int, fileLocation string) error
	GetAll() ([]*entity.Kajian, error)
	GetOne(slug string) (*entity.Kajian, error)
	UpdateOne(slug string, update dto.UpdateKajian) (*entity.Kajian, error)
	DeleteOne(slug string) error
}

type service_kajian struct {
	repository repository.RepositoryKajian
}

func NewServiceKajian(repository repository.RepositoryKajian) *service_kajian {
	return &service_kajian{repository}
}

func (s *service_kajian) Create(input dto.InputKajian) (*entity.Kajian, error) {
	kajian := &entity.Kajian{
		Slug:        slug.Make(input.Title),
		Title:       input.Title,
		Description: input.Description,
	}

	newKajian, err := s.repository.Save(kajian)

	if err != nil {
		return nil, err
	}

	return newKajian, nil
}

func (s *service_kajian) CreateImage(kajianID int, fileLocation string) error {
	image := &entity.KajianImage{
		FileName: fileLocation,
		KajianID: kajianID,
	}

	if err := s.repository.SaveImage(image); err != nil {
		return err
	}

	return nil
}

func (s *service_kajian) GetAll() ([]*entity.Kajian, error) {
	return s.repository.FindAll()
}

func (s *service_kajian) GetOne(slug string) (*entity.Kajian, error) {
	kajian, err := s.repository.FindBySlug(slug)

	if err != nil {
		return nil, err
	}

	if kajian.ID == 0 {
		return nil,	errors.New("kajian not found!")
	}

	return kajian, nil
}

func (s *service_kajian) UpdateOne(slug string, update dto.UpdateKajian) (*entity.Kajian, error) {
	kajian, err := s.repository.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	if kajian.ID == 0 {
		return nil, errors.New("hrd not found")
	}

	kajian.Title = update.Title
	kajian.Description = update.Description

	newHrd, err := s.repository.Update(kajian)
	if err != nil {
		return newHrd, err
	}

	return newHrd, nil
}

func (s *service_kajian) DeleteOne(slug string) error {
	kajian, err := s.repository.FindBySlug(slug)
	if err != nil {
		return err
	}

	if err := s.repository.Delete(kajian); err != nil {
		return err
	}

	return nil
}
