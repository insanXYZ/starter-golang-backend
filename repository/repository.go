package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Where(db *gorm.DB,cond *T, dst *T) error {
	return db.Where(cond).Take(dst).Error
}

func (r *Repository[T]) CountWhere(db *gorm.DB, where any, args ...any) int {
	var count int64
	db.Model(new(T)).Where(where, args).Count(&count)
	return int(count)
}
