# 155. Min Stack

**Difficulty:** Medium  
**Topic:** Stack, Design

## Problem

Реализуй стек который поддерживает push, pop, top и получение минимального элемента за O(1).

## Methods

- `NewMinStack() *MinStack` — инициализация
- `Push(val int)` — добавить элемент
- `Pop()` — удалить верхний элемент
- `Top() int` — вернуть верхний элемент
- `GetMin() int` — вернуть минимальный элемент

## Example

push(-2), push(0), push(-3)
getMin() → -3
pop()
top()    → 0
getMin() → -2

## Constraints

- -2³¹ <= val <= 2³¹ - 1
- pop, top, getMin вызываются только на непустом стеке
- не более 3 * 10⁴ вызовов

## Approach

Подсказка: два стека.
