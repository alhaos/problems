package leetCode

import (
	"bytes"
	"math"
)

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minPrice := math.MaxInt64
	maxProfitValue := 0

	for _, price := range prices {

		if price < minPrice {
			minPrice = price
		}

		profit := price - minPrice

		if profit > maxProfitValue {
			maxProfitValue = profit
		}

	}
	return maxProfitValue
}

func mergeAlternately(word1 string, word2 string) string {
	var i1, i2 int
	var buffer bytes.Buffer

	for i1 < len(word1) || i2 < len(word2) {
		if i1 < len(word1) {
			buffer.WriteByte(word1[i1])
			i1++
		}
		if i2 < len(word2) {
			buffer.WriteByte(word2[i2])
			i2++
		}
	}

	return buffer.String()
}

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLength := gcd(len(str1), len(str2))

	return str1[:gcdLength]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxCandies := 0

	for _, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}
	}

	result := make([]bool, len(candies))

	for index, value := range candies {
		result[index] = value+extraCandies >= maxCandies
	}

	return result
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	index := -1
	lastIndex := len(flowerbed) - 1
	placeCounter := 0

	for {

		index++

		if index > lastIndex {
			break
		}

		if index != 0 && flowerbed[index-1] == 1 {
			continue
		}

		if index != lastIndex && flowerbed[index+1] == 1 {
			continue
		}

		if flowerbed[index] == 1 {
			index++
			continue
		}

		flowerbed[index] = 1
		placeCounter++

	}

	return placeCounter >= n
}

func lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	charCount := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 0x20 {
			if charCount > 0 {
				return charCount
			}
		} else {
			charCount++
		}
	}

	return charCount
}
