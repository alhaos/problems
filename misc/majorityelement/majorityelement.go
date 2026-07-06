package majorityelement

// MajorityElement возвращает элемент, встречающийся в nums более n/2 раз.
// Гарантируется, что такой элемент существует.
//
// Требования: O(n) по времени, O(1) по памяти.
func MajorityElement(nums []int) int {

	// кандидат на
	candidate := 0
	counter := 0

	for _, n := range nums {
		if counter == 0 {
			candidate = n
			counter++
			continue
		}

		if n == candidate {
			counter++
		} else {
			counter--
		}
	}
	return candidate
}
