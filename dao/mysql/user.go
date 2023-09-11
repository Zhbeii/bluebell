package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"learn/bluebell/models"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "xiaoll"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	// 知识点：最后这里为什么不能返回 err ，因为 err 是 if 里面的，在 if 之外是不可见的
	return
}

// InsertUser 数据库添加用户
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行 SQL 语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
