/**
* @Author : henry
* @Data: 2020-08-13 13:22
* @Note:
**/

package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/vouchersAPI/app"
)

var MsDB *gorm.DB
var err error

type VoucherDB struct {
	Type     string `json:"type"`
	Ip       string `json:"ip"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Encrypt  string `json:"encrypt"`
}

func InitMssql(config string) {
	var voucherDB VoucherDB

	if err = json.Unmarshal([]byte(config), &voucherDB); err != nil {
		app.Logger.Fatalln("unmarshal mssqlDb failed : ", err)
	}

	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=%s",
		voucherDB.Ip, voucherDB.Port, voucherDB.Database, voucherDB.User, voucherDB.Pwd, voucherDB.Encrypt)

	MsDB, err = gorm.Open("mssql", connString)
	if err != nil {
		app.Logger.Fatalln("open mssql failed: ", err)
	}

	err = MsDB.DB().Ping()
	if err != nil {
		app.Logger.Fatalln("connect to mssql failed: ", err)
	}
}
