package Router

import (
	"AdresKayitSistemi/App/Services"
	"github.com/gorilla/mux"
	"net/http"
)

func AppRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/AdresKayitSistemi/AcikAdres", Services.AcikAdres{}.Get).Methods("GET")
	router.HandleFunc("/AdresKayitSistemi/GetirIlListe", Services.IlListe{}.Get).Methods("GET")
	router.HandleFunc("/AdresKayitSistemi/GetirIlceListe", Services.IlceListe{}.Get).Methods("GET")
	router.HandleFunc("/AdresKayitSistemi/GetirMahalleListe", Services.MahalleListe{}.Get).Methods("GET")
	router.HandleFunc("/AdresKayitSistemi/GetirCsbmListe", Services.CsbmListe{}.Get).Methods("GET")
	router.HandleFunc("/AdresKayitSistemi/GetirBinaListe", Services.BinaListe{}.Get).Methods("GET")
	router.HandleFunc("/AdresKayitSistemi/GetirBagimsizBolum", Services.BagimsizBolum{}.Get).Methods("GET")

	return router
}
