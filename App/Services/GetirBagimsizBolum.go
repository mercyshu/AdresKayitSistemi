package Services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type BagimsizBolum struct{}

type BagimsizBolumResponseModel struct {
	BagimsizBolumArray []BagimsizBolumResponse
}

type BagimsizBolumResponse struct {
	AdresNo          interface{}
	IcKapiNo         interface{}
	DisKapiNo        interface{}
	YapiKullanimAmac interface{}
	BinaNo           interface{}
	BinaKayitNo      interface{}
	BlokAdi          interface{}
	SiteAdi          interface{}
	AdaNo            interface{}
	ParselNo         interface{}
	PaftaNo          interface{}
	BilesenAdi       interface{}
}

func (Get BagimsizBolum) Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	mahKayitNo := params.Get("MahalleKayitNo")
	binaNo := params.Get("BinaKimlikNo")

	if binaNo != "" && mahKayitNo != "" {
		token := GetRequestToken()
		client := &http.Client{}
		var data = strings.NewReader(`mahalleKoyBaglisiKimlikNo=` + mahKayitNo + `&binaKimlikNo=` + binaNo + `&adresReCaptchaResponse=`)
		req, err := http.NewRequest("POST", "https://adres.nvi.gov.tr/Harita/bagimsizBolumListesi", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Set("Cookie", "TS01eef051=0179b2ce45ee1942e7dd612b915328eef3e5daac590c35b75d66fad56c90b0e4239f6206ef8015dd8aadc9bf2e6ad0900109c39ca9ac9660329433d0fefcab92f590c9857b; __RequestVerificationToken=U44le7ehuJY7Lnoje3BVUjLKcD5Y5xBRu2D2l5eB-k_8yka5LTBNyY2kH-ACrA8IousOpeo_GJZdGKTts_uJjnESrd41; _ga=GA1.3.1408641006.1704641143; _gid=GA1.3.1538294979.1704641143")
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
		req.Header.Set("Content-Length", "79")
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
			var bbArray []BagimsizBolumResponse
			for _, information := range response {
				BbArray := BagimsizBolumResponse{
					information["adresNo"],
					information["icKapiNo"],
					information["disKapiNo"],
					information["yapiKullanimAmacFormatted"],
					information["binaNo"],
					information["binaKayitNo"],
					information["blokAdi"],
					information["siteAdi"],
					information["ada"],
					information["parsel"],
					information["pafta"],
					information["bilesenAdi"],
				}
				bbArray = append(bbArray, BbArray)
			}

			ResponseModel := BagimsizBolumResponseModel{
				BagimsizBolumArray: bbArray,
			}
			ResponseWriter(w, http.StatusOK, "SUCCESS", "Işleminiz başarılı, 'Bağımsız Bölüm' listesi getirildi!", ResponseModel)
		} else {
			ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'MahalleKayitNo & BinaKimlikNo' verisine ait bağımsız bölüm bilgisi bulunamadı!", map[string]interface{}{})
		}
	} else {
		ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'MahalleKayitNo & BinaKimlikNo' verisi belirtiniz!", map[string]interface{}{})
	}
}
