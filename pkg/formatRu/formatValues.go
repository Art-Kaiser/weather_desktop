package formatRu

type ru struct {
	list     map[string]string
	decrease map[string]string
}

func initRu() ru {
	data := ru{list: map[string]string{
		"january":   "Январь",
		"february":  "Февраль",
		"march":     "Март",
		"april":     "Апрель",
		"may":       "Май",
		"june":      "Июнь",
		"july":      "Июль",
		"august":    "Август",
		"september": "Сентябрь",
		"october":   "Октябрь",
		"november":  "Ноябрь",
		"december":  "Декабрь",
		"monday":    "Понедельник",
		"tuesday":   "Вторник",
		"wednesday": "Среда",
		"thursday":  "Четверг",
		"friday":    "Пятница",
		"saturday":  "Суббота",
		"sunday":    "Воскресенье",
	},
		decrease: map[string]string{
			"january":   "Янв",
			"february":  "Фев",
			"march":     "Мар",
			"april":     "Апр",
			"may":       "Май",
			"june":      "Июн",
			"july":      "Июл",
			"august":    "Авг",
			"september": "Сен",
			"october":   "Окт",
			"november":  "Ноя",
			"december":  "Дек",
			"monday":    "Пн",
			"tuesday":   "Вт",
			"wednesday": "Ср",
			"thursday":  "Чт",
			"friday":    "Пт",
			"saturday":  "Сб",
			"sunday":    "Вс",
		}}
	return data
}
