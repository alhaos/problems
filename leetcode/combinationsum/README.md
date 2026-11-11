# 39. Combination Sum

Дан массив уникальных целых `candidates` и целое `target`.
Верни все уникальные комбинации из `candidates`, сумма которых равна `target`.
Одно и то же число можно использовать **неограниченное количество раз**.
Две комбинации разные, если отличается количество хотя бы одного числа.
Порядок комбинаций и порядок чисел внутри комбинации любой.

## Примеры

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Input: candidates = [2], target = 1
Output: []

## Ограничения

- 1 <= len(candidates) <= 30
- 2 <= candidates[i] <= 40
- все candidates уникальны
- 1 <= target <= 40

## Задача

Реализуй `combinationSum` в `combination_sum.go`, прогони `go test`.
Решай через backtracking.
