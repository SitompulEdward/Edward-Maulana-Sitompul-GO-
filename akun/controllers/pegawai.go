package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/akun/connection"
	"perpustakaan/akun/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func init() {
	DB = connection.ConnectToDb()
}

func (c *ctrl) GetDataPegawaii(w http.ResponseWriter, r *http.Request) {
	var listdatpgw []models.Pegawai

	DB.Find(&listdatpgw)
	datajson, err := json.Marshal(listdatpgw)

	if err != nil {
		w.Write([]byte("Error Convert TO JSON"))
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(datajson)
	w.WriteHeader(200)
	return
}

func (c *ctrl) InsertDataPegawaii(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Pegawai

	err := decoder.Decode(&datarequest)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	err = c.us.InsertDataPegawai(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Insert Data")
	return
}

func (c *ctrl) GetDetailPegawai(w http.ResponseWriter, r *http.Request) {
	idpgw := chi.URLParam(r, "id")

	idConvert, err := strconv.Atoi(idpgw)

	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	data, err := c.us.SelectDetailPegawai(idConvert)
	if err != nil {
		ResponseApi(w, 500, nil, err.Error())
	}

	ResponseApi(w, 200, data, "Sukses Get Data")
}

func (c *ctrl) DeleteDataPegawaii(w http.ResponseWriter, r *http.Request) {
	idpgw := chi.URLParam(r, "id")

	if idpgw == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	var datarequest models.Pegawai
	idConvert, err := strconv.Atoi(idpgw)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idConvert
	err = c.us.DeleteDataPegawai(datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func (c *ctrl) UpdateDataPegawaii(w http.ResponseWriter, r *http.Request) {
	idpgw := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Pegawai
	err := decoder.Decode(&datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	idconvert, err := strconv.Atoi(idpgw)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idconvert
	err = c.us.UpdateDataPegawai(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}
