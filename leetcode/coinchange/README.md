# 322. Coin Change

Дан массив `coins` с номиналами монет и целое `amount`.
Верни наименьшее количество монет, которым можно набрать сумму `amount`.
Если набрать нельзя, верни `-1`.
Каждого номинала монет бесконечно много.

## Примеры

Input: coins = [1,2,5], amount = 11
Output: 3
(5 + 5 + 1)

Input: coins = [2], amount = 3
Output: -1

Input: coins = [1], amount = 0
Output: 0

## Ограничения

- 1 <= len(coins) <= 12
- 1 <= coins[i] <= 2^31 - 1
- 0 <= amount <= 10^4

## Задача

Реализуй `coinChange` в `coin_change.go`, прогони `go test`.
Входной слайс менять нельзя.
