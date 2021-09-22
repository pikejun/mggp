package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pikejun/mggp/global"
	"github.com/pikejun/mggp/model"
	"github.com/pikejun/mggp/pkg/app"
	"log"
	"strconv"
)

// 路由的 Handler
type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

// 定义方法
// 增
func (u *UserController) Create(c *gin.Context) {
	param := model.User{}
	response := app.NewResponse(c)
	// 入参校验和绑定
	valid, err := app.BindAndValid(c, &param)
	if valid {
		log.Println(err)
		response.ToResponseFailed(-1,err.Error())
		return
	}

	err = param.Create(global.DBEngine)

	if err != nil {
		response.ToResponseFailed(-1,err.Error())
		return
	}


	response.ToResponse(gin.H{})
}

// 删
func (u *UserController) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	s := c.Param("id")
	fmt.Println("id=",s)

	param:=model.User{}
	i,_:=strconv.Atoi(s)
	param.ID=uint64(i)
	param.DelteById(global.DBEngine)
	response.ToResponse(gin.H{
	})
}

//改
func (u *UserController) Update(c *gin.Context) {
	param := model.User{}
	response := app.NewResponse(c)
	// 入参校验和绑定
	valid, err := app.BindAndValid(c, &param)
	if valid {
		log.Println(err)
		response.ToResponseFailed(-1,err.Error())
		return
	}

	s := c.Param("id")
	fmt.Println("id=",s)
	i,_:=strconv.Atoi(s)
	param.ID=uint64(i)

	err = param.Update(global.DBEngine)
	if err != nil {
		response.ToResponseFailed(-1,err.Error())
	}

	response.ToResponse(gin.H{})
}

// 查
func (u *UserController) GetUserById(c *gin.Context) {
	response := app.NewResponse(c)
	s := c.Param("id")
	fmt.Println("id=",s)
	param := model.User{}
	i,_:=strconv.Atoi(s)
	param.ID=uint64(i)
	response.ToResponse(param.GetById(global.DBEngine))
}

// 查询列表
func (u *UserController) ListUserByPage(c *gin.Context) {
	param :=  model.User{}
	response := app.NewResponse(c)
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	fmt.Println("pageSize=",pageSize)
	pageNo,_ := strconv.Atoi(c.Query("pageNo"))
	fmt.Println("pageNo=",pageNo)

	rows,_:=param.ListUserByPage(global.DBEngine,pageSize,pageNo)
	response.ToResponse(rows)
}
