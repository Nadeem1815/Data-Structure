package Array

func MaxProfit(price []int) int {
	var profit = 0
	var minprice = price[0]
	for i := 1; i < len(price); i++ {
		if price[i] < minprice {
			minprice = price[i]

		} else if (price[i] - minprice) > profit {
			profit = price[i]
		}

	}
	return profit
}
