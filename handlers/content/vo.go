package content

import "github.com/lwabish/typecho-api/models"

type VO struct {
	Content *models.TypechoContent `json:"content"`
	Meta    struct {
		Tags       []models.Tag      `json:"tags"`
		Categories []models.Category `json:"categories"`
	} `json:"meta"`
}

func (v *VO) SetText(c string) *VO {
	v.Content.Text = c
	return v
}

func (v *VO) SetTitle(t string) *VO {
	v.Content.Title = t
	return v
}

func (v *VO) SetCategories(cs []models.Category) *VO {
	v.Meta.Categories = cs
	return v
}

func (v *VO) SetTags(tags []models.Tag) *VO {
	v.Meta.Tags = tags
	return v
}

func (v *VO) SetSlug(s string) *VO {
	v.Content.Slug = s
	return v
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
