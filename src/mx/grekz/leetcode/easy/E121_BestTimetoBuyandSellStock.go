package easy

func maxProfit(prices []int) int {
    buy := 2147483647
    result := 0
    for _, stockPrice := range prices {
        if(stockPrice < buy){
            buy = stockPrice
        }else if ( result < stockPrice - buy ) {
            result = stockPrice - buy
        }
    }
    return result
}