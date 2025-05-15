package options

import (
	"github.com/ketitongxue/fastgo/internal/apiserver"
	genericoptions "github.com/ketitongxue/fastgo/pkg/options"
)

// ConfigFrom defines filepath for configfile.
type ConfigFrom struct {
	Filepath string `json:"filepath,omitempty" mapstructure:"filepath"`
}

func NewConfigFrom() *ConfigFrom {
	return &ConfigFrom{
		Filepath: "default",
	}
}

type ServerOptions struct {
	ConfigFrom   *ConfigFrom                  `json:"config" mapstructure:"config"`
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		ConfigFrom:   NewConfigFrom(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
	}
}

// Validate 校验 ServerOptions 中的选项是否合法.
// 提示：Validate 方法中的具体校验逻辑可以由 Claude、DeepSeek、GPT 等 LLM 自动生成。
func (o *ServerOptions) Validate() error {
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

	return nil
}

// Config 基于 ServerOptions 构建 apiserver.Config.
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
	}, nil
}
