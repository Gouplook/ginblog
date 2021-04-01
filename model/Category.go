/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  Category.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:47
 */
package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type: varchar(20);not null" json:"name"`

}
