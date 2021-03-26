package superReducedString

type Stack struct {
	stack []byte
}

func (s *Stack) Push(b byte) {
	s.stack = append(s.stack, b)
}

func (s *Stack) Pop() byte {
	res := s.stack[len(s.stack)-1]

	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *Stack) Peek() byte {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) String() string {
	return string(s.stack)
}

func superReducedStringStack(s string) string {
	stack := new(Stack)
	stack.Push(s[0])
	for i := 1; i < len(s); i++ {
		if stack.Size() > 0 && stack.Peek() == s[i] {
			stack.Pop()
			continue
		}
		stack.Push(s[i])
	}
	if stack.String() == "" {
		return "Empty String"
	}
	return stack.String()
}

func superReducedString(s string) string {
	for i, l := 0, len(s); i < l-1; {
		if s[i] == s[i+1] {
			s = s[:i] + s[i+2:]
			l = len(s)
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	if s == "" {
		return "Empty String"
	}
	return s
}
