На входе строка содержащая десять цифр от нуля до восьми разделенная запитой и пробелом. "7, 0, 5, 4, 2, 3, 1, 6, 8"

Это массив  ```[int[]]@(7, 0, 5, 4, 2, 3, 1, 6, 8) ```, элементами которого являются ходы в игре "крестики - нолики", позиция элемента означает номер хода, значение клетка в которорую она сделан, клетки нумеруются слево направо, сверху вниз. Игроки ходят по очереди, нечетные ходы крестики, четные нолики.

```
 O | X | X
-----------
 O | O | X
-----------
 O | X | X
```

Определить яляется ли комбинатся выигрошной, если да то на каком ходу и какой игрок победил
На выходе, ("First player win at [] turn"), ("Second player win at [] turn"), ("Draw")