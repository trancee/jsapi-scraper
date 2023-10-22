package lint

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/recoilme/pudge"
	"github.com/sugawarayuuta/sonnet"
)

const EXCHANGE_RATES_BASE_URL = "http://api.exchangeratesapi.io/v1/"
const EXCHANGE_RATES_ENDPOINT = "latest"

type Response struct {
	Success bool `json:"success"` // Returns true or false depending on whether or not your API request has succeeded.

	Timestamp uint64 `json:"timestamp"` // Returns the exact date and time (UNIX time stamp) the given rates were collected.

	Base string `json:"base"` // Returns the three-letter currency code of the base currency used for this request.
	Date string `json:"date"`

	Rates struct { // Returns exchange rate data for the currencies you have requested.
		CHF float64 `json:"CHF"`
		// EUR float64 `json:"EUR"`
	} `json:"rates"`

	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func EUR_CHF_v2() float64 {
	type _Body struct {
		Currencies struct {
			CHF string `json:"CHF"`
		} `json:"currencies"`
	}

	resp, err := http.Get("https://www.tradeinn.com/?action=get_info_pais&id_tienda=16&id_pais=192")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var body _Body
	if err := sonnet.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
		panic(err)
	}

	EUR_CHF, err := strconv.ParseFloat(body.Currencies.CHF, 32)
	if err != nil {
		panic(err)
	}

	return EUR_CHF
}

func EUR_CHF() float64 {
	db, err := pudge.Open("rates", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	url := fmt.Sprintf("%s%s", EXCHANGE_RATES_BASE_URL, EXCHANGE_RATES_ENDPOINT)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	var etag string
	if err := db.Get("ETag", &etag); err == nil {
		req.Header.Set("If-None-Match", etag)
	}
	var date string
	if err := db.Get("Date", &date); err == nil {
		req.Header.Set("If-Modified-Since", date)
	}

	q := req.URL.Query()
	q.Add("access_key", os.Getenv("EXCHANGE_RATES_ACCESS_KEY"))
	// q.Add("base", "EUR")
	// q.Add("symbols", "CHF")
	req.URL.RawQuery = q.Encode()
	// fmt.Printf("%+v\n", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// fmt.Printf("%+v\n", resp)

	var EUR_CHF float64

	if resp.StatusCode == http.StatusNotModified {
		if err := db.Get("EUR_CHF", &EUR_CHF); err != nil {
			panic(err)
		}

		return EUR_CHF
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", string(body))

	var response Response
	if err := sonnet.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", response)

	if !response.Success {
		panic(response.Error.Code + "\n" + response.Error.Message)
	}

	EUR_CHF = response.Rates.CHF

	if err := db.Set("ETag", resp.Header.Get("ETag")); err != nil {
		panic(err)
	}
	if err := db.Set("Date", resp.Header.Get("Date")); err != nil {
		panic(err)
	}

	if err := db.Set("EUR_CHF", EUR_CHF); err != nil {
		panic(err)
	}

	return EUR_CHF
}
