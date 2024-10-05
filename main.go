package main

import (
	Router "AdresKayitSistemi/App/Router"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":1384", Router.AppRoutes())
	if err != nil {
		return
	}
}
