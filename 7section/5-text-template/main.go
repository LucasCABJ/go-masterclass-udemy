package main

import (
	"fmt"
	"os"
	"text/template"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string
	UnreadCount   int
}

func main() {

	fmt.Println("--- Text template example ---")

	emailTemplate := `
Subject: {{ .Subject }}	

{{.Body}}

{{if .Items}}
Related Items:
{{range .Items}}
- {{.}}
{{end}}
{{end}}

{{if gt .UnreadCount 0}}
You have {{.UnreadCount}} unreads.
{{else}}
You have no messages
{{end}}

- Thanks
{{.SenderName}}
`

	tmpl, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := EmailData{
		RecipientName: "Lucas",
		SenderName:    "Gabriel",
		Subject:       "Weekly Update",
		Body:          "Lorem ipsum lorem ipsum lorem ipsum",
		Items:         []string{"Item A", "Item B", "Item C"},
		UnreadCount:   5,
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
