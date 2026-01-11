package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

func BuildMessage(from string, fromName string, to string, subject string, body string, attachmentPath string) ([]byte, error) {

	var buf bytes.Buffer
	boundary := "CERT_MAILER_BOUNDARY"

	// Format From header with name if provided
	fromHeader := from
	if fromName != "" {
		fromHeader = fmt.Sprintf("%s <%s>", fromName, from)
	}

	// Headers
	buf.WriteString(fmt.Sprintf("From: %s\r\n", fromHeader))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", to))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf(
		"Content-Type: multipart/mixed; boundary=%s\r\n\r\n",
		boundary,
	))

	// Body
	buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")
	buf.WriteString(body + "\r\n")

	// Attachment
	fileBytes, err := os.ReadFile(attachmentPath)
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(fileBytes)
	filename := filepath.Base(attachmentPath)

	buf.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	buf.WriteString("Content-Type: application/octet-stream\r\n")
	buf.WriteString("Content-Transfer-Encoding: base64\r\n")
	buf.WriteString(fmt.Sprintf(
		"Content-Disposition: attachment; filename=\"%s\"\r\n\r\n",
		filename,
	))

	// Base64 must be split into lines
	for i := 0; i < len(encoded); i += 76 {
		end := i + 76
		if end > len(encoded) {
			end = len(encoded)
		}
		buf.WriteString(encoded[i:end] + "\r\n")
	}

	buf.WriteString(fmt.Sprintf("--%s--", boundary))
	return buf.Bytes(), nil
}
