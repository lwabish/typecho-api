package models

import (
	"gorm.io/gorm"
)

const (
	metaTypeTag      = "tag"
	metaTypeCategory = "category"
)

type Meta interface {
	Type() string
	String() string
}
type Tag string

func (t Tag) Type() string {
	return metaTypeTag
}
func (t Tag) String() string {
	return string(t)
}

type Category string

func (c Category) Type() string {
	return metaTypeCategory
}

func (c Category) String() string {
	return string(c)
}

func GetMeta(db *gorm.DB, meta Meta) (*TypechoMeta, error) {
	tm := &TypechoMeta{
		Name: meta.String(),
		Type: meta.Type(),
	}
	r := db.FirstOrCreate(tm, tm)
	if r.Error != nil {
		return nil, r.Error
	}
	if r.RowsAffected == 1 {
		// 新增的需要设置默认值
		tm.setDefaultValue(db)
		return tm, db.Save(tm).Error
	}
	return tm, nil
}

func (m *TypechoMeta) setDefaultValue(db *gorm.DB) {
	m.Slug = m.Name
	if m.Type == metaTypeCategory {
		tmp := &TypechoMeta{}
		db.Find(tmp, &TypechoMeta{Type: metaTypeCategory}).Order("`order` DESC").First(tmp)
		m.Order = tmp.Order + 1
	}
}

func (m *TypechoMeta) AddCount(db *gorm.DB) {
	m.Count += 1
	db.Save(m)
}
