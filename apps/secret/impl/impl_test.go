package impl_test

import (
	"context"
	_ "github.com/qiaogy91/cmdb/apps"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"testing"
)

var (
	ctx = context.Background()
	c   secret.Service
)

func init() {
	cfg := "/Users/qiaogy/GolandProjects/projects/github/CloudManager/cmdb/etc/application.yaml"
	if err := ioc.ConfigIocObject(cfg); err != nil {
		panic(err)
	}
	c = secret.GetSvc()
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_DeleteSecret(t *testing.T) {
	req := &secret.DeleteSecretRequest{Id: 1}
	ins, err := c.DeleteSecret(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}
func TestImpl_CreateSecret(t *testing.T) {
	req := &secret.Spec{
		Name:         "key01",
		Desc:         "阿里云杭州1区",
		Vendor:       secret.Vendor_VENDOR_ALI,
		AccessKey:    "LTAI5t7daLrncQLEj1e5yyZt",
		AccessSecret: "kWbqurqSqWI2Mz7DKmtZOxhhVjWBAs",
	}
	ins, err := c.CreateSecret(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}

func TestImpl_DescSecret(t *testing.T) {
	req := &secret.DescSecretRequest{Id: 3}
	ins, err := c.DescSecret(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}

func TestImpl_SyncResource(t *testing.T) {
	req := &secret.SyncResourceRequest{
		Vendor:         secret.Vendor_VENDOR_ALI,
		SecretId:       3,
		AccessEndpoint: "ecs.cn-hangzhou.aliyuncs.com",
		RegionId:       "cn-hangzhou",
	}

	s := NewStream()
	if err := c.SyncResource(req, s); err != nil {
		t.Fatal(err)
	}
}
