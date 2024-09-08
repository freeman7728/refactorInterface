package client

import "fmt"

type CourseClient interface {
	LearnGo()
	LearnJava()
	LearnPiano()
}

type CourseClientImpl struct {
}

func (client *CourseClientImpl) LearnGo() {
	fmt.Println("Learn Go")
}
func (client *CourseClientImpl) LearnJava() {
	fmt.Println("Learn Java")
}
func (client *CourseClientImpl) LearnPiano() {
	fmt.Println("Learn Piano")
}
func NewClient() CourseClient {
	return &CourseClientImpl{}
	//虽然返回的是一个指针，但是编译器自动将其视为实现了接口的类型
	//所以最后返回是一个CourseClient
}
