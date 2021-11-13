package acmp

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
)

func Difficulty(url string) float64 {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return -1
	}
	req.AddCookie(&http.Cookie{Name: "English", Value: "1"})

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return -1
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return -1
	}

	resultString := doc.Find("body.nomargin").Find("center").Find("i").Text()
	resultString = regexp.MustCompile("(\\d+%)").FindString(resultString)

	result, err := strconv.ParseFloat(resultString[:len(resultString)-1], 64)
	if err != nil {
		return -1
	}

	return result
}
