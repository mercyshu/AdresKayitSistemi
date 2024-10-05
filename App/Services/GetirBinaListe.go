package Services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type BinaListe struct{}

type BinaResponseModel struct {
	BinaArray []BinaResponse
}

type BinaResponse struct {
	KimlikNo    interface{}
	BinaAdi     interface{}
	DisKapiNo   interface{}
	DisKapiNo2  interface{}
	CsbmKayitNo interface{}
	AdaNo       interface{}
	ParselNo    interface{}
	PaftaNo     interface{}
	SiteAdi     interface{}
	BlokAdi     interface{}
	BilesenAdi  interface{}
}

func (Get BinaListe) Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	mahKayitNo := params.Get("MahalleKayitNo")
	csbmNo := params.Get("CsbmNo")

	if csbmNo != "" && mahKayitNo != "" {
		token := GetRequestToken()
		client := &http.Client{}
		var data = strings.NewReader(`mahalleKoyBaglisiKimlikNo=` + mahKayitNo + `&yolKimlikNo=` + csbmNo + `&adresReCaptchaResponse=`)
		req, err := http.NewRequest("POST", "https://adres.nvi.gov.tr/Harita/binaListesi", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Set("Cookie", "TS01eef051=0179b2ce4541c587e426ffa74698d1c19875ed8ad19ce2068e6b46c07ead9acecd98eadb5e2a5a79c40ef5db5f6641a7b0961920fb3b8bf0c87d61118e39f2260723c5a8b4; __RequestVerificationToken=U44le7ehuJY7Lnoje3BVUjLKcD5Y5xBRu2D2l5eB-k_8yka5LTBNyY2kH-ACrA8IousOpeo_GJZdGKTts_uJjnESrd41; _ga=GA1.3.1408641006.1704641143; _gid=GA1.3.1538294979.1704641143")
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
		req.Header.Set("Content-Length", "75")
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

		var response []map[string]interface{}
		erreur := json.Unmarshal(bodyText, &response)
		if erreur != nil {
			fmt.Println("Hata Oluştu!", erreur)
		}

		if len(response) > 0 {
			var binaArray []BinaResponse
			for _, information := range response {
				BinaArray := BinaResponse{
					information["kimlikNo"],
					information["adi"],
					information["disKapiNo1"],
					information["disKapiNo2"],
					information["csbmKayitNo"],
					information["ada"],
					information["parsel"],
					information["pafta"],
					information["siteAdi"],
					information["blokAdi"],
					information["bilesenAdi"],
				}
				binaArray = append(binaArray, BinaArray)
			}

			ResponseModel := BinaResponseModel{
				BinaArray: binaArray,
			}
			ResponseWriter(w, http.StatusOK, "SUCCESS", "Işleminiz başarılı, 'Bina' listesi getirildi!", ResponseModel)
		} else {
			ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'MahalleKayitNo & CsbmNo' verisine ait bina bilgisi bulunamadı!", map[string]interface{}{})
		}
	} else {
		ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'MahalleKayitNo & CsbmNo' verisi belirtiniz!", map[string]interface{}{})
	}
}
