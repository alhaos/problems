package solvit

import "strings"

func decodeString(s string) string {
	// Два стека: один для строк, другой для чисел
	strStack := []string{}
	numStack := []int{}

	curStr := ""
	curNum := 0

	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '9':
			// Накапливаем число (может быть многозначным: 12[a])
			curNum = curNum*10 + int(ch-'0')

		case ch == '[':
			// Сохраняем текущий контекст и начинаем новый
			strStack = append(strStack, curStr)
			numStack = append(numStack, curNum)
			curStr = ""
			curNum = 0

		case ch == ']':
			// Извлекаем контекст и применяем повтор
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			curStr = prevStr + strings.Repeat(curStr, num)

		default:
			curStr += string(ch)
		}
	}

	return curStr
}
