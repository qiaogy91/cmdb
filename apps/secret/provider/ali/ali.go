package ali

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/qiaogy91/cmdb/apps/secret"
)

func (i *Impl) SyncResource(req *secret.SyncResourceRequest, server secret.Rpc_SyncResourceServer) error {

	sql := i.db.WithContext(context.Background()).Model(&secret.Secret{})

	// 找出来这把秘钥

	sql.Where("id =? AND ", req.Id, req.Vendor)

	// 初始化客户端

	// 进行同步
	return nil
}

func NewClient(AccessKeyId, AccessKeySecret, Endpoint string) (*ecs20140526.Client, error) {
	conf := &openapi.Config{
		AccessKeyId:     &AccessKeyId,
		AccessKeySecret: &AccessKeySecret,
		EndpointType:    &Endpoint,
	}
	ins, err := ecs20140526.NewClient(conf)

	if err != nil {
		return nil, err
	}
	return ins, nil
}
