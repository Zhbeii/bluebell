package logic

import (
	"learn/bluebell/dao/mysql"
	"learn/bluebell/pkg/snowflake"
)

// 存放业务的逻辑代码

func SignUp() {
	// 0. 判断用户是否存在
	mysql.QueryUserByUserName()
	// 1. 生成 UID
	snowflake.GenID()
	// 密码加密

	// 2. 保存进数据库
	mysql.InsertUser()
}
