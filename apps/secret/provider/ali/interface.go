package ali

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/qiaogy91/cmdb/apps/resource"
	"github.com/qiaogy91/cmdb/apps/secret"
)

func NewClient(AccessKeyId, AccessKeySecret, Endpoint, RegionId string) (*ecs20140526.Client, error) {
	conf := &openapi.Config{
		AccessKeyId:     &AccessKeyId,
		AccessKeySecret: &AccessKeySecret,
		EndpointType:    &Endpoint,
		RegionId:        &RegionId,
	}
	ins, err := ecs20140526.NewClient(conf)

	if err != nil {
		return nil, err
	}
	return ins, nil
}

func NewDescribeInstancesRequest(req *secret.SyncResourceRequest) *ecs20140526.DescribeInstancesRequest {
	return &ecs20140526.DescribeInstancesRequest{
		RegionId: &req.RegionId,
	}
}

func NewResourceSpec(req *ecs20140526.DescribeInstancesResponseBodyInstancesInstance) *resource.Spec {
	return &resource.Spec{
		Vendor:         resource.Vendor_VENDOR_ALI,
		Type:           resource.Type_TYPE_ECS,
		ResourceId:     tea.StringValue(req.InstanceId),
		Name:           tea.StringValue(req.HostName),
		Region:         tea.StringValue(req.RegionId),
		Zone:           tea.StringValue(req.ZoneId),
		Status:         tea.StringValue(req.Status),
		PrivateAddress: tea.StringSliceValue(req.VpcAttributes.PrivateIpAddress.IpAddress),
		PublicAddress:  tea.StringSliceValue(req.PublicIpAddress.IpAddress),
		Cpu:            tea.Int32Value(req.Cpu),
		Memory:         int64(tea.Int32Value(req.Memory)),
		Storage:        0,
		Extra:          nil,
	}

}
