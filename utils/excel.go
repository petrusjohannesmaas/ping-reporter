package utils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// WriteToExcel creates an Excel file with the extracted data
func WriteToExcel(readings []Reading, filename string) error {
	// Create a new Excel file
	f := excelize.NewFile()

	// Create a new sheet
	sheetName := "Results"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("error creating new sheet: %w", err)
	}

	// Set the headers
	headers := []string{"Record ID", "PPP Username", "IP", "Identity", "Avg RTT (ms)"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i)) // Generates cell names A1, B1, etc.
		f.SetCellValue(sheetName, cell, header)
	}

	// Populate the data rows
	for i, reading := range readings {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), reading.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), reading.PPPusername)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), reading.IP)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), reading.Identity)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), reading.AvgRtt)
	}

	// Set the active sheet to the one we created
	f.SetActiveSheet(index)

	// Save the Excel file
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("error saving Excel file: %w", err)
	}

	fmt.Printf("Excel file successfully saved as: %s\n", filename)
	return nil
}
