package models

import (
	"fmt"
	"gorm.io/gorm"
)

// AddRelationship 添加分类/标签和文章的关联关系
func AddRelationship(db *gorm.DB, cid int32, tm *TypechoMeta) error {
	if r := db.Create(&TypechoRelationship{Cid: cid, Mid: tm.Mid}); r.Error != nil || r.RowsAffected != 1 {
		return fmt.Errorf("failed to create relationship: %w", r.Error)
	}
	// 更新meta表中的meta数量字段
	tm.AddCount(db)
	return nil
}
