package api

import "time"

func formatDateES(t time.Time) string {
	var months = map[string]string{
		"January":   "enero",
		"February":  "febrero",
		"March":     "marzo",
		"April":     "abril",
		"May":       "mayo",
		"June":      "junio",
		"July":      "julio",
		"August":    "agosto",
		"September": "septiembre",
		"October":   "octubre",
		"November":  "noviembre",
		"December":  "diciembre",
	}

	day := t.Format("02")
	month := months[t.Format("January")]
	year := t.Format("2006")
	return day + " de " + month + " del " + year
}
