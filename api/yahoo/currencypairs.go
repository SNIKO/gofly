package yahoo

const (
	_ = iota
	stdCurrencyFrom	// UAH
	stdCurrencyTo	// AUD
)

type CurrencyPair struct {
	From string
	To   string
}

type CurrencyPairs []CurrencyPair

func (pair *CurrencyPair) Format(layout string) string {
	var res []byte

	for layout != "" {
		prefix, std, suffix := nextChunk(layout)
		if prefix != "" {
			res = append(res, prefix...)
		}
		if std == 0 {
			break;
		}
		layout = suffix

		switch std {
		case stdCurrencyFrom:
			res = append(res, pair.From...)
		case stdCurrencyTo:
			res = append(res, pair.To...)
		}
	}

	return string(res)
}

func (p *CurrencyPairs) SelectStrings(selector func (pair CurrencyPair) string) []string {
	res := []string{}
	for _, pair := range *p {
		res = append(res, selector(pair))
	}
	return res
}

func (pairs *CurrencyPairs) Contains(from string, to string) bool {
	for _, pair := range *pairs {
		if (pair.From == from && pair.To == to) {
			return true
		}
	}

	return false
}

func nextChunk(layout string) (prefix string, std int, suffix string) {
	for i := 0; i < len(layout); i++ {
		switch c := layout[i]; c {
		case 'U':
			if len(layout) >= i+3 && layout[i:i+3] == "UAH" {
				return layout[0:i], stdCurrencyFrom, layout[i+3:]
			}
		case 'A':
			if len(layout) >= i+3 && layout[i:i+3] == "AUD" {
				return layout[0:i], stdCurrencyTo, layout[i+3:]
			}
		}
	}

	return layout, 0, ""
}
