package yahoo

import (
	"strings"
	"encoding/json"
)

type Rates []Rate

type Rate struct {
	Id 		string
	Name 	string
	Rate 	float64 `json:",string"`
}

func (r *Rates) UnmarshalJSON(b []byte) (err error) {
	rate, rates := Rate{}, []Rate{}

	if err = json.Unmarshal(b, &rate); err == nil {
		*r = Rates([]Rate { rate })
		return
	}

	if err = json.Unmarshal(b, &rates); err == nil {
		*r = Rates(rates)
		return
	}

	return
}

func (r *Rate) From(currency string) bool {
	currencies := strings.Split(r.Name, "/")
	return len(currencies) > 0 && currencies[0] == currency
}

func (r *Rate) To(currency string) bool {
	currencies := strings.Split(r.Name, "/")
	return len(currencies) > 1 && currencies[1] == currency
}

func (rates *Rates) Find(from string, to string) *Rate {
	if rates == nil {
		return nil;
	}

	return rates.Filter(func (r Rate) bool { return r.From(from) && r.To(to) } ).HeadOrNil()
}

func (rates *Rates) Filter(predicate func (Rate) bool) Rates {
	result := Rates {}

	if (rates != nil && predicate != nil) {
		for _, r := range *rates {
			if (predicate(r)) {
				result = append(result, r)
			}
		}
	}

	return result
}

func (rates Rates) HeadOrNil() *Rate {
	if len(rates) > 0 {
		return &(rates)[0]
	} else {
		return nil
	}
}