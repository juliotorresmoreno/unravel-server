package auth

import (
	"bytes"
	"log"
	"net/http"
	"net/smtp"
)

import "../../config"

//import "../../helper"
//import "../../models"

// Recovery recuperacion de contraseña
func Recovery(w http.ResponseWriter, r *http.Request) {
	var email = r.PostFormValue("email")
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(config.SMTP_HOST + ":" + config.SMTP_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// Set the sender and recipient.
	c.Mail("recovery@" + config.HOSTNAME)
	c.Rcpt(email)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
