package excel

import (
	"fmt"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type ExcelService struct {
}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (e *ExcelService) GenerateSheetForStatement(statement *types.Statement) error {
	f, err := excelize.OpenFile("simple.xlsx")
	if err != nil {
		logrus.Error(err)
		return err
	}

	sheetName1 := "Sheet1"
	for i, s := range statement.StatementRecords {
		i++ // Excel sheet starts with 1
		if err := f.SetCellValue(sheetName1, fmt.Sprintf("A%d", i), s.Description); err != nil {
			logrus.Error(err)
		}

		if s.Amount > 0 {
			// I wish to display only expenses
			continue
		}
		// I wish to display expenses as non-negative numbers
		if err := f.SetCellValue(sheetName1, fmt.Sprintf("B%d", i), s.Amount*-1/100); err != nil {
			logrus.Error(err)
		}
	}

	if err := f.SaveAs("simple.xlsx"); err != nil {
		return err
	}

	return nil
}
