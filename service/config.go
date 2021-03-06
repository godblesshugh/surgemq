package service

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"os"
	"path/filepath"
	"surgemq/utils"
)

var (
	// AppConfig is the default config for Application
	AppConfig config.Configer
	// AppPath is the absolute path to the Application
	AppPath string
	// appConfigPath is the path to the config files
	appConfigPath string
)

func init() {
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}
	configExist := false
	_appConfigPath := ""
	_appConfigPath = filepath.Join(workPath, "config", "app.conf")
	if utils.FileExists(_appConfigPath) {
		appConfigPath = _appConfigPath
		configExist = true
	}
	_appConfigPath = filepath.Join(workPath, "app.conf")
	if utils.FileExists(_appConfigPath) {
		appConfigPath = _appConfigPath
		configExist = true
	}
	_appConfigPath = filepath.Join(AppPath, "config", "app.conf")
	if utils.FileExists(_appConfigPath) {
		appConfigPath = _appConfigPath
		configExist = true
	}
	_appConfigPath = filepath.Join(AppPath, "app.conf")
	if utils.FileExists(_appConfigPath) {
		appConfigPath = _appConfigPath
		configExist = true
	}
	if configExist {
		// 暂时只支持 ini 格式的配置文件
		AppConfig, err = config.NewConfig("ini", appConfigPath)
	} else {
		AppConfig = config.NewFakeConfig()
	}
	if err != nil {
		panic(err)
		return
	}
}
