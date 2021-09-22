package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pikejun/mggp/global"
	"github.com/pikejun/mggp/model"
	"github.com/pikejun/mggp/pkg/app"
	"log"
)

type FileController struct {
}

func NewFileController() FileController {
	return FileController{}
}

func (fc *FileController) CreateFile(c *gin.Context){
	param := model.FileModle{}
	response := app.NewResponse(c)
	// 入参校验和绑定
	noValid, err := app.BindAndValid(c, &param)
	if noValid {
		log.Println(err)
		response.ToResponseFailed(-1,err.Error())
		return
	}
	err = param.Create(global.DBEngine)
	if err != nil {
		response.ToResponseFailed(-1,err.Error())
	}

	response.ToResponse(nil)
}


func (fc *FileController) UpdateFile(c *gin.Context){
	param := model.FileModle{}
	response := app.NewResponse(c)
	// 入参校验和绑定
	valid, err := app.BindAndValid(c, &param)
	if valid {
		log.Println(err)
		return
	}

	s := c.Param("id")
	fmt.Println("id=",s)
	param.Id=s

	err = param.Update(global.DBEngine)

	if err != nil {
		response.ToResponseFailed(-1,err.Error())
	}

	response.ToResponse(gin.H{})
}


func (fc *FileController) GetFile(c *gin.Context){
	param := model.FileModle{}
	response := app.NewResponse(c)
	s := c.Param("id")
	fmt.Println("id=",s)
	param.Id=s
	param2 := param.Get(global.DBEngine)
	response.ToResponse(param2)
}

func (f *FileController) DeleteFile(c *gin.Context){
	param := model.FileModle{}
	response := app.NewResponse(c)
	// 入参校验和绑定
	valid, err := app.BindAndValid(c, &param)
	if valid {
		log.Println(err)
		return
	}

	s := c.Param("id")
	fmt.Println("id=",s)
	param.Id=s


	err = param.Delete(global.DBEngine)
	if err != nil {
		response.ToResponseFailed(-1,err.Error())
	}

	response.ToResponse(gin.H{})
}


func (f *FileController) ListFile(c *gin.Context){
	param := model.FileModle{}
	response := app.NewResponse(c)
	// 入参校验和绑定
	valid, err := app.BindAndValid(c, &param)
	if valid {
		log.Println(err)
		return
	}

	s := c.Param("id")
	fmt.Println("id=",s)
	param.Id=s

	//err = param.
	//if err != nil {
	//	return
	//}

	response.ToResponse(gin.H{})
}
