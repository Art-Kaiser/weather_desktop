package formatRu

import (
	"strings"
)

func Format(str string, decrease bool) string {
	data := initRu()
	if value, exists := data.list[strings.ToLower(str)]; exists {
		if decrease {
			valueDecrease := data.decrease[strings.ToLower(str)]
			return valueDecrease
		}
		return value
	}
	return str
}
