package handler

import (
	"html/template"
	"log"
	"net/http"
)

func TopHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/top.gtpl")
		if err != nil {
			log.Printf("[ERROR] TopHandler template.ParseFiles : %e\n", err)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("host:%s uri:%s ua:%s", r.Host, r.RequestURI, r.UserAgent())
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/form.gtpl")
		if err != nil {
			log.Printf("[ERROR] FormHandler template.ParseFiles : %e\n", err)
		}

		t.Execute(w, nil)
		log.Printf("host:%s uri:%s ua:%s", r.Host, r.RequestURI, r.UserAgent())
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		t, err := template.ParseFiles("./templates/submit.gtpl")
		if err != nil {
			log.Printf("[ERROR] TopHandler template.ParseFiles : %e\n", err)
			return
		}

		p := map[string]string{
			"method": "GET",
			"param1": r.Form.Get("param1"),
			"param2": r.Form.Get("param2"),
		}
		err = t.ExecuteTemplate(w, "submit.gtpl", p)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("host:%s uri:%s ua:%s", r.Host, r.RequestURI, r.UserAgent())
	} else if r.Method == "POST" {
		r.ParseForm()
		t, err := template.ParseFiles("./templates/submit.gtpl")
		if err != nil {
			log.Printf("[ERROR] SubmitHandler template.ParseFiles : %e\n", err)
		}

		p := map[string]string{
			"method": "POST",
			"param1": r.Form.Get("param1"),
			"param2": r.Form.Get("param2"),
		}
		t.Execute(w, p)
		log.Printf("host:%s uri:%s ua:%s", r.Host, r.RequestURI, r.UserAgent())
	}
}
