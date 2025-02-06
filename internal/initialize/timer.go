package initialize

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/unicrm/server/internal/globals"
)

func RegisterTimer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 心跳检测，每10秒执行一次
		_, err := globals.UNICRM_TIMER.AddTaskByFunc("heartbeat", "*/10 * * * * *", func() {
			fmt.Println("心跳检测", time.Now().Format("2006-01-02 15:04:05"))
		}, "心跳检测", option...)
		if err != nil {
			fmt.Println("心跳检测定时任务注册失败", err)
		}
	}()
}
