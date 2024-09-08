package test

import (
	"fmt"
	"refactorInterface/service"
	"testing"
)

type MockClientProxy interface {
	LearnGo()
}

type MockClient struct {
	Cli MockClientProxy
}

//同时也方便了单侧的实现，无需实现原先接口中的所有定义方法

// 如果Client方法过多，则导致进行单元测试的时候，需要大量无用代码来实现接口，否则无法通过编译
func (client *MockClient) LearnGo() {
	fmt.Println("Learn Go")
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func Test_CourseServiceImpl_LearnGo(t *testing.T) {
	mockClient := NewMockClient()
	mockService := service.NewCourseService(mockClient)
	mockService.Cli.LearnGo()
}
