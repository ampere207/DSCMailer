package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Recipient struct {
	Name   string
	Email  string
	Status string
}

func ParseCSV(r io.Reader) ([]Recipient, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("csv has no data rows")
	}

	// Find name, email, and status column indices
	header := rows[0]
	nameIdx, emailIdx, statusIdx := -1, -1, -1

	for i, col := range header {
		colLower := strings.ToLower(strings.TrimSpace(col))
		if colLower == "name" {
			nameIdx = i
		}
		if colLower == "email" {
			emailIdx = i
		}
		if colLower == "status" {
			statusIdx = i
		}
	}

	if nameIdx == -1 || emailIdx == -1 {
		return nil, fmt.Errorf("csv must have 'name' and 'email' columns")
	}

	var recipients []Recipient
	seenEmails := make(map[string]bool)

	// Skip header (row 0)
	for i := 1; i < len(rows); i++ {
		row := rows[i]

		if len(row) <= nameIdx || len(row) <= emailIdx {
			continue
		}

		name := strings.TrimSpace(row[nameIdx])
		email := strings.TrimSpace(row[emailIdx])

		// Check status if column exists
		var status string
		if statusIdx != -1 && len(row) > statusIdx {
			status = strings.TrimSpace(row[statusIdx])
			// Only include recipients with "Certificate to be given" status
			if strings.ToLower(status) != "certificate to be given" {
				continue
			}
		}

		if name == "" || email == "" {
			continue
		}

		// Skip duplicate emails
		if seenEmails[email] {
			continue
		}
		seenEmails[email] = true

		// Truncate name to 30 characters
		if len(name) > 30 {
			name = name[:30]
		}

		recipients = append(recipients, Recipient{
			Name:   name,
			Email:  email,
			Status: status,
		})
	}

	return recipients, nil
}

func ParseXLSX(r io.Reader) ([]Recipient, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("xlsx has no data rows")
	}

	// Find name, email, and status column indices
	header := rows[0]
	nameIdx, emailIdx, statusIdx := -1, -1, -1

	for i, col := range header {
		colLower := strings.ToLower(strings.TrimSpace(col))
		if colLower == "name" {
			nameIdx = i
		}
		if colLower == "email" {
			emailIdx = i
		}
		if colLower == "status" {
			statusIdx = i
		}
	}

	if nameIdx == -1 || emailIdx == -1 {
		return nil, fmt.Errorf("xlsx must have 'name' and 'email' columns")
	}

	var recipients []Recipient
	seenEmails := make(map[string]bool)

	// Skip header
	for i := 1; i < len(rows); i++ {
		row := rows[i]

		if len(row) <= nameIdx || len(row) <= emailIdx {
			continue
		}

		name := strings.TrimSpace(row[nameIdx])
		email := strings.TrimSpace(row[emailIdx])

		// Check status if column exists
		var status string
		if statusIdx != -1 && len(row) > statusIdx {
			status = strings.TrimSpace(row[statusIdx])
			// Only include recipients with "Certificate to be given" status
			if strings.ToLower(status) != "certificate to be given" {
				continue
			}
		}

		if name == "" || email == "" {
			continue
		}

		// Skip duplicate emails
		if seenEmails[email] {
			continue
		}
		seenEmails[email] = true

		// Truncate name to 30 characters
		if len(name) > 30 {
			name = name[:30]
		}

		recipients = append(recipients, Recipient{
			Name:   name,
			Email:  email,
			Status: status,
		})
	}

	return recipients, nil
}
