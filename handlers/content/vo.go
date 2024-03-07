package content

import "github.com/lwabish/typecho-api/models"

type VO struct {
	Content *models.TypechoContent `json:"content"`
	Meta    struct {
		Tags       []models.Tag      `json:"tags"`
		Categories []models.Category `json:"categories"`
	} `json:"meta"`
}

func NewVo() *VO {
	return &VO{
		Content: &models.TypechoContent{},
	}
}

func (v *VO) SetCid(cid int) *VO {
	v.Content.Cid = int32(cid)
	return v
}

func (v *VO) aggregateMeta() []models.Meta {
	var metas []models.Meta
	for _, tag := range v.Meta.Tags {
		metas = append(metas, tag)
	}
	for _, category := range v.Meta.Categories {
		metas = append(metas, category)
	}
	return metas
}
