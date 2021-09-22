package model

import "gorm.io/gorm"

// 入参校验
type FileModle struct {
	Id string   `form:"id" json:"id" binding:"max=10" gorm:"primary_key"`
	Name      string `form:"name" json:"name" binding:"required,max=100"`
	FileType     uint8  `form:"file_type,default=1" json:"file_type" binding:"oneof=0 1 2 3"`
	FileSize  string `form:"file_size" json:"file_size" binding:"required,max=10"`
	CreatedTime string `form:"created_time" json:"created_time"`
	CreatedBy string `form:"created_by" json:"created_by"`
	UpdatedTime string `form:"updated_time" json:"updated_time"`
	UpdatedBy string `form:"updated_by" json:"updated_by"`
}
// 定义方法
func (f *FileModle) TableName() string {
	return "t_file"
}

func (f *FileModle) Create(db *gorm.DB) error {
	return db.Debug().Create(f).Error
}

func (f *FileModle) Update(db *gorm.DB) error {
	return db.Debug().UpdateColumns(f).Error
}

func (f *FileModle) Delete(db *gorm.DB) error {
	return db.Debug().Delete(f).Error
}

func (f *FileModle) Get(db *gorm.DB) interface{} {
	db.Debug().Find(&f, "id=?", f.Id)
	return f
}