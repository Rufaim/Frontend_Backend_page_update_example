package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Server struct {
	IndexTemplate      *template.Template
	UserTemplate       *template.Template
	CurrentApiResponce *ApiResponce
}

func getIndexHandler(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		indexTmpl, err := s.IndexTemplate.Clone()
		if err != nil {
			panic(err)
		}
		err = indexTmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

func getUserHandler(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiResp, err := getRandomUser()
		if err != nil {
			panic(err)
		}
		indexTmpl, err := s.UserTemplate.Clone()
		if err != nil {
			panic(err)
		}
		err = indexTmpl.Execute(w, apiResp)
		if err != nil {
			panic(err)
		}
	}
}

func parseTemplate(path string) *template.Template {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func main() {
	indexTmpl := parseTemplate("./html/index.html")
	userTmpl := parseTemplate("./html/user.html")
	s := &Server{
		IndexTemplate: indexTmpl,
		UserTemplate:  userTmpl,
	}

	fmt.Println("Building router...")
	jsFS := http.FileServer(http.Dir(DirJSFiles))
	cssFS := http.FileServer(http.Dir(DirCSSFiles))

	http.Handle("/js/", http.StripPrefix("/js/", jsFS))
	http.Handle("/css/", http.StripPrefix("/css/", cssFS))
	http.HandleFunc("/", getIndexHandler(s))
	http.HandleFunc("/user", getUserHandler(s))

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
