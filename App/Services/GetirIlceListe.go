package Services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IlceListe struct{}

type IlceResponseModel struct {
	IlceArray []Ilce
}
type Ilce struct {
	IlAdi    string `json:"sehir_adi"`
	IlKodu   string `json:"sehir_id"`
	IlceAdi  string `json:"ilce_adi"`
	IlceKodu string `json:"ilce_id"`
}

func (Get IlceListe) Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ilNo := params.Get("IlNo")

	if ilNo != "" {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/metinyildirimnet/turkiye-adresler-json/main/ilceler.json", nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Set("Sec-Fetch-Site", "none")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Host", "raw.githubusercontent.com")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15")
		req.Header.Set("Accept-Language", "tr-TR,tr;q=0.9")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Connection", "keep-alive")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var ilceler []Ilce
		err = json.Unmarshal(body, &ilceler)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		var sehireAitIlceler []Ilce
		for _, ilce := range ilceler {
			if ilce.IlKodu == ilNo {
				sehireAitIlceler = append(sehireAitIlceler, ilce)
			}
		}

		if len(sehireAitIlceler) > 0 {
			ResponseModel := IlceResponseModel{
				IlceArray: sehireAitIlceler,
			}
			ResponseWriter(w, http.StatusOK, "SUCCESS", "Işleminiz başarılı, 'Ilce' listesi getirildi!", ResponseModel)
		} else {
			ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'IlNo' verisine ait ilçe bilgisi bulunamadı!", map[string]interface{}{})
		}
	} else {
		ResponseWriter(w, http.StatusBadRequest, "ERROR", "Işleminiz reddedildi, 'IlNo' verisi göndermelisiniz!", map[string]interface{}{})
	}
}
