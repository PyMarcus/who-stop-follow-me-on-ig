package internal

import (
	"html/template"
	"net/http"
	"path"
)

type Data struct {
	Names []string
	Image string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	templatePath := path.Join("template", "index.html")
	tmpl, err := template.ParseFiles(templatePath)

	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, Data{})
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	insta, err := LoginService(email, password)
	if err != nil {
		data := Data{Names: []string{"Falha ao realizar login."}, Image: ""}
		tmpl.Execute(w, data)
		return
	}
	followersList := GetFollowers(insta)
	followingList := GetFollowing(insta)
	names := Compare(followersList, followingList)

	data := Data{Names: names, Image: ""}

	tmpl.Execute(w, data)

}

func Render() {
	http.HandleFunc("/", renderTemplate)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("template/assets"))))
	http.ListenAndServe(":8080", nil)
}
