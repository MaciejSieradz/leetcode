package main

func main() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([])"))
}

func isValid(s string) bool {

	var stack Stack

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack.Push(char)
		} else {
			if char == ')' && stack.Top() != '(' {
				return false
			}
			if char == ']' && stack.Top() != '[' {
				return false
			}
			if char == '}' && stack.Top() != '{' {
				return false
			}
			stack.Pop()
		}
	}

	if !stack.IsEmpty() {
		return false
	}

	return true
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() rune {
	if s.IsEmpty() {
		return 0
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}

	return false
}
