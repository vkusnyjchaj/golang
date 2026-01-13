# Mutex

Eсли мы хотим убедиться, что один goroutine может получать доступ к переменной без конфликтов, то можно воспользоваться взаимным исключением и структурой данных `mutex`.

Стандартная библиотека Go обеспечивает взаимное исключение с синхронизацией. Mutex и его два метода:
- Lock
- Unlock

Объявление мьютекса:
```
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}
```

```
c.mu.Lock() // lock
// Lock so only one goroutine at a time can access the map c.v.
c.mu.Unlock() // unlock
```