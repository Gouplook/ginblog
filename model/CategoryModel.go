/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  CategoryModel.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:47
 */
package model

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type: varchar(20);not null" json:"name"`

}
