package sliding_window

func maxProfit(prices []int) int {
	lowestPrice := prices[0]
	bestDeal := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < lowestPrice {
			lowestPrice = prices[i]
		}
		deal := prices[i] - lowestPrice
		if deal > bestDeal {
			bestDeal = deal
		}
	}
	return bestDeal
}
