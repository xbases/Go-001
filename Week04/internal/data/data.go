package data

import (
	"Week04/internal/model"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"

	"github.com/go-redis/redis"
)

// ZRepository Repo
type ZRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

// NewZRepository new repo
func NewZRepository(db *gorm.DB, cache *redis.Client) *ZRepository {
	return &ZRepository{
		db:    db,
		cache: cache,
	}
}

// GetZ get
func (d *ZRepository) GetZ(id int64) (*model.ZItem, error) {
	var z model.ZItem
	if err := d.db.Model(&z).Where("id = ?", id).First(&z).Error; err != nil {
		return nil, errors.Wrap(err, "get z error")
	}
	return &z, nil
}
