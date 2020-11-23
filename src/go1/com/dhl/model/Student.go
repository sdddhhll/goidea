package model

type Student struct {
	Name string
	age byte
	monay float64
}


func NewStudent(name string,age byte,monay float64) *Student {
	return &Student{
		Name:  name,
		age:   age,
		monay: monay,
	}
}

func (stu *Student) SetAge(age byte)  {
	stu.age = age
}

func (stu * Student) GetAge() byte {
	return stu.age
}

func (stu *Student) SetMonay(monay float64){
	stu.monay = monay
}

func (stu *Student) GetMonay() float64 {
	return stu.monay
}

