/**
* @Author : henry
* @Data: 2020-08-13 21:15
* @Note: 金蝶旗舰版数据存储
**/

package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/vouchersAPI/app"
	"strconv"
	"strings"
	"time"
)

const intoVoucher = `INSERT INTO t_Voucher (FDate,FTransDate,FYear,FPeriod,FGroupID,FNumber,FReference,
										FExplanation,FAttachments,FEntryCount,FDebitTotal,FCreditTotal,FInternalInd,FChecked,
										FPosted,FPreparerID,FCheckerID,FPosterID,FCashierID,FHandler,FObjectName,
										FSerialNum,FTranType,FOwnerGroupID)
								VALUES (?,?,?,?,?,?,?, ?,?,?,?,?,?,?, ?,?,?,?,?,?,?, ?,?,?)`

var months = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func (kisUVoucher *KisUVoucher) AddVoucher() error {

	// 校验期间 借贷
	if err := kisUIsInPeriod(kisUVoucher); err != nil { // 1.是否在期间
		app.Logger.Errorln("voucher not in period:", err)
		return err
	}

	if ok := kisDcEqual(kisUVoucher.KisUVoucherEntry); !ok {
		app.Logger.Errorln("voucher dc not equal")
		return errors.New("voucher dc not equal")
	}

	// 校验必填项
	err := kisUDataExists(kisUVoucher)
	if err != nil {
		return err
	}

	// 数据转为数据库字段
	err = kisUDataPrepare(kisUVoucher)
	if err != nil {
		return err
	}

	// 新增数据到金蝶
	tx := MsDB.Begin()

	err = tx.Exec(intoVoucher, kisUVoucher.FDate, kisUVoucher.FTransDate, kisUVoucher.FYear, kisUVoucher.FPeriod, kisUVoucher.FGroupID, kisUVoucher.FNumber, kisUVoucher.FReference,
		kisUVoucher.FExplanation, kisUVoucher.FAttachments, kisUVoucher.FEntryCount, kisUVoucher.FDebitTotal, kisUVoucher.FCreditTotal, kisUVoucher.FInternalInd, kisUVoucher.FCheckerID,
		kisUVoucher.FPosted, kisUVoucher.FPreparerID, kisUVoucher.FCheckerID, kisUVoucher.FPosterID, kisUVoucher.FCashierID, kisUVoucher.FHandler, kisUVoucher.FObjectName,
		kisUVoucher.FSerialNum, kisUVoucher.FTranType, kisUVoucher.FOwnerGroupID).Error
	if err != nil {
		app.Logger.Errorln("Add voucher failed:", err)
		return err
	}

	voucherId := KisUVoucher{}
	err = tx.Where("FYear = ? and FPeriod = ? and FGroupID = ? and FNumber = ?",
		kisUVoucher.FYear, kisUVoucher.FPeriod, kisUVoucher.FGroupID, kisUVoucher.FNumber).Find(&voucherId).Error
	if err != nil {
		app.Logger.Errorln("select FVoucherID failed:", err)
		tx.Rollback()
		return err
	}
	kisUVoucher.FVoucherID = voucherId.FVoucherID

	//err = tx.Create(kisUVoucher).Error
	//if err != nil {
	//	app.Logger.Errorln("Add voucher failed:", err)
	//	app.Logger.Errorf("Add voucher failed: %#v", kisUVoucher)
	//	return err
	//}

	for i := 0; i < len(kisUVoucher.KisUVoucherEntry); i++ {
		kisUVoucher.KisUVoucherEntry[i].FVoucherID = kisUVoucher.FVoucherID
		err := tx.Create(kisUVoucher.KisUVoucherEntry[i]).Error
		if err != nil {
			app.Logger.Errorf("Add voucher entry %d failed : %s", i+1, err)
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

// 是否在期间内
func kisUIsInPeriod(voucher *KisUVoucher) error {
	date := voucher.FDate
	FYear := KisUSysPro{FValue: ""}
	FPeriod := KisUSysPro{FValue: ""}

	err := MsDB.Where("FCategory = ? and FKey = ?", "GL", "CurrentYear").Find(&FYear).Error
	if err != nil {
		app.Logger.Errorln("find kisU currentYear failed: ", err)
		return errors.New("find kisU currentYear failed")
	}

	err = MsDB.Where("FCategory = ? and FKey = ?", "GL", "CurrentPeriod").Find(&FPeriod).Error
	if err != nil {
		app.Logger.Errorln("find kisU currentPeriod failed: ", err)
		return errors.New("find kisU currentPeriod failed")
	}

	fYear, err := strconv.Atoi(FYear.FValue)
	if err != nil {
		return err
	}

	fPeriod, err := strconv.Atoi(FPeriod.FValue)
	if err != nil {
		return err
	}

	if date.Month().String() == months[fPeriod-1] && date.Year() == fYear {
		voucher.FYear = fYear
		voucher.FPeriod = fPeriod

		return nil
	} else {
		return errors.New("not in period")
	}

}

// 借贷是否相等
func kisDcEqual(voucherB []*KisUVoucherEntry) bool {
	var cAmount, dAmount float64

	if len(voucherB) < 2 {
		app.Logger.Errorln("KisUVoucherEntry not enough 2 entry")
		return false
	}

	for i := 0; i < len(voucherB); i++ {
		if voucherB[i].FDC == 0 { // 0 - 贷方
			cAmount += voucherB[i].FAmount
		} else if voucherB[i].FDC == 1 { // 1- 借方
			dAmount += voucherB[i].FAmount
		} else {
			app.Logger.Errorln("data error ,not in DC")
			return false
		}
	}

	return cAmount == dAmount
}

// 检查必填项
func kisUDataExists(voucher *KisUVoucher) error {

	// 凭证主表
	if voucher.FCreditTotal == 0 {
		return errors.New("credit total empty")
	}
	if voucher.FDebitTotal == 0 {
		return errors.New("debit total empty")
	}
	if voucher.FPreparer == "" {
		return errors.New("preparer empty")
	}

	// 凭证子表
	for i := 0; i < len(voucher.KisUVoucherEntry); i++ {
		if voucher.KisUVoucherEntry[i].FAmountFor == 0 {
			return errors.New("amountFor empty")
		}
		if voucher.KisUVoucherEntry[i].FDC != 0 && voucher.KisUVoucherEntry[i].FDC != 1 {
			return errors.New("dc empty")
		}
		if voucher.KisUVoucherEntry[i].FCurrency == "" {
			return errors.New("currency empty")
		}
		if voucher.KisUVoucherEntry[i].FAccountName == "" {
			return errors.New("account name empty")
		}
		if voucher.KisUVoucherEntry[i].FAccountName2 == "" {
			return errors.New("account name2 empty")
		}

		voucher.KisUVoucherEntry[i].FEntryID = i
	}

	return nil
}

// 补全数据
func kisUDataPrepare(voucher *KisUVoucher) error {

	var ok bool
	bwb := KisUSysPro{}

	users := make([]KisUUser, 0)
	userMap := make(map[string]int)

	accounts := make([]KisUAccount, 0)
	accountM := make(map[string]int)

	currency := make([]KisUCurrency, 0)
	currencyM := make(map[string]int)
	currencyM2 := make(map[string]float64)

	fields := strings.Fields(fmt.Sprintf("%s", voucher.FDate))
	dataStr := fmt.Sprintf("%s %s", fields[0], fields[1])
	tt, err := time.Parse("2006-01-02 15:04:05", dataStr)
	if err != nil {
		app.Logger.Errorln("failed to convert time: ", err)
		return err
	}
	voucher.FDate = tt
	voucher.FTransDate = tt

	id, err := kisUGetId(voucher.TableName())
	if err != nil {
		return err
	}
	voucher.FVoucherID = id

	voucher.FAttachments = 0
	voucher.FGroupID = 1
	// FNumber FExplanation
	voucher.FEntryCount = len(voucher.KisUVoucherEntry)
	voucher.FOwnerGroupID = 1
	// FSerialNum
	// FTranType FTransDate
	voucher.FFrameWorkID = -1

	number, err := kisUGetNumber(voucher.FYear, voucher.FPeriod, voucher.FGroupID)
	if err != nil {
		app.Logger.Errorln("failed to find FNumber: ", err)
		return err
	}
	voucher.FNumber = number

	err = MsDB.Find(&users).Where("FForbidden = 0").Error
	if err != nil {
		app.Logger.Errorln("failed to find users: ", err)
		return err
	}
	for i := 0; i < len(users); i++ {
		userMap[users[i].FName] = users[i].FUserID
	}

	// 科目
	err = MsDB.Find(&accounts).Where("FDelete = 0 and FDetail = 1").Error
	if err != nil {
		app.Logger.Errorln("failed to find accounts: ", err)
		return err
	}
	for j := 0; j < len(accounts); j++ {
		accountM[accounts[j].FNumber] = accounts[j].FAccountID
	}

	// 币别
	err = MsDB.Where("FDeleted = 0").Find(&currency).Error
	if err != nil {
		app.Logger.Errorln("failed to find currency: ", err)
		return err
	}
	for k := 0; k < len(currency); k++ {
		currencyM[currency[k].FNumber] = currency[k].FCurrencyID
		currencyM2[currency[k].FNumber] = currency[k].FExchangeRate
	}

	err = MsDB.Where("FCategory = ? and FKey = ?", "GL", "FBWB").Find(&bwb).Error
	if err != nil {
		app.Logger.Errorln("failed to find bwb: ", err)
		return err
	}

	if voucher.FPreparerID, ok = userMap[voucher.FPreparer]; !ok {
		app.Logger.Errorln("not find preparer in t_user")
		return errors.New("not find preparer in t_user")
	}

	if voucher.FCashierID, ok = userMap[voucher.FCashier]; !ok {
		voucher.FCashierID = voucher.FPreparerID
	}

	if voucher.FCheckerID, ok = userMap[voucher.FChecker]; !ok {
		voucher.FCheckerID = voucher.FPreparerID
	}

	if voucher.FPosterID, ok = userMap[voucher.FPoster]; !ok {
		voucher.FPosterID = voucher.FPreparerID
	}

	for m := 0; m < len(voucher.KisUVoucherEntry); m++ {

		// 补全科目
		if voucher.KisUVoucherEntry[m].FAccountID, ok = accountM[voucher.KisUVoucherEntry[m].FAccountName]; !ok {
			app.Logger.Errorf("%s not find in account", voucher.KisUVoucherEntry[m].FAccountName)
			return errors.New(voucher.KisUVoucherEntry[m].FAccountName + " not find in account")
		}
		if voucher.KisUVoucherEntry[m].FAccountID2, ok = accountM[voucher.KisUVoucherEntry[m].FAccountName2]; !ok {
			app.Logger.Errorf("%s not find in account", voucher.KisUVoucherEntry[m].FAccountName2)
			return errors.New(voucher.KisUVoucherEntry[m].FAccountName2 + " not find in account")
		}

		// 补全金额
		bwbId, err := strconv.Atoi(bwb.FValue)
		if err != nil {
			return errors.New(bwb.FValue + " not int  ")
		}

		// 币别
		if currencyM[voucher.KisUVoucherEntry[m].FCurrency] == bwbId {
			voucher.KisUVoucherEntry[m].FAmount = voucher.KisUVoucherEntry[m].FAmountFor
			voucher.KisUVoucherEntry[m].FExchangeRate = 1
		} else {
			exchangeRate := currencyM2[voucher.KisUVoucherEntry[m].FCurrency]
			voucher.KisUVoucherEntry[m].FAmount = voucher.KisUVoucherEntry[m].FAmountFor * exchangeRate
			voucher.KisUVoucherEntry[m].FExchangeRate = exchangeRate
		}
		voucher.KisUVoucherEntry[m].FCurrencyID = currencyM[voucher.KisUVoucherEntry[m].FCurrency]
		// FInternalInd
	}
	return nil
}

// 获取内码
func kisUGetId(tableName string) (int, error) {
	var id int
	var maxNum KisUMaxNum
	var identity KisUIdentity

	err := MsDB.Where("FTableName = ?", tableName).Find(&maxNum).Error
	if err != nil {
		app.Logger.Errorln("failed to find icMaxNum: ", err)
		return 0, err
	}

	err = MsDB.Where("FName = ?", tableName).Find(&identity).Error
	if err != nil {
		app.Logger.Errorln("failed to find identity: ", err)
		return 0, err
	}

	if identity.FNext >= maxNum.FMaxNum {
		id = identity.FNext
		maxNum.FMaxNum = identity.FNext + identity.FStep
		identity.FNext = identity.FNext + identity.FStep
	} else {
		id = maxNum.FMaxNum
		maxNum.FMaxNum = maxNum.FMaxNum + identity.FStep
		identity.FNext = maxNum.FMaxNum + identity.FStep
	}

	// update KisU table data
	err = MsDB.Model(&maxNum).Where("FTableName = ?", tableName).Updates(&maxNum).Error
	if err != nil {
		app.Logger.Errorln("failed to update ICMaxNum: ", err)
		return 0, err
	}

	err = MsDB.Model(&identity).Where("FName = ?", tableName).Updates(&identity).Error
	if err != nil {
		app.Logger.Errorln("failed to update identity: ", err)
		return 0, err
	}

	return id, nil
}

// 获取凭证号
func kisUGetNumber(year, period, groupId int) (int, error) {
	var maxNum int
	nums := make([]KisUFNum, 0)

	err = MsDB.Where("FYear = ? AND FPeriod = ? and FGroupID = ?", year, period, groupId).Find(&nums).Error
	if err != nil && err != sql.ErrNoRows {
		app.Logger.Errorln("failed to find voucher  FNumber : ", err)
		return 0, err
	} else if err == sql.ErrNoRows {
		return 0, nil
	}

	for i := 0; i < len(nums); i++ {
		if nums[i].FNumber > maxNum {
			maxNum = nums[i].FNumber
		}
	}
	return maxNum + 1, nil
}
