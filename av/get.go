package av

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Get is responsible for the actual GET request to retrieve data
// from Alpha Vantage and returns a pointer to the response body
func Get(s string) *[]byte {
	c := LoadConfig()
	p := Param{
		Function: "TIME_SERIES_DAILY",
		Symbol:   s,
		APIKey:   c.APIKey,
	}

	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		log.Println(err)
	}

	// Append the query parameters to the URL
	q := req.URL.Query()

	q.Add("function", p.Function)
	q.Add("symbol", string(p.Symbol))
	q.Add("apikey", p.APIKey)

	req.URL.RawQuery = q.Encode()

	// Send the GET request and read the body of the response
	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return &body
}

// GetStock retrieves the data for a given stock Symbol, returning an ordered
// slice of the keys, which are dates for each stock price, and a map of the
// date/price key/value pairs. This is done to provide a way to iterate through
// each price in chronological order.
func GetStock(s string) {
	var body TimeSeriesDaily

	raw := Get(s)
	if err := json.Unmarshal(*raw, &body); err != nil {
		panic(err)
	}

	// map of dates:map[string]price
	series := body.Series
	// keys of series map
	dates := series.MapSort()

	startDate := dates[0]
	prices := make([]float32, len(dates))
	for i, date := range dates {
		// have to assert in order to convert bb
		interfaceMap := series[date].(map[string]interface{})
		floatMap := InterfaceMap(interfaceMap).ConvertToFloats()
		prices[i] = floatMap["4. close"]
	}

	fmt.Println(startDate, prices)
}
