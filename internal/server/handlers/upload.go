package handlers

import (
	"DSCMailer/internal/certificate"
	"DSCMailer/internal/mailer"
	"DSCMailer/internal/parser"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) { // Only POST allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Invalid upload", http.StatusBadRequest)
		return
	}

	// Get template file
	templateFile, templateHeader, err := r.FormFile("template")
	if err != nil {
		http.Error(w, "Template file not found", http.StatusBadRequest)
		return
	}
	defer templateFile.Close()

	// Save template temporarily
	templatePath := filepath.Join("uploads", templateHeader.Filename)
	os.MkdirAll("uploads", 0755)

	templateOut, err := os.Create(templatePath)
	if err != nil {
		http.Error(w, "Failed to save template", http.StatusInternalServerError)
		return
	}
	defer templateOut.Close()

	_, err = templateOut.ReadFrom(templateFile)
	if err != nil {
		http.Error(w, "Failed to save template", http.StatusInternalServerError)
		return
	}

	// Get recipients file
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Recipients file not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Decide parser
	ext := strings.ToLower(filepath.Ext(header.Filename))

	var recipients []parser.Recipient
	if ext == ".csv" {
		recipients, err = parser.ParseCSV(file)
	} else if ext == ".xlsx" {
		recipients, err = parser.ParseXLSX(file)
	} else {
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to parse file", http.StatusBadRequest)
		return
	}

	// Get email subject and body from form
	subjectTemplate := r.FormValue("subject")
	bodyTemplate := r.FormValue("body")

	if subjectTemplate == "" || bodyTemplate == "" {
		http.Error(w, "Email subject and body are required", http.StatusBadRequest)
		return
	}

	// Process each recipient and track successful sends
	var successful []parser.Recipient

	for _, recipient := range recipients {

		certPath, err := certificate.GenerateWithTemplate(recipient.Name, templatePath)
		if err != nil {
			log.Println("Certificate error:", err)
			continue
		}

		// Personalize subject and body with recipient name
		subject := strings.ReplaceAll(subjectTemplate, "{name}", recipient.Name)
		body := strings.ReplaceAll(bodyTemplate, "{name}", recipient.Name)

		// Wrap body in HTML with pre-wrap to preserve formatting
		body = "<div style=\"white-space: pre-wrap; font-family: Arial, sans-serif;\">" + body + "</div>"

		err = mailer.SendSMTP(
			recipient.Email,
			subject,
			body,
			certPath,
		)

		if err != nil {
			log.Println("Mail failed:", err)
			continue
		}

		// Add to successful list
		successful = append(successful, recipient)
	}

	// Render success page
	RenderSuccess(w, successful)
}
