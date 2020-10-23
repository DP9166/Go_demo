package main

import "fmt"

type Student struct {
	id uint
	name string
	male bool
	score float64
}

func NewStudent(id uint, name string,  score float64) *Student {
	return &Student{id:id, name:name, score:score}
}

func (s Student) GetName() string  {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s Student) String() string  {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}", s.id, s.name, s.male, s.score)
}


// ------------------------------  为基本数据类型添加成员方法
type Integer int

func (a Integer) Equal (b Integer) bool {
	return a == b
}


func main()  {
	var a Integer = 2
	if a.Equal(2) {
		fmt.Println(a, "is equal to 2")
	}
}