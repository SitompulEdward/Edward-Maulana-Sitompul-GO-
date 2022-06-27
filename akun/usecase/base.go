package usecase

import (
	"perpustakaan/akun/models"
	"perpustakaan/akun/repositories"
)

type UC struct {
	queryrepo repositories.Repo
}

type Usecase interface {
	GetDataAdmin() ([]models.Admin, error)
	InsertDataAdmin(models.Admin) error
	UpdateDataAdmin(models.Admin) error
	DeleteDataAdmin(models.Admin) error
	GetDataDetailAdmin(id int) (data models.Admin, err error)

	GetDataPegawai() ([]models.Pegawai, error)
	InsertDataPegawai(models.Pegawai) error
	SelectDetailPegawai(id int) (data models.Pegawai, err error)
	DeleteDataPegawai(models.Pegawai) error
	UpdateDataPegawai(models.Pegawai) error
}

func NewUsecase(r repositories.Repo) Usecase {
	return &UC{queryrepo: r}
}
