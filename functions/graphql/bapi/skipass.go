package bapi

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"strings"
	"time"
	"unicode"
)

type Card struct {

}

type Skipass struct {
	Plan         string
	CardNumber   string
	TicketNumber string
	PurchaseDate time.Time
}

func (c *Client) CardNumber(ticketNumber string) (string, error) {
	url := fmt.Sprintf("%s?NumTicket=%s", c.base, ticketNumber)

	res, err := c.http.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	return parseCardNumber(doc)
}

func (c *Client) Skipass(ticketNumber string) (Skipass, error) {
	cardNumber, err := c.CardNumber(ticketNumber)
	if err != nil {
		return Skipass{}, err
	}

	url := fmt.Sprintf("%s?NumTicket=%s&Card=%s", c.base, ticketNumber, cardNumber)

	res, err := c.http.Get(url)
	if err != nil {
		return Skipass{}, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Skipass{}, err
	}

	s := Skipass{
		TicketNumber: ticketNumber,
	}

	err = parseSkipass(doc, &s)
	if err != nil {
		return Skipass{}, err
	}

	return s, nil
}

func parseCardNumber(doc *goquery.Document) (string, error) {
	s := doc.Find("#result tr:first-child strong").Text()

	return stripCardNumber(s), nil
}

func stripCardNumber(s string) string {
	result := strings.TrimSpace(s)

	if len(result) == 0 {
		return result
	}

	// card number format
	// XX-XXXX XXXX XXXX XXXX XXXX-X
	startIndex := strings.IndexFunc(result, isHyphen) + 1
	endIndex := strings.LastIndexFunc(result, isHyphen)

	result = result[startIndex:endIndex]


	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		if unicode.Is(unicode.Hyphen, r) {
			return -1
		}

		return r
	}, result)
}

func isHyphen(r rune) bool {
	return unicode.Is(unicode.Hyphen, r)
}

func parseSkipass(doc *goquery.Document, s *Skipass) error {
	if doc == nil {
		return errors.New("parseSkipass no goquery.Document provided (doc == nil)")
	}

	res := doc.Find("#result")

	if res.Length() == 0 {
		return errors.New("parseSkipass no #result table where found")
	}

	if res.Length() != 2 {
		return errors.New(fmt.Sprintf("parseSkipass expecting 2 #result table, got %d", res.Length()))
	}

	fillSkipassHead(res.Eq(0), s)

	return nil
}

func fillSkipassHead (s *goquery.Selection, skipass *Skipass) {
	s.Find("tr").Each(func (i int, s *goquery.Selection) {
		columns := s.Find("td")
		key, _ := decodeWindows1251(columns.Eq(0).Text())
		value, _ := decodeWindows1251(columns.Eq(1).Text())

		switch strings.TrimSpace(key) {
		case "№ картки":
			skipass.CardNumber = strings.TrimSpace(value)
		case "Квиток":
			skipass.Plan = stripCardPlan(value)
		case "Дата продажу":
			skipass.PurchaseDate, _ = parseCardTime(value)
		}
	})
}

func parseCardTime (s string) (time.Time, error) {
	var date time.Time
	loc, err := time.LoadLocation("Europe/Kiev")

	if err != nil {
		return date, err
	}

	date, err = time.ParseInLocation("02.01.2006 15:04", s, loc)

	if err != nil {
		return date, err
	}

	return date.UTC(), nil
}

func stripCardPlan (s string) string {
	var result []rune
	hadSpace := false

	for _, val := range s {
		// skip duplicated space
		if hadSpace && unicode.IsSpace(val) {
			continue
		}
		// no spaces before closing parenthesis
		if hadSpace && val == ')' {
			result = result[:len(result) - 1]
		}

		result = append(result, val)
		hadSpace = unicode.IsSpace(val)
	}

	return string(result)
}

func decodeWindows1251(s string) (string, error) {
	ba := []byte(s)

	dec := charmap.Windows1251.NewDecoder()
	out, err := dec.Bytes(ba)

	if err != nil {
		return "", err
	}

	return string(out), nil
}

