package main

import (
	"refactorInterface/client"
	"refactorInterface/service"
)

func main() {
	c := client.NewClient()
	//此时的c拥有CourseClient的全部方法
	//以值传递的形式交给service后，service只接受proxy中声明的方法，所以不能使用指针传递
	LearnService := service.NewCourseService(c)
	LearnService.Learn()
}
