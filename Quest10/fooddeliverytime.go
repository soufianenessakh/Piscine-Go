package Quest10

type Food struct {
	PrepTime int
}

var menu = map[string]Food{
	"burger":  {15},
	"chips":   {10},
	"nuggets": {12},
}

func FoodDeliveryTime(order string) int {
	total := 0
	word := ""

	for _, ch := range order {
		if ch == ' ' { 
			if word != "" {
				food, ok := menu[word]
				if !ok {
					return 404
				}
				total += food.PrepTime
				word = ""
			}
		} else {
			word += string(ch) 
		}
	}
	if word != "" {
		food, ok := menu[word]
		if !ok {
			return 404
		}
		total += food.PrepTime
	}

	return total
}