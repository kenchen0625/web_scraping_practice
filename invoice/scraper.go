package invoice

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scraper interface {
	Fetch()
}

type Invoice struct {
	URL      string
	response *http.Response
}

func (i *Invoice) Fetch() error {
	response, err := http.Get(i.URL)
	if err != nil {
		return err
	}

	i.response = response

	return nil
}

func (i *Invoice) GetJackpot() ([]string, error) {
	var prizes []string

	defer i.response.Body.Close()

	document, err := goquery.NewDocumentFromReader(i.response.Body)
	if err != nil {
		return nil, err
	}

	// Find and print image URLs
	document.Find("p.etw-tbiggest.mb-md-4").Each(func(index int, element *goquery.Selection) {
		if index > 2 {
			return
		}
		prize := strings.TrimSpace(element.Text())
		// fmt.Println(index, prize)
		prizes = append(prizes, prize)
	})

	return prizes, nil
}
