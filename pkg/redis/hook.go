package redis

import (
	"context"
	"net"

	"github.com/redis/go-redis/v9"
)

// 添加的钩子，必须实现Hook接口
// DialHook: 当创建网络连接时调用的hook
// ProcessHook: 执行命令时调用的hook
// ProcessPipelineHook: 执行管道命令时调用的hook
type Hook interface {
	DialHook(next redis.DialHook) redis.DialHook
	ProcessHook(next redis.ProcessHook) redis.ProcessHook
	ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook
}

type RedisHook struct{}

func (RedisHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

func (RedisHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		next(ctx, cmd)
		return nil
	}
}
func (RedisHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}
