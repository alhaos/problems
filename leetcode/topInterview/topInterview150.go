// https://leetcode.com/studyplan/top-interview-150/

package topInterview

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
