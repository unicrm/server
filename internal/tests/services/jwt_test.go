package services

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	defer os.Exit(m.Run())
}

func TestJWTBlackList(t *testing.T) {
	// TODO
	// 黑名单测试
	t.Run("黑名单测试", func(t *testing.T) {
		// 黑名单测试
		t.Log("黑名单测试")
	})
}
