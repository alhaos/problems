# Course Schedule

Нужно пройти `numCourses` курсов, пронумерованных от `0` до `numCourses-1`.

Дан список `prerequisites`, где `prerequisites[i] = [a, b]` значит:
чтобы пройти курс `a`, надо сначала пройти курс `b`.

Верни `true`, если можно пройти все курсы, `false` — если это невозможно
(то есть есть циклическая зависимость).

## Примеры

numCourses = 2, prerequisites = [[1,0]]
-> true   (0 -> 1, порядок есть)

numCourses = 2, prerequisites = [[1,0],[0,1]]
-> false  (1 требует 0, 0 требует 1 — цикл)

numCourses = 4, prerequisites = [[1,0],[2,1],[3,2]]
-> true   (0 -> 1 -> 2 -> 3)

numCourses = 3, prerequisites = []
-> true   (зависимостей нет)

## Ограничения
- 1 <= numCourses <= 2000
- 0 <= a, b < numCourses
- a != b
