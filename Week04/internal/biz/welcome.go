package biz

import (
	"Week04/internal/data"
	"Week04/internal/model"
)

// ZRepo inter
type ZRepo interface {
	GetZ(id int64) (*model.ZItem, error)
}

// ZBiz biz
type ZBiz struct {
	repo ZRepo
}

// NewZBiz new
func NewZBiz(repo *data.ZRepository) *ZBiz {
	return &ZBiz{repo: repo}
}

// GetZ get zitem by id
func (b *ZBiz) GetZ(id int64) (*model.ZItem, error) {
	return b.repo.GetZ(id)
}
