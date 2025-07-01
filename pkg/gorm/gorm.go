package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// 將 GORM 錯誤轉換為中文
func GormErrorToMessage(err error) string {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return "查無資料"
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return "無效的交易操作"
	case errors.Is(err, gorm.ErrNotImplemented):
		return "此功能尚未實作"
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return "缺少 WHERE 條件，操作無效"
	case errors.Is(err, gorm.ErrUnsupportedRelation):
		return "不支援的關聯操作"
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return "需要主鍵欄位"
	case errors.Is(err, gorm.ErrModelValueRequired):
		return "必須提供模型數據"
	case errors.Is(err, gorm.ErrModelAccessibleFieldsRequired):
		return "模型必須指定可存取的欄位"
	case errors.Is(err, gorm.ErrSubQueryRequired):
		return "需要子查詢"
	case errors.Is(err, gorm.ErrInvalidData):
		return "不支援的資料格式"
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return "不支援的資料庫驅動"
	case errors.Is(err, gorm.ErrRegistered):
		return "已註冊，請勿重複操作"
	case errors.Is(err, gorm.ErrInvalidField):
		return "無效的欄位"
	case errors.Is(err, gorm.ErrEmptySlice):
		return "傳入的資料為空"
	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return "不支援模擬執行模式"
	case errors.Is(err, gorm.ErrInvalidDB):
		return "資料庫無效或未正確初始化"
	case errors.Is(err, gorm.ErrInvalidValue):
		return "無效的值，應為 struct 或 slice 的指標"
	case errors.Is(err, gorm.ErrInvalidValueOfLength):
		return "關聯資料長度不一致"
	case errors.Is(err, gorm.ErrPreloadNotAllowed):
		return "使用 count 查詢時不允許預載入"
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return "資料重複，違反唯一性限制"
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return "外鍵約束錯誤，關聯資料不存在"
	case errors.Is(err, gorm.ErrCheckConstraintViolated):
		return "資料不符合檢查約束條件"
	default:
		return "資料庫錯誤：" + err.Error()
	}
}
