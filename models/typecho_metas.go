package models

import "gorm.io/gorm"

type Meta interface {
	Type() string
	String() string
}
type Tag string

func (t Tag) Type() string {
	return "tag"
}
func (t Tag) String() string {
	return string(t)
}

type Category string

func (c Category) Type() string {
	return "category"
}

func (c Category) String() string {
	return string(c)
}

func GetMeta(db *gorm.DB, meta Meta) (*TypechoMeta, error) {
	tm := &TypechoMeta{
		Name: meta.String(),
		Type: meta.Type(),
	}
	if r := db.FirstOrCreate(tm); r.Error != nil {
		return nil, r.Error
	}
	return tm, nil
}
