package content

import "github.com/lwabish/typecho-api/models"

type vo struct {
	Content *models.TypechoContent `json:"content"`
	Meta    struct {
		Tags       []models.Tag      `json:"tags"`
		Categories []models.Category `json:"categories"`
	} `json:"meta"`
}

func (v vo) aggregateMeta() []models.Meta {
	var metas []models.Meta
	for _, tag := range v.Meta.Tags {
		metas = append(metas, tag)
	}
	for _, category := range v.Meta.Categories {
		metas = append(metas, category)
	}
	return metas
}
