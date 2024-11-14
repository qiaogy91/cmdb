package ali

import (
	"context"
	"fmt"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-playground/validator/v10"
	"github.com/qiaogy91/cmdb/apps/secret"
	"strings"
)

func (i *Impl) SyncResource(req *secret.SyncResourceRequest, stream secret.Rpc_SyncResourceServer) error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	// 找出秘钥
	sql := i.db.WithContext(context.Background()).Model(&secret.Secret{})
	sec := &secret.Secret{}
	if err := sql.Where("id =? AND vendor = ?", req.Id, req.Vendor).First(sec).Error; err != nil {
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

		privateIps := tea.StringSliceValue(ecs.VpcAttributes.PrivateIpAddress.IpAddress)
		publicIps := tea.StringSliceValue(ecs.PublicIpAddress.IpAddress)
		dic := map[string]any{
			"InstanceId":          tea.StringValue(ecs.InstanceId),
			"StartTime":           tea.StringValue(ecs.StartTime),
			"OSName":              tea.StringValue(ecs.OSName),
			"HostName":            tea.StringValue(ecs.HostName),
			"Status":              tea.StringValue(ecs.Status),
			"Memory":              tea.Int32Value(ecs.Memory),
			"Cpu":                 tea.Int32Value(ecs.Cpu),
			"InstanceNetworkType": tea.StringValue(ecs.InstanceNetworkType),
			"PrivateIpAddress":    strings.Join(privateIps, ","),
			"PublicIpAddress":     strings.Join(publicIps, ","),
			"ZoneId":              tea.StringValue(ecs.ZoneId),
			"RegionId":            tea.StringValue(ecs.RegionId),
		}

		rsp := &secret.SyncResourceResponse{
			Status:  false,
			Message: fmt.Sprintf("%+v", dic),
		}

		if err := stream.Send(rsp); err != nil {
			return err
		}
	}
	return nil
}
