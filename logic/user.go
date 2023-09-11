package logic

import (
	"learn/bluebell/dao/mysql"
	"learn/bluebell/models"
	"learn/bluebell/pkg/snowflake"
)

// 存放业务的逻辑代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 0. 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 1. 生成 UID
	userID := snowflake.GenID()

	// 构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 2. 保存进数据库
	return mysql.InsertUser(user)
}
