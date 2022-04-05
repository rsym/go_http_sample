package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func TopHandler(w http.ResponseWriter, r *http.Request) {
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/top.gtpl")
		if err != nil {
			log.Printf("[ERROR] TopHandler template.ParseFiles : %e\n", err)
			return
		}

		p1, err := rd.Get("param1").Result()
		if err != nil {
			log.Printf("[ERROR] %e\n", err)
		}
		p2, err := rd.Get("param2").Result()
		if err != nil {
			log.Printf("[ERROR] %e\n", err)
		}

		p := map[string]string{
			"param1": p1,
			"param2": p2,
		}
		err = t.ExecuteTemplate(w, "top.gtpl", p)
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
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if r.Method == "GET" {
		r.ParseForm()
		t, err := template.ParseFiles("./templates/submit.gtpl")
		if err != nil {
			log.Printf("[ERROR] TopHandler template.ParseFiles : %e\n", err)
			return
		}

		p1 := r.Form.Get("param1")
		p2 := r.Form.Get("param2")

		err = rd.Set("param1", p1, 0).Err()
		if err != nil {
			log.Printf("[ERROR] %e\n", err)
		}
		err = rd.Set("param2", p2, 0).Err()
		if err != nil {
			log.Printf("[ERROR] %e\n", err)
		}

		p := map[string]string{
			"method": "GET",
			"param1": p1,
			"param2": p2,
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

		p1 := r.Form.Get("param1")
		p2 := r.Form.Get("param2")
		err = rd.Set("param1", p1, 0).Err()
		if err != nil {
			log.Printf("[ERROR] %e\n", err)
		}
		err = rd.Set("param2", p2, 0).Err()
		if err != nil {
			log.Printf("[ERROR] %e\n", err)
		}

		p := map[string]string{
			"method": "POST",
			"param1": p1,
			"param2": p2,
		}
		t.Execute(w, p)
		log.Printf("host:%s uri:%s ua:%s", r.Host, r.RequestURI, r.UserAgent())
	}
}
