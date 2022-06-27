package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan/akun/connection"
	"perpustakaan/akun/models"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDb()
}

func (c *ctrl) GetData(w http.ResponseWriter, r *http.Request) {
	var listdatadm []models.Admin

	DB.Find(&listdatadm)
	datajson, err := json.Marshal(listdatadm)

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

func (c *ctrl) PostDataAdmin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Admin

	err := decoder.Decode(&datarequest)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	err = c.us.InsertDataAdmin(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Insert Data")
	return
}

func (c *ctrl) UpdateDataAdminn(w http.ResponseWriter, r *http.Request) {
	idadmin := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Admin
	err := decoder.Decode(&datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	idconvert, err := strconv.Atoi(idadmin)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idconvert
	err = c.us.UpdateDataAdmin(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}

func (c *ctrl) DeleteDataAdminn(w http.ResponseWriter, r *http.Request) {
	idAdmin := chi.URLParam(r, "id")

	if idAdmin == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	var datarequest models.Admin
	idConvert, err := strconv.Atoi(idAdmin)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idConvert
	err = c.us.DeleteDataAdmin(datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func (c *ctrl) GetDataAdminDetail(w http.ResponseWriter, r *http.Request) {
	idadmin := chi.URLParam(r, "id")

	idConvert, err := strconv.Atoi(idadmin)

	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	data, err := c.us.GetDataDetailAdmin(idConvert)
	if err != nil {
		ResponseApi(w, 500, nil, err.Error())
	}

	ResponseApi(w, 200, data, "Sukses Get Data")
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {

	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resevice)

}
