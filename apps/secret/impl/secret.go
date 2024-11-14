package impl

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/cmdb/apps/secret/provider"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&secret.Secret{})
}

func (i *Impl) CreateSecret(ctx context.Context, req *secret.Spec) (*secret.Secret, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}
	ins := &secret.Secret{Spec: req}
	if err := i.db.WithContext(ctx).Model(&secret.Secret{}).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) DeleteSecret(ctx context.Context, req *secret.DeleteSecretRequest) (*secret.Secret, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins, err := i.DescSecret(ctx, &secret.DescSecretRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Model(&secret.Secret{}).Where("id = ?", req.Id).Delete(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) QuerySecret(ctx context.Context, req *secret.QuerySecretRequest) (*secret.SecretSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	sql := i.db.WithContext(ctx).Model(&secret.Secret{})
	ins := &secret.SecretSet{}

	off := int((req.PageNum - 1) * req.PageSize)
	lim := int(req.PageSize)

	switch req.QueryType {
	case secret.QueryType_QUERY_TYPE_DESC:
		sql = sql.Where("desc like ?", "%"+req.Keyword+"%")

	}

	if err := sql.Count(&ins.Total).Offset(off).Limit(lim).Find(&ins.Items).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) DescSecret(ctx context.Context, req *secret.DescSecretRequest) (*secret.Secret, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}
	ins := &secret.Secret{}
	if err := i.db.WithContext(ctx).Model(&secret.Secret{}).Where("id = ?", req.Id).First(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) SyncResource(req *secret.SyncResourceRequest, stream secret.Rpc_SyncResourceServer) error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	ins := provider.GetProvider(req.Vendor)
	if ins == nil {
		return fmt.Errorf("no such provider, because provider not added")
	}

	return ins.SyncResource(req, stream)
}
