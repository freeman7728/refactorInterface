package service

type ICourseService interface {
}

// ClientProxy 我这个服务需要使用底层的哪一个组件，就声明哪一个组件
type ClientProxy interface {
	LearnGo()
}

type CourseServiceImpl struct {
	Cli ClientProxy
}

func NewCourseService(cli ClientProxy) *CourseServiceImpl {
	//使得cli不再是需要频繁创建的对象了，节省了资源
	return &CourseServiceImpl{Cli: cli}
}

func (c *CourseServiceImpl) Learn() {
	//fmt.Println("learnGo")
	//屏蔽底层细节
	//cli := client.NewClient()
	//cli.LearnGo()
	//使得cli不再是需要频繁创建的对象了，节省了资源
	c.Cli.LearnGo()
}
