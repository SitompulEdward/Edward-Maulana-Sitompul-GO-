package controllers

import (
	"net/http"
	"perpustakaan/akun/usecase"
)

type ctrl struct {
	us usecase.Usecase
}

type ControllerInterface interface {
	GetData(w http.ResponseWriter, r *http.Request)
	PostDataAdmin(w http.ResponseWriter, r *http.Request)
	UpdateDataAdminn(w http.ResponseWriter, r *http.Request)
	DeleteDataAdminn(w http.ResponseWriter, r *http.Request)
	GetDataAdminDetail(w http.ResponseWriter, r *http.Request)

	GetDataPegawaii(w http.ResponseWriter, r *http.Request)
	InsertDataPegawaii(w http.ResponseWriter, r *http.Request)
	GetDetailPegawai(w http.ResponseWriter, r *http.Request)
	DeleteDataPegawaii(w http.ResponseWriter, r *http.Request)
	UpdateDataPegawaii(w http.ResponseWriter, r *http.Request)
}

func NewController(us usecase.Usecase) ControllerInterface {
	return &ctrl{us: us}
}
