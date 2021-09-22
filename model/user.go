package model

import (
	"github.com/pikejun/mggp/pageUtil"
	"gorm.io/gorm"
)

// 创建 user model
type User struct {
	ID         uint64 `form:"id"  binding:"max=10" gorm:"primary_key" json:"id"`
	Name      string `form:"name" json:"name" binding:"max=100"`
	State     uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`
	PassWord  string `form:"pass_word" json:"passWord" binding:"required,min=6,max=10"`
	CreatedBy string `form:"created_by" json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// 定义方法
func (u *User) TableName() string {
	return "t_user"
}

// 创建
func (u *User) Create(db *gorm.DB) error {
	return db.Debug().Create(u).Error
}

func (u *User)Update(db *gorm.DB) error {
	return db.Debug().UpdateColumns(u).Error
}

//delete
func (m *User) DelteById(db *gorm.DB) error {
	return db.Debug().Delete(m).Error
}

//get
func (m *User) GetById(db *gorm.DB) interface{} {
	db.Debug().Find(m, "id=?", m.ID)
	return m
}

//list
func (m *User) ListUserByPage(db *gorm.DB,pageSize int,pageNo int) (interface{},error) {
	if pageNo<=0{
		pageNo=1
	}
	if(pageSize<=0){
		pageSize=10
	}
	//你的空的结果数组
	resultSet := make([]*User,0,30)
	//准备一个写好查询条件的gorm.DB，注意要执行过Module()
	handler := db.Model(&User{}).Where(&User{})
	//进行查询
	return pageUtil.PageQuery(pageNo,pageSize,handler,&resultSet)
}

