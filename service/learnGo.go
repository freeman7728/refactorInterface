package service

import (
	"refactorInterface/client"
)

type ICourseService interface {
}
type CourseServiceImpl struct {
	cli client.CourseClient
}

func NewCourseService(cli client.CourseClient) *CourseServiceImpl {
	//使得cli不再是需要频繁创建的对象了，节省了资源
	return &CourseServiceImpl{cli: cli}
}

func (c *CourseServiceImpl) LearnGo() {
	//fmt.Println("learnGo")
	//屏蔽底层细节
	//cli := client.NewClient()
	//cli.LearnGo()
	//使得cli不再是需要频繁创建的对象了，节省了资源
	c.cli.LearnGo()
}
