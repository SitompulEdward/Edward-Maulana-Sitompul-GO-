package usecase

import (
	"errors"
	"perpustakaan/akun/models"
)

func (r *UC) GetDataPegawai() ([]models.Pegawai, error) {
	var Modelpgw []models.Pegawai

	err := r.queryrepo.SelectAllPegawai(Modelpgw)
	if err != nil {
		return Modelpgw, err
	}

	return Modelpgw, nil
}

func (r *UC) InsertDataPegawai(data models.Pegawai) error {
	err := r.queryrepo.InsertDataPegawai(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) SelectDetailPegawai(id int) (data models.Pegawai, err error) {
	data, err = r.queryrepo.SelectPegawaiById(id)
	if err != nil {
		return data, errors.New("data tidak ada")
	}

	return data, nil
}

func (r *UC) DeleteDataPegawai(data models.Pegawai) error {
	pegawai, err := r.queryrepo.SelectPegawaiById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	pegawai.Id = data.Id
	pegawai.Nama = data.Nama
	pegawai.No_telp = data.No_telp
	pegawai.Email = data.Email
	pegawai.Password = data.Password

	err = r.queryrepo.DeleteDataPegawai(pegawai)
	if err != nil {
		return err
	}

	return nil
}

func (r *UC) UpdateDataPegawai(data models.Pegawai) error {
	pegawai, err := r.queryrepo.SelectPegawaiById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	pegawai.Nama = data.Nama
	pegawai.No_telp = data.No_telp
	pegawai.Email = data.Email
	pegawai.Password = data.Password

	err = r.queryrepo.UpdateDataPegawai(pegawai)
	if err != nil {
		return err
	}

	return nil
}
