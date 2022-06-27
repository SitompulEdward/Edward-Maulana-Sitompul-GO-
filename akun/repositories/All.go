package repositories

import (
	"perpustakaan/akun/models"
)

func (r *repo) SelectAll(i interface{}) error {
	result := r.apps.Find(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) SelectAdminById(id int) (data models.Admin, err error) {
	err = r.apps.First(&data, id).Error
	return
}

func (r *repo) InserData(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) UpdateData(i models.Admin) error {
	result := r.apps.Save(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) DeleteData(i models.Admin) error {
	result := r.apps.Delete(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//__PEGAWAI__

func (r *repo) SelectAllPegawai(i interface{}) error {
	result := r.apps.Find(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) InsertDataPegawai(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) SelectPegawaiById(id int) (data models.Pegawai, err error) {
	err = r.apps.First(&data, id).Error
	return
}

func (r *repo) DeleteDataPegawai(i models.Pegawai) error {
	result := r.apps.Delete(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) UpdateDataPegawai(i models.Pegawai) error {
	result := r.apps.Save(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
