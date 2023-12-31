package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"learn/bluebell/logic"
	"learn/bluebell/models"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	//c.Query()
	//c.Param()  路径参数的校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUP with invalid param", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
		})
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.RePassword) == 0 || len(p.Password) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUP with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(p)
	ResponseSuccess(c, nil)
	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "注册失败"})
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
