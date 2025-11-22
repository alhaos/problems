// https://leetcode.com/studyplan/top-interview-150/

package topInterview

import (
	"bytes"
	"math"
	"math/rand"
	"sort"
	"strings"
)

// merge problem
//
// 88. Merge Sorted Array
//
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
// representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To
// accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
//
// Example 1:
//
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The expected of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
// Example 2:
//
// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The expected of the merge is [1].
// Example 3:
//
// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The expected of the merge is [1].
// Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge expected can fit in nums1.
//
// Вам даны два целочисленных массива nums1 и nums2, отсортированных в неубывающем порядке, и два целых числа m и n,
// представляющих количество элементов в nums1 и nums2 соответственно.
//
// Объедините nums1 и nums2 в один массив, отсортированный в неубывающем порядке.
// Окончательный отсортированный массив не должен возвращаться функцией, а вместо этого должен храниться внутри массива
// nums1. Чтобы учесть это, nums1 имеет длину m + n, где первые m элементов обозначают элементы, которые должны быть
// объединены, а последние n элементов устанавливаются в 0 и должны игнорироваться. nums2 имеет длину n.
func merge(nums1 []int, m int, nums2 []int, n int) {

	// Инициализировать переменные в которых будут индексы слайсов
	index1 := m - 1          // Индекс просматриваемого элемента в слайсе nums1
	index2 := n - 1          // Индекс просматриваемого элемента в слайсе nums2
	commonIndex := m + n - 1 // Индекс заполняемого элемента в слайсе nums1

	// Остановить цикл когда перебраны все элементы в слайсе nums2
	for index2 >= 0 {
		// Если еще не перебраны все элементы слайса num1
		// и текущий элемент слайса num2 больше элемента в стайсе nums1
		if index1 >= 0 && nums1[index1] > nums2[index2] {
			// Заменить крайний элемент в слайсе nums1 на текущий элемент
			// слайса nums1
			nums1[commonIndex] = nums1[index1]
			// Уменьшить индекс текущего элемента в nums1
			index1--
		} else {
			// В случае когда перебраны все элементы слайса nums1 или
			// текущий элемент слайса nums2 больше текущего элемента слайса nums1
			nums1[commonIndex] = nums2[index2]
			// Уменьшить индекс текущего элемента в nums2
			index2--
		}
		// Уменьшить индекс заполняемого элемента в nums1
		commonIndex--
	}
}

// removeElement
//
// 27. Remove element
// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
//
// Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Custom Judge:
//
// The judge will test your solution with the following code:
//
// int[] nums = [...]; // Input array
// int val = ...; // Value to remove
// int[] expectedNums = [...]; // The expected answer with correct length.
//
//	// It is sorted with no values equaling val.
//
// int k = removeElement(nums, val); // Calls your implementation
//
// assert k == expectedNums.length;
// sort(nums, 0, k); // Sort the first k elements of nums
//
//	for (int i = 0; i < actualLength; i++) {
//	    assert nums[i] == expectedNums[i];
//	}
//
// If all assertions pass, then your solution will be accepted.
//
// Example 1:
//
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are underscores).
// Example 2:
//
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order.
// It does not matter what you leave beyond the returned k (hence they are underscores).
//
// Constraints:
//
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 50
// 0 <= val <= 100
// Тут задача избавится от всех элементов в слайсе nums которые равны переменной val
// изменить данные в слайсе и вернуть количество элементов не равных val
func removeElement(nums []int, val int) int {

	// Объявить переменную в которой будет индекс крайнего вставленного элемента
	insertIndex := 0

	// Инициировать цикл for range по слайсу
	for _, num := range nums {
		// если текущий элемент слайса nums не равен проверяемому значению
		if num != val {
			// заменить элемент в слайсе nums c индексом крайнего вставленного элемента.
			nums[insertIndex] = num
			// Увеличить индекс вставленного элемента на один
			insertIndex++
		}
	}

	// Вернуть индекс
	return insertIndex
}

// removeDuplicates
//
// 26. Remove Duplicates from Sorted Array
// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
//
// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
//
// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Custom Judge:
//
// The judge will test your solution with the following code:
//
// int[] nums = [...]; // Input array
// int[] expectedNums = [...]; // The expected answer with correct length
//
// int k = removeDuplicates(nums); // Calls your implementation
//
// assert k == expectedNums.length;
// for (int i = 0; i < k; i++) {
// assert nums[i] == expectedNums[i];
// }
// If all assertions pass, then your solution will be accepted.
//
// Example 1:
//
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).
// Example 2:
//
// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).
//
// Constraints:
//
// 1 <= nums.length <= 3 * 104
// -100 <= nums[i] <= 100
// nums is sorted in non-decreasing order.
func removeDuplicates(nums []int) int {

	// Инициировать переменную содержащую длину слайса nums
	lenNums := len(nums)

	// Если длина слайса nums меньше двух вернуть длину слайса
	if lenNums < 2 {
		return lenNums
	}

	// Инициализировать переменную содержащую индекс изменяемого элемента
	insertIndex := 0

	// Инициализировать переменную содержащую индекс проверяемого элемента
	index := 1

	// Прервать цикл когда индекс проверяемого элемента сравнялся с длиной слайса nums
	for index < lenNums {
		// Если элемент с индексом изменяемого элемента не равен элементу с индексом просматриваемого элемента
		if nums[insertIndex] != nums[index] {
			// Увеличить индекс изменяемого элемента на один
			insertIndex++
			// Прировнять значение элемента с индексом изменяемого элемента к значению проверяемого элемента
			nums[insertIndex] = nums[index]
		}
		// Увеличить индекс проверяемого элемента на единицу
		index++
	}

	// Вернуть индекс изменяемого элемента + 1
	return insertIndex + 1
}

// removeDuplicates
//
// 80. Remove Duplicates from Sorted Array II
// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.
//
// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
//
// Return k after placing the final result in the first k slots of nums.
//
// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
//
// Custom Judge:
//
// The judge will test your solution with the following code:
//
// int[] nums = [...]; // Input array
// int[] expectedNums = [...]; // The expected answer with correct length
//
// int k = removeDuplicates(nums); // Calls your implementation
//
// assert k == expectedNums.length;
//
//	for (int i = 0; i < k; i++) {
//	    assert nums[i] == expectedNums[i];
//	}
//
// If all assertions pass, then your solution will be accepted.
//
// Example 1:
//
// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]
// Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).
// Example 2:
//
// Input: nums = [0,0,1,1,1,1,2,3,3]
// Output: 7, nums = [0,0,1,1,2,3,3,_,_]
// Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).
//
// Constraints:
//
// 1 <= nums.length <= 3 * 104
// -104 <= nums[i] <= 104
// nums is sorted in non-decreasing order.//
func removeDuplicatesII(nums []int) int {

	// Инициировать переменные
	indexMutable := 0       // Индекс изменяемого элемента
	indexChecked := 1       // Индекс проверяемого элемента
	idDubFound := false     // Флаг, который указывает на то был ли уже найден дубликат
	lengthNums := len(nums) // Длина слайса nums

	// Если в слайсе nums меньше двух элементов
	if lengthNums < 2 {
		// Возвращаем длину слайса nums
		return lengthNums
	}

	// Прервать цикл если индекс проверяемого элемента достиг длинны слайса
	for indexChecked < lengthNums {

		// Если изменяемый элемент и проверяемый элемент равны и дубликат ранее не найден
		if nums[indexMutable] == nums[indexChecked] && idDubFound {
			// Увеличить индекс проверяемого элемента на один
			indexChecked++
			// Прервать текущую итерацию цикла
			continue
		}

		// Если изменяемый элемент и проверяемый элемент равны и дубликат ранее найден
		if nums[indexMutable] == nums[indexChecked] && !idDubFound {
			// Увеличить индекс изменяемого элемента на один
			indexMutable++
			// Изменить значение изменяемого элемента на значение проверяемого элемента
			nums[indexMutable] = nums[indexChecked]
			// Увеличить индекс проверяемого элемента на один
			indexChecked++
			// Установить признак того что дубликат найден
			idDubFound = true
			// Прервать текущую итерацию цикла
			continue
		}

		// Если проверяемый элемент больше изменяемого
		if nums[indexChecked] > nums[indexMutable] {
			// Увеличить индекс изменяемого элемента на один
			indexMutable++
			// Изменить значение изменяемого элемента на значение проверяемого элемента
			nums[indexMutable] = nums[indexChecked]
			// Увеличить индекс проверяемого элемента на один
			indexChecked++
			// Сбросить признак того что дубликат найден
			idDubFound = false
		}
	}
	// Вернуть индекс изменяемого элемента + 1
	return indexMutable + 1
}

// majorityElement
//
// 169. Majority Element
// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
//
// Example 1:
//
// Input: nums = [3,2,3]
// Output: 3
// Example 2:
//
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
//
// Constraints:
//
// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109//
func majorityElement(nums []int) int {

	// Инициировать переменную длинной слайса
	lengthNums := len(nums)

	// Инициировать карту, ключ - значение переменной в nums, значение - количество элемента в слайсе nums
	elementsMap := make(map[int]int)

	// Перебрать элементы в слайсе nums
	for _, num := range nums {
		// К значению ключа карты добавить единицу
		elementsMap[num]++
	}

	// Перебрать карту
	for k, v := range elementsMap {
		// Если значение элемента больше половины длины слайса nums
		if v > lengthNums/2 {
			// вернуть ключ элемента
			return k
		}
	}

	// По условиям задачи мажорный элемент обязан быть
	// и так как это код не должен быть достигнут тут паника
	panic("no majority element found")
}

// rotate
//
// 189. Rotate Array
//
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
// Example 1:
//
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:
//
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]
//
// Constraints:
//
// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105
func rotate(nums []int, k int) {
	// Инициализировать переменную длинной слайса nums
	lengthNums := len(nums)
	// Инициализировать переменную остатком от деления количества шагов на длину слайса
	remainder := k % lengthNums
	// Скопировать в nums результат объединения части слайса nums от разности
	// длины слайса и остатка с каждым элементом части слайса от начала до разности длины с остатком
	copy(nums, append(nums[lengthNums-remainder:], nums[:lengthNums-remainder]...))
}

// maxProfit
//
// 121. Best Time to Buy and Sell Stock
// Solved
// Easy
// Topics
// Companies
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
// Example 1:
//
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.
//
// Constraints:
//
// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104
func maxProfit(prices []int) int {
	// Инициировать переменную содержащую минимальную цену максимальным возможным int64
	minPrice := math.MaxInt64
	// Инициировать переменную содержащую максимально возможным прибыли нулем
	maxProfitValue := 0
	// Перебрать цены
	for _, price := range prices {
		// Если текущая цена меньше минимальной цены
		if price < minPrice {
			// Минимальная цена равна текущей цене
			minPrice = price
		}
		// Прибыль равна разности текущей цены и минимальной
		profit := price - minPrice
		// Если прибыль больше максимально возможной прибыли
		if profit > maxProfitValue {
			// Обновить значение максимально возможной прибыли текущей
			maxProfitValue = profit
		}
	}
	// Вернуть максимально возможную прибыль
	return maxProfitValue
}

// maxProfitII
//
// 122. Best Time to Buy and Sell Stock II
//
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
//
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
//
// Find and return the maximum profit you can achieve.
//
// Example 1:
//
// Input: prices = [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Total profit is 4 + 3 = 7.
// Example 2:
//
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Total profit is 4.
// Example 3:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
//
// Constraints:
//
// 1 <= prices.length <= 3 * 104
// 0 <= prices[i] <= 104
func maxProfitII(prices []int) int {

	// Инициировать накопленную прибыль
	profit := 0

	// Инициировать индекс текущей цены
	index := 1

	// Инициировать цикл
	for {

		// Если индекс достиг длины слайса цен
		if index >= len(prices) {

			// Прервать цикл
			break
		}

		// Если текущая цена больше предыдущей
		if prices[index] > prices[index-1] {

			// Прибавить к накопленной прибыли разность между текущей ценой и предыдущей
			profit += prices[index] - prices[index-1]
		}

		// Увеличить индекс текущей цены на один
		index++
	}
	// Вернуть накопленную прибыль
	return profit
}

// canJump
//
// 55. Jump Game
// You are given an integer array nums. You are initially positioned at the array'strs first index, and each element in the array represents your maximum jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
//
// Example 1:
//
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:
//
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
//
// Constraints:
//
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105
func canJump(nums []int) bool {

	// Инициировать индекс позиции длиной стлайса минус единицей
	indexPosition := len(nums) - 1

	// Инициировать индекс доступной позиции длиной стлайса минус единицей
	indexReachPosition := indexPosition

	for {
		// Если индекс позиции меньше нуля
		if indexPosition < 0 {
			// Прервать цикл
			break
		}

		// Если индекс позиции плюс значение элемента с индексом позиции
		// больше или равно индексу доступной позиции
		if indexPosition+nums[indexPosition] >= indexReachPosition {
			// Индекс доступной позиции равен индексу позиции
			indexReachPosition = indexPosition
		}

		// Уменьшить индекс позиции на единицу
		indexPosition--
	}

	// Если индекс доступной позиции равен нулю вернуть истину
	if indexReachPosition == 0 {
		return true
	}

	// Вернуть ложь
	return false
}

// jump
//
// 45. Jump Game II
//
// You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
//
// Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums// [i + j] where:
//
// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
//
// Example 1:
//
// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:
//
// Input: nums = [2,3,0,1,4]
// Output: 2
//
// Constraints:
//
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// It'strs guaranteed that you can reach nums[n - 1].
func jump(nums []int) int {

	// lastNumsIndex содержит крайний индекс слайса
	lastNumsIndex := len(nums) - 1

	// currentAvailableIndex содержит индекс элемента до которого мы можем допрыгнуть в текущей итерации цикла
	currentAvailableIndex := 0

	// farthestAvailableIndex содержит индекс самого дальнего элемента из обнаруженных за все предыдущие итерации цикла
	farthestAvailableIndex := 0

	// jumps счетчик прыжков
	jumps := 0

	// index итерируемого элемента цикла
	index := 0

	for {

		// Обновить элемент до которого можно допрыгнуть вообще, учитывая предыдущие итерации цикла
		farthestAvailableIndex = max(farthestAvailableIndex, index+nums[index])

		// Если элемент до которого можно допрыгнуть равен крайнему элементу - цель достигнута
		if farthestAvailableIndex >= lastNumsIndex {
			// Увеличить счетчик прыжков на один
			jumps++
			// Прервать цикл
			break
		}

		// Если текущий индекс равен индексу элемента до которого мы можем допрыгнуть за эту
		// итерацию - значит мы дошли в тупик, но по условия задачи гарантируется наличие пути к крайнему
		// элементу
		if index == currentAvailableIndex {
			// Увеличить счетчик прыжков на один
			jumps++
			// Индекс элемента до которого можно допрыгнуть приравнять
			// к индексу наиболее дальнего элемента который найден за все предыдущие итерации цикла
			currentAvailableIndex = farthestAvailableIndex
		}

		// Увеличить текущий индекс на единицу
		index++
	}
	// Вернуть счетчик прыжков
	return jumps
}

// hIndex
//
// 274. H-Index
// Medium
// Topics
// Companies
// Hint
// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher'strs h-index.
//
// According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
//
// Example 1:
//
// Input: citations = [3,0,6,1,5]
// Output: 3
// Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
// Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
// Example 2:
//
// Input: citations = [1,3,1]
// Output: 1
//
// Constraints:
//
// n == citations.length
// 1 <= n <= 5000
// 0 <= citations[i] <= 1000//
func hIndex(citations []int) int {

	// Воспользуюсь сортировкой целочисленных слайсов с произвольной функцией сортировки
	// Отсортировать слайс по убыванию
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	// Переменная hIndexValue содержит крайний h-индекс подходящий под критерии
	hIndexValue := 0
	// Цикл range по элементам слайса citations
	// i индекс текущего элемента слайса
	// v значение текущего элемента слайса
	for i, v := range citations {
		// если значение текущего элемента слайса больше индекса элемента
		if v >= i+1 {
			// Обновляем hIndexValue становится индексом текущего элемента плюс единица так как чети индекса ведется с единицы
			hIndexValue = i + 1
			// Прервать итерацию цикла
			continue
		}
		// Прервать цикл
		break
	}

	// Вернуть значение h-индекса
	return hIndexValue
}

// region RandomizedSet 380

// RandomizedSet 380. Insert Delete GetRandom O(1)
// Implement the RandomizedSet class:
//
// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements (it'strs guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
// You must implement the functions of the class such that each function works in average O(1) time complexity.
//
// Example 1:
//
// Input
// ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// Output
// [null, true, false, true, 2, true, false, 2]
//
// Explanation
// RandomizedSet randomizedSet = new RandomizedSet();
// randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
// randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
// randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
// randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
// randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
// randomizedSet.insert(2); // 2 was already in the set, so return false.
// randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
//
// Constraints:
//
// -231 <= val <= 231 - 1
// At most 2 * 105 calls will be made to insert, remove, and getRandom.
// There will be at least one element in the data structure when getRandom is called.

// RandomizedSet структура содержащая данные
type RandomizedSet struct {
	elements        []int       // слайс элементов
	elementIndexMap map[int]int // карта содержащая значение элемента и массив элемента
}

// Constructor конструктор для структуры RandomizedSet
func Constructor() RandomizedSet {
	// Вернуть RandomizedSet
	return RandomizedSet{
		elements:        []int{},           // Поле elements инициировать пустым слайсом целочисленных элементов
		elementIndexMap: make(map[int]int), // Создать пустую карту
	}
}

// Insert вставляет элемент при отсутствии данного элемента
// Возвращает ИСТИНУ в случае успеха, ЛОЖЬ в противоположном
func (this *RandomizedSet) Insert(val int) bool {

	// Инициировать exist бинарным признаком наличия элемента в карте элементов
	_, exist := this.elementIndexMap[val]
	// В случае наличия элемента
	if exist {
		// Вернуть ЛОЖЬ
		return false
	}

	// Добавить элемент в слайс элементов
	this.elements = append(this.elements, val)

	// Добавить элемент в карту индексов элементов
	this.elementIndexMap[val] = len(this.elements) - 1

	// Вернуть ИСТИНУ
	return true
}

// Remove удаляет элемент при наличии данного элемента
// Возвращает ИСТИНУ в случае успеха, ЛОЖЬ в противоположном
func (this *RandomizedSet) Remove(val int) bool {

	// Инициировать exist бинарным признаком наличия элемента в карте элементов
	_, exist := this.elementIndexMap[val]
	// В случае отсутствия элемента
	if !exist {
		// Вернуть ЛОЖЬ
		return false
	}

	// Определить индекс элемента по карте
	index := this.elementIndexMap[val]

	// Заменить значение элемента с индексом index на значение крайнего элемента в слайсе элементов
	this.elements[index] = this.elements[len(this.elements)-1]

	// Заменить значение индекса, нового значения элемента в карте индексов элементов
	this.elementIndexMap[this.elements[index]] = index

	// Переопределить слайс элементов им же без крайнего элемента
	this.elements = this.elements[:len(this.elements)-1]

	// Удалить элемент из карты элементов
	delete(this.elementIndexMap, val)

	// Вернуть ИСТИНУ
	return true
}

// GetRandom возвращает случайный элемент
func (this *RandomizedSet) GetRandom() int {
	// Проверить наличие элементов в слайсе
	if len(this.elements) == 0 {
		// Вернуть ноль
		return 0
	}
	// Получить индекс случайного элемента
	index := rand.Intn(len(this.elements))

	// Вернуть значение случайного элемента
	return this.elements[index]
}

// endregion

// productExceptSelf
//
// 238. Product of Array Except Self
// Medium
// Topics
// Companies
// Hint
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// Example 1:
//
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:
//
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
//
// Constraints:
//
// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.
func productExceptSelf(nums []int) []int {

	// Инициировать переменную содержащую накопленное произведение единицей
	cumulativeProduct := 1
	// Инициировать переменную содержащую индекс итерируемого элемента слайса нулем
	index := 0
	// Инициировать промежуточный слайс нулями, длиной аналогичной nums
	output := make([]int, len(nums))

	// Запустить цикл
	for {
		// Значение элемента промежуточного слайса с индексом Index
		// прировнять к значению переменной накопленной произведения
		output[index] = cumulativeProduct
		// Значение накопленного произведения приравнять к произведению накопленного
		// произведения и значения элемента nums c индексом Index
		cumulativeProduct *= nums[index]
		// Если значение Index выходит за приделы крайнего элемента слайса nums
		if index >= len(nums)-1 {
			// Прервать цикл
			break
		}
		// Инкрементировать Index
		index++
	}

	// Значение накопленного предвидения прировнять к единице
	cumulativeProduct = 1

	for {
		// Значение элемента промежуточного слайса с индексом Index прировнять к произведению
		// значение элемента промежуточного слайса с индексом Index и накопленного произведения
		output[index] *= cumulativeProduct
		// Значение накопленного произведения прировнять к произведению наколенного произведения
		// и элемента слайса nums с индексом Index
		cumulativeProduct *= nums[index]
		// Если Index меньше нуля
		if index <= 0 {
			// Прервать цикл
			break
		}
		// Декриментировать Index
		index--

	}

	// Вернуть промежуточный слайс
	return output
}

// canCompleteCircuit
//
// 134. Gas Station
//
// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
//
// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.
//
// Given two integer arrays gas and cost, return the starting gas station'strs index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a  solution, it is guaranteed to be unique.
//
// Example 1:
//
// Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// Output: 3
// Explanation:
// Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
// Travel to station 4. Your tank = 4 - 1 + 5 = 8
// Travel to station 0. Your tank = 8 - 2 + 1 = 7
// Travel to station 1. Your tank = 7 - 3 + 2 = 6
// Travel to station 2. Your tank = 6 - 4 + 3 = 5
// Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
// Therefore, return 3 as the starting index.
// Example 2:
//
// Input: gas = [2,3,4], cost = [3,4,3]
// Output: -1
// Explanation:
// You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
// Let'strs start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
// Travel to station 0. Your tank = 4 - 3 + 2 = 3
// Travel to station 1. Your tank = 3 - 3 + 3 = 3
// You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
// Therefore, you can't travel around the circuit once no matter where you start.
//
// Constraints:
//
// n == gas.length == cost.length
// 1 <= n <= 105
// 0 <= gas[i], cost[i] <= 104
func canCompleteCircuit(gas []int, cost []int) int {

	sumGas := 0 // Сюда суммирую элементы слайса gas
	for _, n := range gas {
		sumGas += n
	}

	sumCost := 0 // Сюда суммирую элементы слайса cost
	for _, n := range cost {
		sumCost += n
	}

	// Пути нет в том случае если сумма топлива меньше суммы топлива необходимого на перемещение
	if sumGas < sumCost {
		return -1
	}

	tank := 0           // Количество топлива в баке
	iterationIndex := 0 // Индекс итерации по слайсу
	startIndex := 0     // Индекс элемента с которого возможно начать путь

	// Прерываю цикл когда все элементы перебраны
	for len(gas) > iterationIndex {
		// Количество топлива в баке равно оставшееся топливо полюс топлива наи итерируемой станции
		// минус количество топлива необходимое чтобы добраться до следующей станции.
		tank += gas[iterationIndex] - cost[iterationIndex]
		// Если расчетное количество топлива в баке меньше нуля
		if tank < 0 {
			// Обнулить количество топлива в баке
			tank = 0
			// Начать путь со следующего индекса
			startIndex = iterationIndex + 1
		}
		// Инкрементировать индекс итерации
		iterationIndex++
	}

	// Вернуть индекс начала пусти
	return startIndex
}

// candy
//
// 135. Candy
// There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
//
// You are giving candies to these children subjected to the following requirements:
//
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// Return the minimum number of candies you need to have to distribute the candies to the children.
//
// Example 1:
//
// Input: ratings = [1,0,2]
// Output: 5
// Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
// Example 2:
//
// Input: ratings = [1,2,2]
// Output: 4
// Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
// The third child gets 1 candy because it satisfies the above two conditions.
//
// Constraints:
//
// n == ratings.length
// 1 <= n <= 2 * 104
// 0 <= ratings[i] <= 2 * 104
func candy(ratings []int) int {

	// Length содержит длину массива ratings
	length := len(ratings)

	// Создаем слайс для подсчета и сравнения полученных конфет
	candies := make([]int, length)

	// Раздаем по конфете
	// Условие "Each child must have at least one candy."
	for i, _ := range candies {
		candies[i] = 1
	}

	// Обходим слайс слева направо до предпоследнего элемента
	for i := range length - 1 {
		// Если рейтинг текущего ребенка меньше рейтинга ребенка справа
		if ratings[i] < ratings[i+1] {
			// Ребенок справа получает на одну конфету больше чем есть у текущего ребенка
			candies[i+1] = candies[i] + 1
		}
	}

	// Обходим слайс справа налево до второго элемента
	for i := length - 1; i > 0; i-- {
		// Если у текущего ребенка рейтинг ниже чему у ребенка слева
		// и у конфет у текущего ребенка больше или столько же как у ребенка слева
		if ratings[i] < ratings[i-1] && candies[i] >= candies[i-1] {
			// У ребенка слева становится конфет больше на одну чем у текущего
			candies[i-1] = candies[i] + 1
		}
	}

	// Возвращаем сумму конфет
	sum := 0
	for _, n := range candies {
		sum += n
	}
	return sum
}

// trap
//
// 42. Trapping Rain Water
//
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
//
// Example 1:
//
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
//
// Example 2:
//
// Input: height = [4,2,0,3,2,5]
// Output: 9
//
// Constraints:
//
// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105//
func trap(height []int) int {

	// length содержит длину слайса height
	length := len(height)

	// Далее инициализируются переменные ссылающиеся на первый элемент слайса
	// и на последний, надо проверить что эти элементы есть чтобы не было паники
	if length == 0 {
		return 0
	}

	// left, right Указатели на текущие элементы справа и слева
	left, right := 0, length-1

	// leftMax, rightMax обнаруженные максимумы слева и справа
	leftMax, rightMax := height[left], height[right]

	// water Общее количество воды
	water := 0

	// Обходим слайс пока не сравняются указатели
	for left < right {

		// Если текущая значение слева меньше текущего значения справа
		if height[left] < height[right] {
			// Если текущее значение слева больше или равно обнаруженного на текущий момент
			// максимума слева
			if height[left] >= leftMax {
				// Обнаруженный максимум слева равен ткущему элементу
				leftMax = height[left]
			} else { // Если текущий элемент слева меньше чем текущий элемент справа
				// Повышаем общий уровень воды на разность обнаруженного максимума
				// слева и значения текущего элемента слева
				water += leftMax - height[left]
			}
			// Увеличиваем указатель на левый элемент
			left++
		} else { // Если правый элемент меньше или равно левый элемент
			// Если правый элемент больше обнаруженного максимума справа
			if height[right] >= rightMax {
				// Обнаруженный максимум справа равен текущему элементу справа
				rightMax = height[right]
			} else { // Если текущий элемент справа меньше обнаруженного максимума справа
				// Общий уровень воды повышаем на разность обнаруженного максимума справа и значение элемента справа
				water += rightMax - height[right]
			}
			// Увеличиваем указатель правого элемента
			right--
		}
	}
	// Возвращаем общий уровень воды
	return water
}

// romanToInt
//
// 13. Roman to Integer
//
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.//
func romanToInt(s string) int {

	// Создаю карту соответствия римской цифры и десятичного числа
	m := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Значение текущей римской цифры зависит от предыдущей
	// previous содержит предыдущее число
	previous := 0

	// Sum содержит сумму
	sum := 0

	// Обходим строку циклом range при таком обходе текущий элемент это руна
	for i, _ := range s {
		// Обходить слайс нужно в обратном порядке поэтому вычисляем реверсивней указатель
		reverseI := len(s) - i - 1

		// Если число из карты по текущей руне строки меньше предыдущего числа
		if m[s[reverseI]] < previous {
			// Вычитаем это число из общей суммы
			sum -= m[s[reverseI]]
		} else { // Если число из карты по текущей руне строки не меньше предыдущего числа
			// Увеличиваем общую сумму на это число
			sum += m[s[reverseI]]
		}
		// Обновляем предыдущее число для следящей итерации цикла
		previous = m[s[reverseI]]
	}
	// Возвращаем сумму
	return sum
}

// intToRoman
//
// 12. Integer to Roman
//
// Seven different symbols represent Roman numerals with the following values:
//
// Symbol      Value
// I     1
// V     5
// X     10
// L     50
// C     100
// D     500
// M     1000
// Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:
//
// If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the // remainder to a Roman numeral.
// If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 // (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
// Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a // symbol 4 times use the subtractive form.
// Given an integer, convert it to a Roman numeral.
//
// Example 1:
//
// Input: num = 3749
//
// Output: "MMMDCCXLIX"
//
// Explanation:
//
// 3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
//
//	700 = DCC as 500 (D) + 100 (C) + 100 (C)
//	 40 = XL as 10 (X) less of 50 (L)
//	  9 = IX as 1 (I) less of 10 (X)
//
// Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
// Example 2:
//
// Input: num = 58
//
// Output: "LVIII"
//
// Explanation:
//
// 50 = L
//
//	8 = VIII
//
// Example 3:
//
// Input: num = 1994
//
// Output: "MCMXCIV"
//
// Explanation:
//
// 1000 = M
//
//	900 = CM
//	 90 = XC
//	  4 = IV
//
// Constraints:
//
// 1 <= num <= 3999//
func intToRoman(num int) string {

	// Инициировать буфер для сбора римских цифр
	b := bytes.Buffer{}

	// Повторять цикл пока num > 0
	for num > 0 {
		// В конструкции switch будем определять какое максимально возможное число мы можем вычесть
		// В каждом случае, добавляем соответствующую римскую цифру и уменьшаем num на соответствующее число
		// пока от num ничего не останется
		switch {
		case num >= 1000:
			b.WriteString("M")
			num -= 1000
		case num >= 900:
			b.WriteString("CM")
			num -= 900
		case num >= 500:
			b.WriteString("D")
			num -= 500
		case num >= 400:
			b.WriteString("CD")
			num -= 400
		case num >= 100:
			b.WriteString("C")
			num -= 100
		case num >= 90:
			b.WriteString("XC")
			num -= 90
		case num >= 50:
			b.WriteString("L")
			num -= 50
		case num >= 40:
			b.WriteString("XL")
			num -= 40
		case num >= 10:
			b.WriteString("X")
			num -= 10
		case num >= 9:
			b.WriteString("IX")
			num -= 9
		case num >= 5:
			b.WriteString("V")
			num -= 5
		case num >= 4:
			b.WriteString("IV")
			num -= 4
		case num >= 1:
			b.WriteString("I")
			num -= 1
		}
	}
	// Выводим строку полученную из буфера
	return b.String()
}

// lengthOfLastWord
//
// 58. Length of Last Word
// Given a string strs consisting of words and spaces, return the length of the last word in the string.
//
// A word is a maximal
// substring
//
//	consisting of non-space characters only.
//
// Example 1:
//
// Input: strs = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:
//
// Input: strs = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
// Example 3:
//
// Input: strs = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.
//
// Constraints:
//
// 1 <= strs.length <= 104
// strs consists of only English letters and spaces ' '.
// There will be at least one word in strs.//
func lengthOfLastWord(s string) int {

	// Далее указатель цикла будет инициирован длиной строки минус один
	// что приведет к панике, поэтому добавляю проверку.
	if len(s) == 0 {
		return 0
	}

	// Счетчик символов
	charCount := 0

	// Обходим строку в справа налево
	for i := len(s) - 1; i >= 0; i-- {
		// Если текущий элемент пробел
		if s[i] == 0x20 {
			// Если мы уже что-то посчитали
			if charCount > 0 {
				// Возвращаем счетчик символов
				return charCount
			}
		} else { // Если не пробел
			// Счетчик символов увеличиваем на один
			charCount++
		}
	}

	// Эта часть кода сработает если строка состоит только из пробелов
	return charCount
}

// longestCommonPrefix
//
// 14. Longest Common Prefix
//
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
// Constraints:
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters if it is non-empty.//
func longestCommonPrefix(strs []string) string {

	// Далее minLength инициируется длиной первой строки в strs
	// необходимо проверить что этот элемент есть
	if len(strs) == 0 {
		return ""
	}

	// Инициируем переменную minLength длиной первого элемента strs
	minLength := len(strs[0])

	// Обходим strs вычисляем минимум длины среди элементов strs
	for i := 1; i < len(strs); i++ {
		minLength = min(minLength, len(strs[i]))
	}

	// Если минимальная длина 0 возвращаем пустую сроку
	if minLength == 0 {
		return ""
	}

	// Тут крайний индекс префикса
	prefixCharsCounter := 0

	// rootLoop: это метка так как цикл может быть прерван из внутреннего цикла
rootLoop:
	// Цикл по диапазону [0..minLength-1]
	for i := range minLength {
		// Текущий байт первого элемента будет эталоном
		// с которым будем сравнивать остальные байты элементов strs
		reference := strs[0][i]
		// Обходим strs
		for j := 1; j < len(strs); j++ {
			// Если текущий байт в текущем элементе strs не совпадает с эталоном
			if reference != strs[j][i] {
				// Перерываем внешний цикл rootLoop
				break rootLoop
			}
		}
		// Актуализируем крайний индекс префикса
		prefixCharsCounter++
	}

	// Возвращаем префикс первого элемента strs
	return strs[0][0:prefixCharsCounter]
}

// reverseWords
//
// 151. Reverse Words in a String
//
// Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra // spaces.
//
// Example 1:
//
// Input: s = "the sky is blue"
// Output: "blue is sky the"
// Example 2:
//
// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:
//
// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
//
// Constraints:
//
// 1 <= s.length <= 104
// s contains English letters (upper-case and lower-case), digits, and spaces ' '.
// There is at least one word in s.
func reverseWords(s string) string {

	// Разбиваем строку на подстроки по условию
	words := strings.Fields(s)

	// Меняем порядок элементов в слайсе классическим методом
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Возвращаем строку, полученную после обледенения элементов слайча через пробел
	return strings.Join(words, " ")
}

// convert
//
// 6. Zigzag Conversion
//
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// string convert(string s, int numRows);
//
// Example 1:
//
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Example 2:
//
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// Example 3:
//
// Input: s = "A", numRows = 1
// Output: "A"
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000//
func convert(s string, numRows int) string {

	// Если строка одна возвращаем ее сразу
	if numRows == 1 {
		return s
	}

	// Создаем слайс строк длиной numRows
	rows := make([]string, numRows)
	// Указательна текущую строку
	currentRowIndex := 0
	// Флаг движения вниз
	directionDown := true

	// Цикл по строке
	for i := range len(s) {

		// Если указатель строк на первой строке
		if currentRowIndex == 0 {
			// Подымаем флаг движения вниз
			directionDown = true
		}

		// Если указатель строк на крайней строке
		if currentRowIndex == len(rows)-1 {
			// Опускаем флаг движения вниз
			directionDown = false
		}

		// Добавляем символ в строку на которой указатель текущей строки
		rows[currentRowIndex] += string(s[i])

		// Если поднят флаг движения вниз
		if directionDown {
			// Увеличиваем указатель текущей строки
			currentRowIndex++
		} else { // Если не поднят флаг движения вниз
			// Уменьшаем указатель текущей строки
			currentRowIndex--
		}
	}

	// Возвращаем строку из объединенного слайса строк
	return strings.Join(rows, "")
}

// strStr
//
// 28. Find the Index of the First Occurrence in a String
//
// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
// Example 1:
//
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.
// Example 2:
//
// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
// Constraints:
//
// 1 <= haystack.length, needle.length <= 104
// haystack and needle consist of only lowercase English characters.
func strStr(haystack string, needle string) int {

	if haystack == needle {
		return 0
	}

	// Сдвиг
	needleIndex := 0

	// Обходим строку haystack
	for i := range len(haystack) {

		if haystack[i] != needle[needleIndex] {
			needleIndex = 0
			continue
		}

		needleIndex++

		if needleIndex == len(needle) {
			return i - needleIndex + 1
		}
	}

	return -1
}

// 69. Sqrt(x)
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.
//
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
//
//
// Example 1:
//
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.
// Example 2:
//
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
//
//
// Constraints:
//
// 0 <= x <= 231 - 1

func mySqrt(x int) int {
	xFloat := float64(x)
	sqrtX := math.Sqrt(xFloat)
	flooredX := math.Floor(sqrtX)
	return int(flooredX)
}
