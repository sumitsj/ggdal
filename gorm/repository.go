package gorm

import (
	"context"
	"gorm.io/gorm"
)

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *repository[T] {
	return &repository[T]{
		db: db,
	}
}

func (r *repository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository[T]) CreateAll(ctx context.Context, entity *[]T) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository[T]) GetFirst(ctx context.Context, params *T) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Where(&params).FirstOrInit(&entity).Error
	return &entity, err
}

func (r *repository[T]) GetAll(ctx context.Context, params *T) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Where(&params).Find(&entities).Error
	return &entities, err
}

func (r *repository[T]) GetByPage(ctx context.Context, pageNo int, pageSize int) (*[]T, error) {
	skip := (pageNo - 1) * pageSize
	var entities []T
	err := r.db.WithContext(ctx).Offset(skip).Limit(pageSize).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *repository[T]) Delete(ctx context.Context, params *T) error {
	return r.db.WithContext(ctx).Where(&params).Delete(params).Error
}

func (r *repository[T]) Count(ctx context.Context) (int64, error) {
	var entity T
	var count int64
	tx := r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count, tx.Error
}

func (r *repository[T]) CountBy(ctx context.Context, params *T) (int64, error) {
	var entity T
	var count int64
	tx := r.db.WithContext(ctx).Model(&entity).Where(&params).Count(&count)
	return count, tx.Error
}
