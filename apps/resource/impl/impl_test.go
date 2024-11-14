package impl_test

import (
	"context"
	_ "github.com/qiaogy91/cmdb/apps"
	"github.com/qiaogy91/cmdb/apps/resource"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"testing"
)

var (
	ctx = context.Background()
	c   resource.Service
)

func init() {
	cfg := "/Users/qiaogy/GolandProjects/projects/github/CloudManager/cmdb/etc/application.yaml"
	if err := ioc.ConfigIocObject(cfg); err != nil {
		panic(err)
	}
	c = resource.GetSvc()
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateResource(t *testing.T) {
	req := &resource.Spec{
		Vendor:         resource.Vendor_VENDOR_ALI,
		Type:           resource.Type_TYPE_ECS,
		ResourceId:     "0001",
		Name:           "tewst",
		Region:         "aa",
		Zone:           "aa",
		Status:         "aa",
		PrivateAddress: []string{"10.0.0.1", "10.0.0.2"},
		PublicAddress:  []string{"1.1.1.1", "2.2.2.2"},
		Cpu:            2,
		Memory:         4,
		Storage:        50,
		Extra:          nil,
	}

	ins, err := c.CreateResource(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}
