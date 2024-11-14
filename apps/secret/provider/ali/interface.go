package ali

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
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
