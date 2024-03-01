package models

import (
	"fmt"
	"gorm.io/gorm"
)

// AddRelationship 添加分类/标签和文章的关联关系
// todo: 用发布订阅模式，更新meta表中的meta数量字段
func AddRelationship(db *gorm.DB, cid, mid int32) error {
	if r := db.Create(&TypechoRelationship{Cid: cid, Mid: mid}); r.Error != nil || r.RowsAffected != 1 {
		return fmt.Errorf("failed to create relationship: %w", r.Error)
	}
	return nil
}
