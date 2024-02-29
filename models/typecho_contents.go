package models

import "time"

// PostDefault 文章默认模板
// 对于数据库无默认值，且默认值和go的零值不匹配的字段，需要重设默认值
// 仅用于新blog
func (c *TypechoContent) PostDefault() {
	c.AuthorID = 1
	c.AllowComment = "1"
	c.AllowPing = "1"
	c.AllowFeed = "1"
	c.Created = int32(time.Now().Unix())
	c.Modified = int32(time.Now().Unix())
	// todo: wordCount created modified
	// todo: markdown basic
	// markdown wrapper
	// todo: markdown advanced
	// TOC标记
	// AI摘要
	// 原文链接
}

func (c *TypechoContent) UpdatePost(new *TypechoContent) {
	c.Title = new.Title
	c.Slug = new.Slug
	c.Text = new.Text
	c.Modified = int32(time.Now().Unix())
}

func (c *TypechoContent) IsNewPost() bool {
	return c.Cid == 0
}
