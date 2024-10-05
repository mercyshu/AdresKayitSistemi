package Services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type AcikAdres struct{}

type AcikAdresResponse struct {
	AdresNo          interface{}
	AcikAdres        interface{}
	IlAdi            interface{}
	IlceAdi          interface{}
	IcKapiNo         interface{}
	DisKapiNo        interface{}
	DisKapiNo2       interface{}
	YapiKullanimAmac interface{}
	KayitNo          interface{}
	BinaNo           interface{}
	BinaKayitNo      interface{}
	BlokAdi          interface{}
	SiteAdi          interface{}
	AdaNo            interface{}
	ParselNo         interface{}
	PaftaNo          interface{}
	CsbmAdi          interface{}
	CsbmKayitNo      interface{}
	MahalleAdi       interface{}
	MahalleKayitNo   interface{}
	KoyAdi           interface{}
	KoyKayitNo       interface{}
	BilesenAdi       interface{}
}

func (Get AcikAdres) Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	adresNo := params.Get("AdresNo")

	if adresNo != "" {
		token := GetRequestToken()
		client := &http.Client{}
		var data = strings.NewReader(`bagimsizBolumKayitNo=&bagimsizBolumAdresNo=` + adresNo)
		req, err := http.NewRequest("POST", "https://adres.nvi.gov.tr/Harita/AcikAdres", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Set("Cookie", "TS01eef051=0179b2ce4511c15dbcdda2a85496d33d55c549823b4e32b0c0e3b558d5a3f934a6c5ede6e1e8e649d07e27fa248ceed215f3c6cfbc4173184ea14ea8c8fad6615f5efb0699; __RequestVerificationToken=U44le7ehuJY7Lnoje3BVUjLKcD5Y5xBRu2D2l5eB-k_8yka5LTBNyY2kH-ACrA8IousOpeo_GJZdGKTts_uJjnESrd41; _ga=GA1.3.1408641006.1704641143; _gid=GA1.3.1538294979.1704641143; browser-check=%22done%22")
		req.Header.Set("Pragma", "no-cache")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("Host", "adres.nvi.gov.tr")
		req.Header.Set("Accept-Language", "tr-TR,tr;q=0.9")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Origin", "https://adres.nvi.gov.tr")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15")
		req.Header.Set("Referer", "https://adres.nvi.gov.tr/VatandasIslemleri/AdresSorgu")
		req.Header.Set("Content-Length", "50")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("__RequestVerificationToken", token)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var response map[string]interface{}
		erreur := json.Unmarshal(bodyText, &response)
		if erreur != nil {
			fmt.Println("Hata Oluştu!")
		}

		if response != nil {
			adressNo, ok := response["adresNo"].(float64)
			acikAdresModel, ok := response["acikAdresModel"].(map[string]interface{})
			if ok {
				if adressNo != 0 {
					ResponseModel := AcikAdresResponse{
						response["adresNo"],
						acikAdresModel["acikAdresAciklama"],
						acikAdresModel["ilAdi"],
						acikAdresModel["ilceAdi"],
						acikAdresModel["icKapiNo"],
						acikAdresModel["disKapiNo1"],
						acikAdresModel["disKapiNo2"],
						response["yapiKullanimAmacFormatted"],
						acikAdresModel["bagimsizBolumKayitNo"],
						response["binaNo"],
						response["binaKayitNo"],
						acikAdresModel["blokAdi"],
						acikAdresModel["siteAdi"],
						acikAdresModel["ada"],
						acikAdresModel["pafta"],
						acikAdresModel["parsel"],
						acikAdresModel["csbmAdi"],
						acikAdresModel["csbmKayitNo"],
						acikAdresModel["mahalleAdi"],
						acikAdresModel["mahalleKayitNo"],
						acikAdresModel["koyAdi"],
						acikAdresModel["koyKayitNo"],
						response["bilesenAdi"],
					}

					ResponseWriter(
						w, http.StatusOK,
						"SUCCESS", "Işleminiz başarılı, 'AdresNo' ile ilişkin veriler getirildi!", ResponseModel)
				} else {
					ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'AdresNo' verisine ait sonuç bulunamadı!", map[string]interface{}{})
				}
			} else {
				ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'AdresNo' verisine ait sonuç bulunamadı!", map[string]interface{}{})
			}
		}
	} else {
		ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'AdresNo' verisi göndermelisiniz!", map[string]interface{}{})
	}
}
