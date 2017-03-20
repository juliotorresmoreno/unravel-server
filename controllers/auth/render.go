package auth

import "../../config"

type page struct {
	Title string
	Link  string
}

type result struct {
	html string
}

type writer struct {
	result *result
}

func (w writer) Write(p []byte) (n int, err error) {
	w.result.html = w.result.html + string(p)
	return 0, nil
}

func (w writer) String() string {
	return w.result.html
}

// render renderiza el correo de recuperacion
func render(email, id string) []byte {
	p := page{
		Title: "Recuperacion de contraseñas",
		Link:  "https://" + config.HOSTNAME + "/recuperar?id=" + id,
	}
	b := writer{result: &result{}}
	templates.ExecuteTemplate(b, "recovery.html", p)
	body := []byte(
		"To: " + email + "\r\n" +
			"Subject: Recuperacion de contraseña!\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
			b.String() + "\r\n",
	)
	return body
}
