// https://programmers.co.kr/learn/courses/30/lessons/42883

package main

// Stack type init
type Stack struct {
	Data  []byte
	Count int
}

// Push Stack insert byte type Data
func (s *Stack) Push(a byte) {
	s.Data = append(s.Data[:s.Count], a)
	s.Count++
}

// Pop Stack remove Top Data
func (s *Stack) Pop() {
	s.Count--
}

// Top Stack Return Top Data
func (s *Stack) Top() byte {
	return s.Data[s.Count-1]
}

func p42883(n string, k int) string {
	var stack Stack

	str := []byte(n)
	for _, v := range str {
		for stack.Count > 0 && stack.Top() < v && k > 0 {
			stack.Pop()
			k--
		}
		stack.Push(v)
	}
	for ; k > 0; k-- {
		stack.Pop()
	}

	return string(stack.Data[:stack.Count])
}
