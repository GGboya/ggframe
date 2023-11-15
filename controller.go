package main

import "ggframe/framework"

func UserLoginController(c *framework.Context) error {
	// 打印控制器的名字
	c.Json(200, "ok, UserLoginController")
	return nil
}

func SubjectAddController(c *framework.Context) error {
	// 打印控制器的名字
	c.Json(200, "ok, SubjectAddController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	// 打印控制器的名字
	c.Json(200, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	// 打印控制器的名字
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	// 打印控制器的名字
	c.Json(200, "ok, SubjectGetController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	// 打印控制器的名字
	c.Json(200, "ok, SubjectListController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	// 打印控制器的名字
	c.Json(200, "ok, SubjectNameController")
	return nil
}
