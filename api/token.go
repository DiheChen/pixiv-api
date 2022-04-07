package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DiheChen/PixivAPI/auth"
)

// PathExists 判断给定 path 是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || errors.Is(err, os.ErrExist)
}

func GetRefreshToken() string {
	if !pathExists("refresh_token") {
		fmt.Println("未找到 refresh_token 文件，请先登录")
		refreshToken, err := auth.Login()
		if err == nil {
			_ = os.WriteFile("refresh_token", []byte(refreshToken), 0644)
			fmt.Println("已保存 refresh_token 到文件, 请重新启动程序。")
		}
		fmt.Println("登录失败:", err)
		os.Exit(0)
	}
	file, err := os.Open("refresh_token")
	if err != nil {
		fmt.Println("打开 refresh_token 文件失败:", err)
		os.Exit(0)
	}
	defer func() { _ = file.Close() }()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(all)
}
