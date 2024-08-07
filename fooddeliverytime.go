package piscine

type food struct {
	preptime int
}

var foodPrepTimes = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

func FoodDeliveryTime(order string) int {
	if item, exists := foodPrepTimes[order]; exists {
		return item.preptime
	}
	return 404
}
