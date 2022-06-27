package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"perpustakaan/akun/connection"
	"perpustakaan/akun/controllers"
	"perpustakaan/akun/repositories"
	"perpustakaan/akun/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config File !")
	}
}

func main() {
	koneksidb := connection.ConnectToDb()
	repo := repositories.NewRepo(koneksidb)
	usecase := usecase.NewUsecase(repo)
	ctrl := controllers.NewController(usecase)

	r := chi.NewRouter()

	r.Get("/get-data-admin", ctrl.GetData)
	r.Post("/post-data-admin", ctrl.PostDataAdmin)
	r.Put("/update-data-admin/{id}", ctrl.UpdateDataAdminn)
	r.Delete("/delete-data-admin/{id}", ctrl.DeleteDataAdminn)
	r.Get("/get-data-detail-admin/{id}", ctrl.GetDataAdminDetail)

	r.Get("/get-data-pegawai", ctrl.GetDataPegawaii)
	r.Post("/post-data-pegawai", ctrl.InsertDataPegawaii)
	r.Get("/get-detail-pegawai/{id}", ctrl.GetDetailPegawai)
	r.Delete("/delete-data-pegawai/{id}", ctrl.DeleteDataPegawaii)
	r.Put("/update-data-pegawai/{id}", ctrl.UpdateDataPegawaii)

	if err := http.ListenAndServe(":"+os.Getenv("HOST")+"", r); err != nil {
		fmt.Println("Error Starting Service !")
	}
}
