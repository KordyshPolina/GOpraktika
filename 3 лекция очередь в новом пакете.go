
package queue

type Queue struct {
    s    []interface{} // слайс в котором хранятся значения
    low  int           // индекс нижней границы очереди
    high int           // индекс верхней границы очереди
    size int           // размер очереди
}

func NewQueue(size int) *Queue {
    return &Queue{
        s:    make([]interface{}, size),
        size: size,
        low:  -1,
        high: -1,
    }
}

// Push - добавление в очередь значения
func Push(q *Queue, v interface{}) {
    if q.high == q.size-1 {
        return
    }
    q.high++
    q.s[q.high] = v
    if q.low == -1 {
        q.low = 0
    }
}

// Pop - получение значения из очереди и его удаление
func Pop(q *Queue) interface{} {
    if q.low == -1 || q.low > q.high {
        return nil
    }
    v := q.s[q.low]
    q.low++
    return v
}

