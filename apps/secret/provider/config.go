package provider

type Conf struct {
	Ali     *Ali     `json:"ali" yaml:"ali"`
	Tencent *Tencent `json:"tencent" yaml:"tencent"`
}

type Ali struct{}

type Tencent struct{}
