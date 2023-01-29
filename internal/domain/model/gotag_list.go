package model

type GotagList []Gotag

func (gl GotagList) Take(target GotagList) GotagList {
	result := gl

	for _, i := range target {
		list := make([]Gotag, 0)

		for _, j := range result {
			if i != j {
				list = append(list, j)
			}
		}

		result = list
	}

	return result
}
