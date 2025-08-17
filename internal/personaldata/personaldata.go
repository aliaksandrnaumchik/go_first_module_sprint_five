package personaldata

import "fmt"

type Personal struct {
	Weight float64
	Height float64
	Name   string
}

func (p Personal) Print() {
	info := fmt.Sprintf(
		"Имя: %s\n"+
			"Вес: %.2f кг.\n"+
			"Рост: %.2f м.",
		p.Name,
		p.Weight,
		p.Height,
	)

	fmt.Println(info)
}
