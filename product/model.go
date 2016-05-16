package product

import "encoding/json"

type Model interface{}

func JSON(m Model) string {
	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return "{}"
	}
	return string(b)
}
