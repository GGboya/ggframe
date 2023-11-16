package main

import (
	"ggframe/framework"
	"ggframe/framework/middleware"
	"time"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// 需求1+2：HTTP方法+静态路由匹配
	core.Get("/user/login", middleware.Test3(), UserLoginController)
	// 需求3：批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", framework.TimeoutHandler(SubjectNameController, time.Second))
		}

		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
	}

}
