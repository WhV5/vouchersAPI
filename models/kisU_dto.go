/**
* @Author : henry
* @Data: 2020-08-13 15:47
* @Note: 金蝶凭证的表
**/

package models

import "time"

type Test struct {
	ID   int
	Name string
	Age  int
}

type KisUVoucher struct {
	FApproveID       int       `gorm:"-"`
	FAttachments     int       `gorm:"column:FAttachments"`
	FBrNo            string    `gorm:"-"`
	FCashierID       int       `gorm:"column:FCashierID"`
	FChecked         bool      `gorm:"column:FChecked"`
	FCheckerID       int       `gorm:"column:FCheckerID"`
	FCreditTotal     float64   `gorm:"column:FCreditTotal" json:"fCreditTotal"`
	FDate            time.Time `gorm:"column:FDate" json:"fDate"`
	FDebitTotal      float64   `gorm:"column:FDebitTotal" json:"fDebitTotal"`
	FEntryCount      int       `gorm:"column:FEntryCount"`
	FExplanation     string    `gorm:"column:FExplanation" json:"fExplanation"`
	FFootNote        string    `gorm:"column:FFootNote"`
	FFrameWorkID     int       `gorm:"column:FFrameWorkID"`
	FGroupID         int       `gorm:"column:FGroupID"`
	FHandler         string    `gorm:"column:FHandler" json:"fHandler"`
	FInternalInd     string    `gorm:"column:FInternalInd"`
	FModifyTime      []uint8   `gorm:"-"`
	FNumber          int       `gorm:"column:FNumber"`
	FObjectName      string    `gorm:"column:FObjectName" json:"fObjectName"`
	FOwnerGroupID    int       `gorm:"column:FOwnerGroupID"`
	FParameter       string    `gorm:"column:FParameter"`
	FPeriod          int       `gorm:"column:FPeriod"`
	FPosted          bool      `gorm:"column:FPosted"`
	FPosterID        int       `gorm:"column:FPosterID"`
	FPreparerID      int       `gorm:"column:FPreparerID"`
	FReference       string    `gorm:"-"`
	FSerialNum       int       `gorm:"column:FSerialNum"`
	FTransDate       time.Time `gorm:"column:FTransDate" json:"fTransDate"`
	FTranType        int       `gorm:"column:FTranType" json:"fTranType"`
	FVoucherID       int       `gorm:"column:FVoucherID"`
	FYear            int       `gorm:"column:FYear"`
	FCashier         string    `gorm:"-" json:"fCashier"`
	FChecker         string    `gorm:"-" json:"fChecker"`
	FPoster          string    `gorm:"-" json:"fPoster"`
	FPreparer        string    `gorm:"-" json:"fPreparer"`
	KisUVoucherEntry []*KisUVoucherEntry
}

type KisUVoucherEntry struct {
	FAccountID     int     `gorm:"column:FAccountID" `
	FAccountID2    int     `gorm:"column:FAccountID2 " `
	FAmount        float64 `gorm:"column:FAmount " `
	FAmountFor     float64 `gorm:"column:FAmountFor" json:"fAmountFor"`
	FBrNo          string  `gorm:"column:FBrNo " `
	FCashFlowItem  int     `gorm:"column:FCashFlowItem " `
	FCurrencyID    int     `gorm:"column:FCurrencyID " `
	FDC            int     `gorm:"column:FDC " json:"fDc"`
	FDetailID      int     `gorm:"column:FDetailID " `
	FEntryID       int     `gorm:"column:FEntryID" `
	FExchangeRate  float64 `gorm:"column:FExchangeRate " `
	FExplanation   string  `gorm:"column:FExplanation" json:"fExplanation"`
	FInternalInd   string  `gorm:"column:FInternalInd" `
	FMeasureUnitID int     `gorm:"column:FMeasureUnitID" `
	FQuantity      float64 `gorm:"column:FQuantity " `
	FResourceID    int     `gorm:"column:FResourceID " `
	FSettleNo      string  `gorm:"column:FSettleNo " `
	FSettleTypeID  int     `gorm:"column:FSettleTypeID " `
	FTaskID        int     `gorm:"column:FTaskID " `
	FTransNo       string  `gorm:"column:FTransNo" `
	FUnitPrice     float64 `gorm:"column:FUnitPrice" `
	FVoucherID     int     `gorm:"column:FVoucherID" `
	FAccountName   string  `gorm:"-" json:"fAmountName"`
	FAccountName2  string  `gorm:"-" json:"fAccountName2"`
	FCurrency      string  `gorm:"-" json:"fCurrency"`
}

type KisUUser struct {
	FUserID int    `gorm:"column:FUserID"`
	FName   string `gorm:"column:FName"`
}

type KisUAccount struct {
	FAccountID int    `gorm:"column:FAccountID"`
	FNumber    string `gorm:"column:FNumber"`
}

type KisUCurrency struct {
	FCurrencyID   int     `gorm:"column:FCurrencyID"`
	FNumber       string  `gorm:"column:FNumber"`
	FExchangeRate float64 `gorm:"column:FExchangeRate"`
}

type KisUSysPro struct {
	FValue    string `gorm:"column:FValue"`
	FCategory string `gorm:"column:FCategory"`
	FKey      string `gorm:"column:FKey"`
}

type KisUMaxNum struct {
	FTableName string `gorm:"column:FTableName"`
	FMaxNum    int    `gorm:"column:FMaxNum"`
}

type KisUIdentity struct {
	FName string `gorm:"column:FName"`
	FNext int    `gorm:"column:FNext"`
	FStep int    `gorm:"column:FStep"`
}

func (kisUVoucher KisUVoucher) TableName() string {
	return "t_Voucher"
}

func (kisUVoucherEntry KisUVoucherEntry) TableName() string {
	return "t_VoucherEntry"
}

func (kisUUser KisUUser) TableName() string {
	return "t_User"
}

func (kisUAccount KisUAccount) TableName() string {
	return "t_Account"
}

func (kisUCurrency KisUCurrency) TableName() string {
	return "t_Currency"
}

func (kisUSysPro KisUSysPro) TableName() string {
	return "t_SystemProfile"
}

func (kisUMaxNum KisUMaxNum) TableName() string {
	return "IcMaxNum"
}

func (kisUIdentity KisUIdentity) TableName() string {
	return "t_Identity"
}

type KisUFNum struct {
	FNumber int `gorm:"column:FNumber"`
}

func (kisUNum KisUFNum) TableName() string {
	return "t_Voucher"
}
