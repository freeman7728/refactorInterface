package main

import (
	"refactorInterface/client"
	"refactorInterface/service"
)

func main() {
	c := client.NewClient()
	LearnService := service.NewCourseService(c)
	LearnService.LearnGo()
}
