/**
* @Author : henry
* @Data: 2020-08-13 13:01
* @Note: 初始化config.json
**/

package app

import (
	"io/ioutil"
	"os"
)

var dbInfo string

func GetDBInfo() string {
	return dbInfo
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func initConfig() {
	configJson := "config.json"

	if Exist(configJson) {
		info, err := ioutil.ReadFile(configJson)
		if err != nil {
			Logger.Fatalln("reading config.json failed: ", err)
		}
		dbInfo = string(info)
	} else {
		Logger.Fatalln("No config.json")
	}
}
