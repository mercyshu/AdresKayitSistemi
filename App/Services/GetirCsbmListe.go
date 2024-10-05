package Services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type CsbmListe struct{}

type CsbmResponseModel struct {
	CsbmArray []CsbmResponse
}

type CsbmResponse struct {
	MahalleKayitNo interface{}
	CsbmAdi        interface{}
	BilesenAdi     interface{}
	CsbmKimlikNo   interface{}
}

func (Get CsbmListe) Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	mahalleNo := params.Get("MahalleNo")

	if mahalleNo != "" {
		token := GetRequestToken()
		client := &http.Client{}
		var data = strings.NewReader(`mahalleKoyBaglisiKimlikNo=` + mahalleNo + `&adresReCaptchaResponse=`)
		req, err := http.NewRequest("POST", "https://adres.nvi.gov.tr/Harita/yolListesi", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Set("Cookie", "TS01eef051=0179b2ce45ab15897ced0b892ff6ec0b62b3da6c0a3e9a86c85b90167b1674dc6ea9379ecae8d8c3897e9ab689da70ff2f138269c8ec3eb58af0e36f2ca285c5453d3aa27d; __RequestVerificationToken=U44le7ehuJY7Lnoje3BVUjLKcD5Y5xBRu2D2l5eB-k_8yka5LTBNyY2kH-ACrA8IousOpeo_GJZdGKTts_uJjnESrd41; _ga=GA1.3.1408641006.1704641143; _gid=GA1.3.1538294979.1704641143")
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
		req.Header.Set("Content-Length", "54")
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
			var csbmArray []CsbmResponse
			for _, information := range response {
				CsbmArray := CsbmResponse{
					information["mahalleKayitNo"].(float64),
					information["adi"].(string),
					information["bilesenAdi"].(string),
					information["kimlikNo"].(float64),
				}
				csbmArray = append(csbmArray, CsbmArray)
			}

			ResponseModel := CsbmResponseModel{
				CsbmArray: csbmArray,
			}
			ResponseWriter(w, http.StatusOK, "SUCCESS", "Işleminiz başarılı, 'Sokak' listesi getirildi!", ResponseModel)
		} else {
			ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'MahalleNo' verisine ait sokak bilgisi bulunamadı!", map[string]interface{}{})
		}
	} else {
		ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'MahalleNo' verisi belirtiniz!", map[string]interface{}{})
	}
}
