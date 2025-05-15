package options

import (
	"fmt"
	"net"
	"strconv"
	
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
	Addr         string                       `json:"addr" mapstructure:"addr"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		ConfigFrom:   NewConfigFrom(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
        Addr:         "0.0.0.0:6666",
	}
}

// Validate 校验 ServerOptions 中的选项是否合法.
// 提示：Validate 方法中的具体校验逻辑可以由 Claude、DeepSeek、GPT 等 LLM 自动生成。
func (o *ServerOptions) Validate() error {
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

    // 验证服务器地址
    if o.Addr == "" {
        return fmt.Errorf("server address cannot be empty")
    }

    // 检查地址格式是否为host:port
    _, portStr, err := net.SplitHostPort(o.Addr)
    if err != nil {
        return fmt.Errorf("invalid server address format '%s': %w", o.Addr, err)
    }

    // 验证端口是否为数字且在有效范围内
    port, err := strconv.Atoi(portStr)
    if err != nil || port < 1 || port > 65535 {
        return fmt.Errorf("invalid server port: %s", portStr)
    }

	return nil
}

// Config 基于 ServerOptions 构建 apiserver.Config.
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
        Addr:         o.Addr,
	}, nil
}
