package initialize

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/unicrm/server/internal/config"
	"github.com/unicrm/server/internal/globals"
)

// LoadConfig 加载配置文件，并解析到全局变量中
// 优先级: 函数参数 > 命令行 > 环境变量 > 默认值（使用gin模式加载配置文件）
func LoadConfig(path ...string) *viper.Viper {
	// 这里可以添加加载配置的逻辑，例如从文件、环境变量等读取配置信息

	var configPath string

	if len(path) == 0 {
		flag.StringVar(&configPath, "c", "", "配置文件路径")
		flag.Parse()
		if configPath == "" {
			if configEnv := os.Getenv(config.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					configPath = config.ConfigDebugFile
				case gin.ReleaseMode:
					configPath = config.ConfigReleaseFile
				case gin.TestMode:
					configPath = config.ConfigTestFile
				default:
					configPath = config.ConfigDefaultFile
				}
				fmt.Printf("使用gin模式加载配置文件: %s\n", configPath)
			} else {
				configPath = configEnv
				fmt.Printf("使用环境变量加载配置文件: %s\n", configPath)
			}
		} else {
			fmt.Printf("使用命令行参数加载配置文件: %s\n", configPath)
		}
	} else {
		configPath = path[0]
		fmt.Printf("使用函数参数加载配置文件: %s\n", configPath)
	}

	// 初始化配置文件
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}
	v.WatchConfig()

	// 监听配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已更改:", e.Name)
		if err := v.Unmarshal(&globals.UNICRM_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	// 解析配置文件，并写入全局变量
	if err := v.Unmarshal(&globals.UNICRM_CONFIG); err != nil {
		panic(fmt.Errorf("解析配置文件失败: %s", err))
	}

	// 写入环境变量
	if config.ConfigAutomaticEnv {
		v.AutomaticEnv()
	}

	return v
}
