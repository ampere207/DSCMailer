package mailer

import (
	"bytes"
	"html/template"
)

type EmailData struct {
	Name string
}

var emailTemplate = template.Must(template.New("email").Parse(`
<!DOCTYPE html>
<html>
  <body style="font-family: Arial, sans-serif;">
    <h2>Hello {{.Name}},</h2>
    <p>
      Congratulations!<br>
      Please find your certificate attached.
    </p>
    <p>
      Regards,<br>
      <strong>Developer Student Club</strong>
    </p>
  </body>
</html>
`))

func RenderEmail(name string) (subject string, body string, err error) {
	subject = "Your Certificate 🎓"

	var buf bytes.Buffer
	err = emailTemplate.Execute(&buf, EmailData{
		Name: name,
	})

	return subject, buf.String(), err
}
