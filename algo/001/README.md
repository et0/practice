# 001. Meeting rooms (жадные алгоритмы, приоритетные очереди)

**Ссылка на решение:**
https://github.com/et0/leetcode/tree/master/Hard/2402 

**Условие задачи:**  
Вам дано целое число `n`. Существует `n` комнат, пронумерованных от `0` до `n - 1`.  

Также дан двумерный массив встреч `meetings`, где `meetings[i] = [starti, endi]` означает, что встреча будет проходить в полуинтервале времени `[starti, endi)`. Все значения `starti` уникальны.  

**Правила распределения встреч по комнатам:**  
1. Каждая встреча должна проводиться в **свободной комнате с наименьшим номером**.  
2. Если свободных комнат нет, встреча **откладывается** до момента освобождения любой комнаты. При этом длительность отложенной встречи остается такой же, как у исходной.  
3. Когда комната освобождается, **первой должна быть проведена встреча с самым ранним исходным временем начала** (т.е. приоритет у встреч, которые изначально должны были начаться раньше).  

**Требуется:**  
Вернуть номер комнаты, в которой было проведено **наибольшее количество встреч**. Если таких комнат несколько, вернуть комнату с **наименьшим номером**.  

**Примечание:**  
Полуинтервал `[a, b)` включает `a`, но не включает `b`.  

### Что нужно сделать?  
Необходимо реализовать алгоритм, который:  
1. Моделирует процесс распределения встреч по комнатам согласно указанным правилам.  
2. Считает, сколько встреч было проведено в каждой комнате.  
3. Находит комнату с максимальным числом встреч (если таких несколько, выбирает комнату с меньшим номером).  

### Пример (для понимания):  
Допустим, `n = 2` (комнаты `0` и `1`), а `meetings = [[0, 10], [1, 5], [2, 7], [3, 4]]`.  

1. Встреча `[0, 10]` занимает комнату `0` (самую маленькую свободную).  
2. Встреча `[1, 5]` занимает комнату `1`.  
3. Встреча `[2, 7]` не может начаться, так как обе комнаты заняты. Она откладывается.  
4. Встреча `[3, 4]` также откладывается.  
5. В момент `5` освобождается комната `1` (завершается встреча `[1, 5]`).  
   - Из отложенных встреч `[2, 7]` и `[3, 4]` раньше должна начаться `[2, 7]` (по исходному времени начала).  
   - Она занимает комнату `1` и продлится до `5 + (7 - 2) = 10`.  
6. В момент `4` освобождается комната `0` (но встреча `[0, 10]` еще идет, поэтому это неверно — здесь нужна более точная симуляция).  

В итоге правильное распределение требует аккуратного моделирования с приоритетными очередями.  

**Ответ:** `0` (если провести симуляцию корректно).  

### Ключевые моменты:  
- Использовать **мини-кучу (min-heap)** для отслеживания свободных комнат.  
- Использовать **другую кучу** для отслеживания завершающихся встреч (чтобы обрабатывать освобождение комнат).  
- Учитывать **отложенные встречи**, сортируя их по исходному времени начала.  

---

### **Полное решение задачи о встречах в комнатах на Go**

Используем **две кучи**:
1. **`freeRooms`** — min-heap с номерами свободных комнат (чтобы брать комнату с минимальным номером).
2. **`busyRooms`** — min-heap по времени окончания встречи (чтобы находить ближайшее освобождение).

#### **Алгоритм:**
1. **Инициализируем** все комнаты как свободные.
2. **Сортируем встречи** по времени начала.
3. **Обрабатываем каждую встречу**:
   - Освобождаем комнаты, у которых встреча завершилась.
   - Если есть свободная комната — назначаем встречу.
   - Если нет — откладываем встречу (увеличиваем её время начала и окончания).
4. **Считаем статистику** по комнатам.

---

### **Ключевые моменты:**
1. **Сортировка встреч** по `startTime` гарантирует обработку в хронологическом порядке.
2. **`busyRooms`** отсортирована по `endTime`, чтобы быстро находить ближайшее освобождение.
3. **Отложенные встречи** получают новое время `earliestEnd + duration`.
4. **Статистика** ведётся в `roomUsage`, чтобы найти комнату с максимумом встреч.

---

### **Сложность:**
- Время: **O(m log n)**, где `m` — количество встреч, `n` — количество комнат.
- Память: **O(n)** для хранения куч.

Это оптимальное решение с использованием приоритетных очередей.