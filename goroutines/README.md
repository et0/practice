# Горутины
| # | Description | Tags |
| --- | --- | --- |
| 001 | Простой вывод чисел в горутине | WaitGroup |
| 002 | Передача данных через канал | chan, nobuf |
| 003 | Синхронизация двух горутин (Pipeline) | chan, nobuf |
| 004 | Работа с несколькими горутинами (Worker Pool) | chan, buf, wp |
| 005 | Конкурентный доступ к данным (Mutex) | mutex, race condition |
| 006 | Pipeline с фильтрацией | deadlock |
| 007 | Остановка горутин по сигналу (Context) | context |
| 008 | Fan-out, Fan-in | fan-in, fan-out, chan, buf |
| 009 | Ограничение скачивания файлов (Semaphore) | chan, buf, semaphore |
| 010 | Ограничение запросов к API (Semaphore) | chan, buf, semaphore |
| 011 | Параллельная обработка данных (Worker Pool, Errgroup) | chan, buf, wp, errg |
| 012 | Ограниченный пул рабочих (Worker Pool) с ожиданием задач | sync.Cond, wp |
| 013 | Одна горутина (main) пишет в канал, две горутины читают из канала числа | chan, nobuf, WaitGroup | 