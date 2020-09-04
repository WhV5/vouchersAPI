/**
* @Author : henry
* @Data: 2020-08-12 17:28
* @Note:
**/

package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/vouchersAPI/app"
	"github.com/vouchersAPI/ctrls/routers"
	"github.com/vouchersAPI/models"
	"net/http"
)

func main() {

	config := app.GetDBInfo()
	models.InitMssql(config)

	router := routers.InitRouter()

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", "127.0.0.1", "8899"),
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		app.Logger.Fatalln("listen error: ", err)
	}
}
