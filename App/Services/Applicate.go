package Services

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func init() {
}

func GetRequestToken() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adres.nvi.gov.tr/VatandasIslemleri/AdresSorgu", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Accept-Language", "tr-TR,tr;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Host", "adres.nvi.gov.tr")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "TS01eef051=0179b2ce4515226e1701bbc3956a2d6ad2b2842715856bd69c2793178e4a922597f65a0c5de346f55685574426520976ded9211141005bcaf90002144b1ea56430058e17fc; __RequestVerificationToken=U44le7ehuJY7Lnoje3BVUjLKcD5Y5xBRu2D2l5eB-k_8yka5LTBNyY2kH-ACrA8IousOpeo_GJZdGKTts_uJjnESrd41; _ga=GA1.3.1408641006.1704641143; _gid=GA1.3.1538294979.1704641143; browser-check=%22done%22")
	req.Header.Set("Sec-Fetch-Dest", "document")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	requestVerificationToken, exists := doc.Find("input[name=__RequestVerificationToken]").Attr("value")
	if !exists {
		log.Fatal("RequestVerificationToken bulunamadı")
	}

	return requestVerificationToken
}

func ResponseWriter(w http.ResponseWriter, statusCode int, statusText string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseData := map[string]interface{}{
		"MessageCode":       statusText,
		"MessageList":       []string{message},
		"RestGenericObject": data,
	}
	jsonData, err := json.MarshalIndent(responseData, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, _ = w.Write(jsonData)
}
