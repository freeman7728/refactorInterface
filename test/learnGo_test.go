package test

import (
	"fmt"
	"refactorInterface/service"
	"testing"
)

type MockClient struct {
}

// 如果Client方法过多，则导致进行单元测试的时候，需要大量无用代码来实现接口，否则无法通过编译
func (client *MockClient) LearnGo() {
	fmt.Println("Learn Go")
}
func (client *MockClient) LearnJava() {
	fmt.Println("Learn Java")
}
func (client *MockClient) LearnPiano() {
	fmt.Println("Learn Piano")
}
func NewMockClient() *MockClient {
	return &MockClient{}
}

func Test_CourseServiceImpl_LearnGo(t *testing.T) {
	mockClient := NewMockClient()
	mockService := service.NewCourseService(mockClient)
	mockService.LearnGo()
}
