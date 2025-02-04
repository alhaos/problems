// https://leetcode.com/studyplan/top-interview-150/

package topInterview

import (
	"math"
	"math/rand"
	"sort"
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
// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
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
// It's guaranteed that you can reach nums[n - 1].
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
// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.
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

// RandomizedSet 380. Insert Delete GetRandom O(1)
// Implement the RandomizedSet class:
//
// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
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
