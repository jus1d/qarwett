package ssau

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

const (
	HeadURL   = "https://ssau.ru"
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
)

type SearchGroupResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"text"`
	ScheduleURL string `json:"url"`
}

func GetScheduleDocument(groupID string, week int) (*goquery.Document, error) {
	client := http.Client{}

	cookies, csrf, err := GetCookiesAndToken()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/rasp?groupId=%s&selectedWeek=%d", HeadURL, groupID, week)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	AddHeaders(req, csrf)
	AddCookies(req, cookies)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func GetGroupBySearchQuery(query string) ([]SearchGroupResponse, error) {
	cookies, token, err := GetCookiesAndToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get csrf token: %s", err.Error())
	}

	client := http.Client{}

	req, err := http.NewRequest("POST", HeadURL+"/rasp/search?text="+query, nil)
	if err != nil {
		return nil, err
	}

	AddHeaders(req, token)
	AddCookies(req, cookies)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var list []SearchGroupResponse
	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, &list); err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("responce %s: %s", res.Status, req.URL)
	}

	return list, nil
}

func GetCookiesAndToken() ([]*http.Cookie, string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", HeadURL+"/rasp", nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Add("User-Agent", UserAgent)

	res, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}

	cookies := res.Cookies()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return cookies, "", err
	}
	token, exists := doc.Find("meta[name='csrf-token']").Attr("content")
	if !exists {
		return cookies, "", err
	}

	return cookies, token, nil
}

func AddHeaders(req *http.Request, token string) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CSRF-TOKEN", token)
}

func AddCookies(req *http.Request, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
}
