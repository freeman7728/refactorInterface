package client

import "fmt"

// CourseClient 直接创建类，而不是接口
type CourseClient struct {
}

func (client *CourseClient) LearnGo() {
	fmt.Println("Learn Go")
}
func (client *CourseClient) LearnJava() {
	fmt.Println("Learn Java")
}
func (client *CourseClient) LearnPiano() {
	fmt.Println("Learn Piano")
}

func NewClient() *CourseClient {
	return &CourseClient{}
}
