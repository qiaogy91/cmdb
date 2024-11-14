package impl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/cmdb/apps/resource"
	"gorm.io/gorm/clause"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&resource.Resource{})
}

func (i *Impl) DescResource(ctx context.Context, req *resource.DescResourceRequest) (*resource.Resource, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	sql := i.db.WithContext(ctx).Model(&resource.Resource{})
	ins := &resource.Resource{}

	if err := sql.Where("resource_id = ?", req.ResourceId).First(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}
func (i *Impl) CreateResource(ctx context.Context, req *resource.Spec) (*resource.Resource, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	// 逻辑描述：
	// 1. 更新：如果本地存在，则更新
	// 2. 创建：如果本地不存在，则创建
	// 2. 删除：如果云商不存在的资源，则需要删除，但是从这里无法判断哪些资源是要删除的，因此需要开一个后台定时任务，拿着本地一个个资源ID 去云商查询，查询不到则删除
	sql := i.db.WithContext(ctx).Model(&resource.Resource{})

	exp := clause.OnConflict{
		Columns:   []clause.Column{{Name: "resource_id"}}, // 检查哪些字段是否重复，前提该字段定义了唯一约束
		DoNothing: false,                                  // 当重复键冲突时不执行任何操作
		UpdateAll: true,                                   // 更新所有字段
	}

	ins := &resource.Resource{Spec: req}
	if err := sql.Clauses(exp).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) QueryResource(ctx context.Context, req *resource.QueryResourceRequest) (*resource.ResourceSet, error) {
	//TODO implement me
	panic("implement me")
}
