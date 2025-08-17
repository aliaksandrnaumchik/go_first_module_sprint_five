package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			fmt.Printf("Ошибка парсинга данных: %v", err)
			continue
		}

		if info, err := dp.ActionInfo(); err != nil {
			fmt.Printf("Ошибка получения информации")
		} else {
			fmt.Println(info)
		}
	}
}
