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
}
