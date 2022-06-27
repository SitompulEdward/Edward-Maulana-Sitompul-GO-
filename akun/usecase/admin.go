package usecase

import (
	"errors"
	"perpustakaan/akun/models"
)

func (r *UC) GetDataAdmin() ([]models.Admin, error) {
	var modelAdmin []models.Admin

	err := r.queryrepo.SelectAll(modelAdmin)
	if err != nil {
		return modelAdmin, err
	}

	return modelAdmin, nil
}

func (r *UC) InsertDataAdmin(data models.Admin) error {
	err := r.queryrepo.InserData(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) UpdateDataAdmin(data models.Admin) error {
	admin, err := r.queryrepo.SelectAdminById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	admin.Nama = data.Nama
	admin.No_telp = data.No_telp
	admin.Email = data.Email
	admin.Password = data.Password

	err = r.queryrepo.UpdateData(admin)
	if err != nil {
		return err
	}

	return nil
}

func (r *UC) DeleteDataAdmin(data models.Admin) error {
	admin, err := r.queryrepo.SelectAdminById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	admin.Id = data.Id
	admin.Nama = data.Nama
	admin.No_telp = data.No_telp
	admin.Email = data.Email
	admin.Password = data.Password

	err = r.queryrepo.DeleteData(admin)
	if err != nil {
		return err
	}

	return nil
}

func (r *UC) GetDataDetailAdmin(id int) (data models.Admin, err error) {
	data, err = r.queryrepo.SelectAdminById(id)
	if err != nil {
		return data, errors.New("data tidak ada")
	}

	return data, nil
}
