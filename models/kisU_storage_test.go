/**
* @Author : henry
* @Data: 2020-08-17 19:29
* @Note:
**/

package models

import (
	"fmt"
	"github.com/vouchersAPI/app"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	addVoucherTest()
}

func addVoucherTest() {
	config := app.GetDBInfo()
	InitMssql(config)

	kisUVoucher := KisUVoucher{
		FCreditTotal:     100,
		FDate:            time.Date(2014, time.January, 1, 1, 1, 1, 1, time.UTC),
		FDebitTotal:      100,
		FPreparer:        "Administrator",
		KisUVoucherEntry: make([]*KisUVoucherEntry, 2),
	}

	kisUVoucher.KisUVoucherEntry[0] = &KisUVoucherEntry{
		FAmountFor:    100,
		FDC:           1,
		FAccountName:  "1002.01",
		FAccountName2: "1001",
		FCurrency:     "RMB",
	}

	kisUVoucher.KisUVoucherEntry[1] = &KisUVoucherEntry{
		FAmountFor:    100,
		FDC:           0,
		FAccountName:  "1001",
		FAccountName2: "1002.01",
		FCurrency:     "RMB",
	}

	err := kisUVoucher.AddVoucher()
	if err != nil {
		fmt.Printf("add voucher err :%s", err)
	}
}
