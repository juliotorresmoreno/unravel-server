package test

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
)

func post(url string, params []byte) (error, string, []byte) {
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, "", []byte("")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return nil, resp.Status, body
}

func test(w http.ResponseWriter, url string, params []byte) {
	err, status, body := post(url, params)
	if err != nil {
		fmt.Fprintln(w, "url: " + url + ", method: post, status: Error")
		fmt.Fprintln(w, "response Status: " + status)
		fmt.Fprintln(w, "response Body: " + string(body))
		fmt.Fprintln(w, "error: " + err.Error())
	} else {
		fmt.Fprintln(w, "url: " + url + ", method: post, status: OK")
		fmt.Fprintln(w, "response Status: " + status)
		fmt.Fprintln(w, "response Body: " + string(body))
	}
	fmt.Fprintln(w, "")
}

func Test(w http.ResponseWriter, r *http.Request)  {
	var url string
	var params []byte
	w.WriteHeader(http.StatusOK)

	url = "http://localhost:8080/api/v1/auth/registrar"
	params = []byte("nombres=nombres&apellidos=apellidos&" +
		  "email=email@dominio.com&usuario=username&" +
		  "passwd=123456&passwdConfirm=123456")
	//test(w, url, params)

	url = "http://localhost:8080/api/v1/auth/login"
	params = []byte("usuario=username&passwd=123456")
	//test(w, url, params)

	url = "http://localhost:8080/api/v1/auth/login"
	params = []byte("usuario=username&passwd=12345678")
	//test(w, url, params)

	url = "http://localhost:8080/api/v1/chats/mensaje"
	params = []byte("tipo=usuario&usuario=dmonsalve&mensaje=hola+mundo")

	test(w, url, params)
}