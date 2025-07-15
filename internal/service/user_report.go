package service

import (
	"app/internal/model"
	"app/internal/repository"
	"app/internal/response"
	"bytes"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

type IUserReportService interface {
	GenerateExcel(users *[]model.User) ([]byte, error)
}

type userReportService struct {
	userRepository repository.IUserRepository
}

func NewUserReportService(
	userRepository repository.IUserRepository,
) IUserReportService {
	return &userReportService{userRepository}
}

func (s *userReportService) GenerateExcel(users *[]model.User) ([]byte, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	headers := []string{"ID", "姓名", "使用者名稱", "電子郵件"}
	maxLens := make([]int, len(headers))

	// 寫入標題列
	for colIdx, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
		f.SetCellValue(sheet, cell, header)
		if len(header) > maxLens[colIdx] {
			maxLens[colIdx] = len(header)
		}
	}

	// 寫入每筆 user 資料
	for rowIdx, user := range *users {
		values := []string{strconv.Itoa(int(user.ID)), user.Names, user.Username, user.Email}
		for colIdx, val := range values {
			cell, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
			f.SetCellValue(sheet, cell, val)

			if len(val) > maxLens[colIdx] {
				maxLens[colIdx] = len(val)
			}
		}
	}

	// 設定欄寬
	for colIdx, maxLen := range maxLens {
		colName, _ := excelize.ColumnNumberToName(colIdx + 1)
		f.SetColWidth(sheet, colName, colName, float64(maxLen)*1.2)
	}

	// 將 excelize.File 寫入 bytes.Buffer
	buf := new(bytes.Buffer)
	if err := f.Write(buf); err != nil {
		return nil, response.NewErrorRes(fiber.StatusInternalServerError, []string{"無法寫入 Excel 檔案"})
	}

	return buf.Bytes(), nil
}
