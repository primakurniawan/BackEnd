package main

func findMaxElArray(array []int) int {
	max := array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	return max
}

func sumArr(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func MaximumBuyProduct(money int, productPrice []int) int {

	// your code here
	productBoughtArr := make([]int, len(productPrice))
	moneySpend := 0

	for _, v := range productPrice {
		productBoughtArr = append(productBoughtArr, v)
		if sumArr(productBoughtArr) > money {
			if findMaxElArray(productBoughtArr) > v {

			}
		}
	}

}

func main() {

	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}) // 3

	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4

	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}) // 4

	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}) // 1

	MaximumBuyProduct(0, []int{10000, 30000}) // 0

}
