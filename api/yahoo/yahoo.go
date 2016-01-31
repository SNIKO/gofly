package yahoo

import (
	"net/http"
	"strings"
	"fmt"
	"net/url"
	"encoding/json"
)

type YahooResponse struct {
	Query Query
}

type Query struct {
	Count   int
	Lang    string
	Results QueryResults
}

type QueryResults struct {
	Rate Rates
}

func GetExchangeRates(pairs CurrencyPairs) (Rates, error) {
	pairsAsStrings := pairs.SelectStrings(func (p CurrencyPair) string { return p.Format("\"UAHAUD\"") })
	query := url.QueryEscape(fmt.Sprintf("select * from yahoo.finance.xchange where pair in (%s)", strings.Join(pairsAsStrings, ", ")))

	u := fmt.Sprintf("http://query.yahooapis.com/v1/public/yql?q=%s&env=store://datatables.org/alltableswithkeys&format=json", query)
	response, err := http.Get(u)
	if (err != nil) {
		return nil, err
	}

	defer response.Body.Close()

	var res YahooResponse

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&res)
	if (err != nil) {
		return nil, err
	}

	return Rates(res.Query.Results.Rate), nil
}