package initialize

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/tasks/tables"
	"go.uber.org/zap"
)

func RegisterTimer() {
	go func() {

		// 设置定时任务支持秒级别
		var option []cron.Option
		var err error
		option = append(option, cron.WithSeconds())

		// 心跳检测，每10秒执行一次
		_, err = globals.UNICRM_TIMER.AddTaskByFunc("heartbeat", "*/10 * * * * *", func() {
			fmt.Println("心跳检测", time.Now().Format("2006-01-02 15:04:05"))
		}, "心跳检测", option...)
		if err != nil {
			zap.L().Panic("心跳检测定时任务注册失败", zap.Error(err))
		}

		// 清理数据库定时任务，每天执行一次
		_, err = globals.UNICRM_TIMER.AddTaskByFunc("ClearDB", "@daily", func() {
			if err := tables.ClearTables(globals.UNICRM_DB); err != nil {
				zap.L().Panic("清理数据库定时任务注册失败", zap.Error(err))
			}
		}, "清理数据库定时任务", option...)
		if err != nil {
			zap.L().Panic("清理数据库定时任务注册失败", zap.Error(err))
		}

	}()
}
