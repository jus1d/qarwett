package ssau

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const (
	HeadURL   = "https://ssau.ru/rasp"
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
)

func GetScheduleDocument(groupID string, week int) (*goquery.Document, error) {
	client := http.Client{}

	csrf, err := GetCsrfToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?groupId=%s&selectedWeek=%d", HeadURL, groupID, week)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	AddHeaders(req, csrf)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func AddHeaders(req *http.Request, token string) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CSRF-TOKEN", token)
}

func GetCsrfToken() (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", HeadURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", UserAgent)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	token, exists := doc.Find("meta[name='csrf-token']").Attr("content")
	if !exists {
		return "", err
	}

	return token, nil
}
