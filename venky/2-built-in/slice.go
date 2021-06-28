package main

import "fmt"

type slice struct {
	length int
	data   map[int]interface{}
}

type isArray interface {
	get(index int) interface{}
	size() int
	add(index int, data interface{})
	push(data interface{})
	pop()
	delete(index int)
}

func NewSlice() *slice {
	return &slice{
		length: 0,
		data:   map[int]interface{}{},
	}
}

func (s slice) String() string {
	return fmt.Sprint("length:", s.length, "\ndata: ", s.data)
}

func main() {
	slice := NewSlice()
	slice.push("toi")
	slice.push("la")
	slice.push("vuong")

	slice.size()
	fmt.Println(slice)
	slice.add(2, "khanh")
	slice.push("day")
	fmt.Println(slice)
	slice.pop()
	fmt.Println(slice)
	slice.delete(0)
	fmt.Println(slice)

}

func (s *slice) get(index int) interface{} {
	if s.length < index {
		return nil
	}

	return s.data[index]
}

func (s *slice) size() int {
	return s.length
}

func (s *slice) add(index int, data interface{}) {
	for i := s.length; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = data
	s.length++
}

func (s *slice) push(data interface{}) {
	s.data[s.length] = data
	s.length++
}

func (s *slice) pop() {
	delete(s.data, s.length-1)
	s.length--
}

func (s *slice) delete(index int) {
	delete(s.data, index)
	unShift(s, index)
	s.length--
}

func unShift(s *slice, index int) {
	for i := index; i < s.length; i++ {
		s.data[i] = s.data[i+1]
	}
	delete(s.data, s.length-1)
}
