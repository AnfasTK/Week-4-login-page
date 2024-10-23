package main

import (
	"html/template"
	"net/http"
	"github.com/icza/session"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("templates/*.html"))
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil && sess.CAttr("username") != nil {
		http.Redirect(w, r, "/main", http.StatusSeeOther)
		return
	}
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	if r.Method==http.MethodGet {
		temp.ExecuteTemplate(w, "index.html", nil)
	}else{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
}
func loginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodPost {
		username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "Anfas" && password == "Anfas@123" {
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{
				"username": username,
				"password": password,
			},
		})
		session.Add(sess, w)
		http.Redirect(w, r, "/main", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"error": "Invalid Username and Password",
		}
		temp.ExecuteTemplate(w, "index.html", data)
	}
	}else{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
}
func mainHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodGet {	
	w.Header().Set("Cache-Control", "no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	sess := session.Get(r)
	if sess == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} 
	username := sess.CAttr("username")
	data := map[string]interface{}{
		"username": username,
	}
	temp.ExecuteTemplate(w, "main.html", data)
	}else{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
}
func logoutHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost{
		sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	}else{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/main", mainHandle)
	http.HandleFunc("/logout", logoutHandle)
	http.ListenAndServe(":8000", nil)
}
