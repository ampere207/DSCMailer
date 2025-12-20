package handlers

import (
	"DSCMailer/internal/parser"
	"html/template"
	"net/http"
)

var templates = template.Must(
	template.ParseFiles(
		"web/templates/index.html",
		"web/templates/success.html",
	),
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type SuccessData struct {
	Recipients []struct {
		Name  string
		Email string
	}
	Count int
}

func RenderSuccess(w http.ResponseWriter, recipients []parser.Recipient) {
	data := SuccessData{
		Count: len(recipients),
	}

	for _, r := range recipients {
		data.Recipients = append(data.Recipients, struct {
			Name  string
			Email string
		}{
			Name:  r.Name,
			Email: r.Email,
		})
	}

	if err := templates.ExecuteTemplate(w, "success.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
