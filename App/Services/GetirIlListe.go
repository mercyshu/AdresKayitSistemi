package Services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type IlListe struct{}

type IlResponseModel struct {
	IlArray []IlResponse
}

type IlResponse struct {
	IlAdi      interface{}
	IlKayitNo  interface{}
	KimlikNo   interface{}
	BilesenAdi interface{}
}

func (Get IlListe) Get(w http.ResponseWriter, r *http.Request) {
	token := GetRequestToken()
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://adres.nvi.gov.tr/Harita/ilListesi", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie", "TS01eef051=0179b2ce4551bafa143b819362c832a387979face466b9655f9ee6dc47c6a3637be108157cf182e5923df95b93f506a79166b6ae558ad45b867681d70fb30de4fd6bb3b265; __RequestVerificationToken=U44le7ehuJY7Lnoje3BVUjLKcD5Y5xBRu2D2l5eB-k_8yka5LTBNyY2kH-ACrA8IousOpeo_GJZdGKTts_uJjnESrd41; _ga=GA1.3.1408641006.1704641143; _gid=GA1.3.1538294979.1704641143; browser-check=%22done%22")
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
	req.Header.Set("Content-Length", "0")
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
		var ilarray []IlResponse
		for _, information := range response {
			IlArray := IlResponse{
				information["adi"].(string),
				information["ilKayitNo"].(float64),
				information["kimlikNo"].(float64),
				information["bilesenAdi"].(string),
			}
			ilarray = append(ilarray, IlArray)
		}

		ResponseModel := IlResponseModel{
			IlArray: ilarray,
		}
		ResponseWriter(w, http.StatusOK, "SUCCESS", "Işleminiz başarılı, 'Il' listesi getirildi!", ResponseModel)
	} else {
		ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'Il' bilgisi bulunamadı!", map[string]interface{}{})
	}
}
