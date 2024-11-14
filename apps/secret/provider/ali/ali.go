package ali

import (
	"context"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/cmdb/apps/secret"
)

func (i *Impl) SyncResource(req *secret.SyncResourceRequest, stream secret.Rpc_SyncResourceServer) error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	// 找出秘钥
	sql := i.db.WithContext(context.Background()).Model(&secret.Secret{})
	sec := &secret.Secret{}
	if err := sql.Where("id =? AND vendor = ?", req.SecretId, req.Vendor).First(sec).Error; err != nil {
		return err
	}

	// 构建客户端
	cli, err := NewClient(sec.Spec.AccessKey, sec.Spec.AccessSecret, req.AccessEndpoint, req.RegionId)
	if err != nil {
		return err
	}
	// 发起一次调用
	in := NewDescribeInstancesRequest(req)
	ins, err := cli.DescribeInstancesWithOptions(in, &util.RuntimeOptions{})
	if err != nil {
		return err
	}

	// 进行同步
	for _, ecs := range ins.Body.Instances.Instance {
		// gorm 入库
		rsp, err := i.resource.CreateResource(context.Background(), NewResourceSpec(ecs))
		if err != nil {
			return err
		}

		// stream 回调
		if err := stream.Send(rsp); err != nil {
			return err
		}
	}
	return nil
}
