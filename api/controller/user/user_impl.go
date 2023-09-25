package user

import (
	"github.com/gin-gonic/gin"
	"sectran/api/model"
	"sectran/api/service"
)

//// 首先定义你需要操作这个表的方法
//type UserInterface interface {
//	GetUserById(int32) (*model.User, error)
//	DelUserById(int32) (*model.User, error)
//	EditUserById(int32, *model.User) error
//}
//
//// 实现这些方法
//func GetUserById(int32) (*model.User, error) {
//	return new(model.User), nil
//}

func PostUserLogin(c *gin.Context) {
	p := model.UserLogin{}
	err := c.BindJSON(&p)
	if err != nil {
		return
	}
	Db := service.Db
	//if err := Db.Where("user_name = ?", p.UserName).First(&Kp).Error; err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusBadRequest,
	//		"msg":  "无此账号",
	//	})
	//
	//}

	//return nil

}
