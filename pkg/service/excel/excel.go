package excel

import (
	"fmt"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"github.com/xuri/excelize/v2"
)

type ExcelService struct {
}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (e *ExcelService) GenerateSheetForStatement(statement *types.Statement) error {
	f := excelize.NewFile()

	sheetName1 := "Sheet1"
	for i, j := 0, 1; i < len(statement.StatementRecords); i++ {
		if statement.StatementRecords[i].Amount > 0 {
			// I wish to display only expenses
			continue
		}

		err := f.SetCellValue(sheetName1, fmt.Sprintf("A%d", j), statement.StatementRecords[i].Description)
		if err != nil {
			return err
		}
		// I wish to display expenses as non-negative numbers
		err = f.SetCellValue(sheetName1, fmt.Sprintf("B%d", j), statement.StatementRecords[i].Amount*-1/100)
		if err != nil {
			return err
		}
		j++
	}

	if err := f.SaveAs("simple.xlsx"); err != nil {
		return err
	}

	return nil
}
