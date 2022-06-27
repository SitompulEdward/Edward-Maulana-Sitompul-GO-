package models

type Pegawai struct {
	Id       int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	No_telp  string `json:"no_telp"`
}
