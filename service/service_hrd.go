package service

import (
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
	"math/rand"
	"strings"
	"time"
)

type ServiceHrd interface {
	CreateHrd(input dto.CreateHrdDTO) (*entity.Hrd, error)
	GetAllHrd() ([]*entity.Hrd, error)
	DeleteHrd(slug string) (*entity.Hrd, error)
	GetOneHrd(slug string) (*entity.Hrd, error)
	UpdateHrd(input dto.UpdateHrdDTO, slug string) (*entity.Hrd, error)
}

type service_hrd struct {
	repository repository.RepositoryHrd
}

func NewServiceHrd(repository repository.RepositoryHrd) *service_hrd {
	return &service_hrd{repository}
}

func (s *service_hrd) CreateHrd(input dto.CreateHrdDTO) (*entity.Hrd, error) {
	slugName := strings.ReplaceAll(strings.ToLower(input.Nama), " ", "-")
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)

	newHrd := entity.Hrd{
		Nama: input.Nama,
		Departement: input.Departement,
		Role: input.Role,
		Status: input.Status,
		Slug: fmt.Sprintf("%s-%d", slugName, randomNumber),
	}

	hrd, err := s.repository.Save(&newHrd)

	if err != nil {
		return hrd, err
	}

	return hrd, nil
}

func (s *service_hrd) GetAllHrd() ([]*entity.Hrd, error) {
	hrds, err := s.repository.FindAll()

	if err != nil {
		return hrds, err
	}

	return hrds, nil
}

func (s *service_hrd) DeleteHrd(slug string) (*entity.Hrd, error) {
	hrd, err := s.repository.FindBySlug(slug)
	if err != nil {
		return hrd, err
	}

	newHrd, err := s.repository.Delete(hrd)
	if err != nil {
		return newHrd, err
	}

	return newHrd, nil
}

func (s *service_hrd) GetOneHrd(slug string) (*entity.Hrd, error) {
	hrd, err := s.repository.FindBySlug(slug)
	if err != nil {
		return hrd, err
	}

	return hrd, nil
}

func (s *service_hrd) UpdateHrd(input dto.UpdateHrdDTO, slug string) (*entity.Hrd, error) {
	hrd, err := s.repository.FindBySlug(slug)
	if err != nil {
		return hrd, err
	}

	slugName := strings.ReplaceAll(strings.ToLower(input.Nama), " ", "-")
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)

	hrd.Slug = fmt.Sprintf("%s-%d", slugName, randomNumber)
	hrd.Nama = input.Nama
	hrd.Departement = input.Departement
	hrd.Role = input.Role
	hrd.Status = input.Status

	newHrd, err := s.repository.Update(hrd)
	if err != nil {
		return newHrd, err
	}

	return newHrd, nil
}
