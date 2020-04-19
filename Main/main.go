package main

import (
	"errors"
	log "github.com/jeanphorn/log4go"
)

var err  = errors.New("category Test error test ...")
func main() {
	log.LoadConfiguration("../configLog/example.json")

	// 其中category和存储路径是一一对应的，根据配置文件可知，会输入到当前路径中的test.log文件中
	log.LOGGER("Test").Info("category Test info test ...")
	log.LOGGER("Test").Info("category Test info test message: %s", "new test msg")
	log.LOGGER("Test").Debug("category Test debug test ...")
	log.LOGGER("Test").Error("error = %s", err)

	log.LOGGER("Other").Debug("category Other debug test ...")
	log.LOGGER("TestSocket").Debug("category TestSocket debug test ...")

	//原始的log4go测试
	log.Info("nomal info test ...")
	log.Debug("nomal debug test ...")

	log.Close()
}