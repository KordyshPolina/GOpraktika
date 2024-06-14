
package stack

type Stack struct {
    s    []interface{}
    head int
}

func NewStack(size int) *Stack {
    return &Stack{
        s:    make([]interface{}, size),
        head: -1,
    }
}

func Push(s *Stack, v interface{}) {
    s.head++
    s.s[s.head] = v
}

func Pop(s *Stack) interface{} {
    if s.head < 0 {
        return nil
    }
    v := s.s[s.head]
    s.head--
    return v
}

func Peek(s *Stack) interface{} {
    if s.head < 0 {
        return nil
    }
    return s.s[s.head]
}
