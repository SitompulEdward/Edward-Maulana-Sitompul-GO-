package repositories

import (
	"perpustakaan/akun/models"

	"gorm.io/gorm"
)

type repo struct {
	apps *gorm.DB
}

type Repo interface {
	SelectAll(i interface{}) error
	InserData(i interface{}) error
	UpdateData(i models.Admin) error
	DeleteData(i models.Admin) error
	SelectAdminById(id int) (data models.Admin, err error)

	SelectAllPegawai(i interface{}) error
	InsertDataPegawai(i interface{}) error
	SelectPegawaiById(id int) (data models.Pegawai, err error)
	DeleteDataPegawai(i models.Pegawai) error
	UpdateDataPegawai(i models.Pegawai) error
}

func NewRepo(apps *gorm.DB) Repo {
	return &repo{apps}
}
