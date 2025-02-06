package logging

import (
	"errors"
	"fmt"
	"os"
)

// PathExist 判断文件或目录是否存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsDir 判断是否为目录
func IsDir(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

// 支持批量创建目录
func CreateDir(paths ...string) error {
	for _, path := range paths {
		if ok, _ := PathExists(path); !ok {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				fmt.Printf("创建目录失败: %s\n", path)
				return err
			}
		}
	}
	return nil
}
