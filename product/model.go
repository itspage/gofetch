package product

import "encoding/json"

type Model struct{}

func (p *ProductList) JSON() string {
	b, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return "{}"
	}
	return string(b)
}
