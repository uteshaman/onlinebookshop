package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Data struct {
	Tasks []string
	Error string
}

var tasks []string

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.HandleFunc("/register", handleRegister)
	// http.HandleFunc("/logout", handleLogout)



	http.HandleFunc("/", tasksHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/search", searchHandler)

	fmt.Println("Successfully run!")
	http.ListenAndServe(":8080", nil)
}







func tasksHandler(w http.ResponseWriter, r *http.Request) {

	search := strings.TrimSpace(r.URL.Query().Get("search"))

	// if !isLoggedIn(r) {
	// 	http.Redirect(w, r, "/register", http.StatusSeeOther)
	// 	return
	// }

	// else:

	filteredTasks := tasks
	if search != "" {
		filteredTasks = []string{}
		for _, task := range tasks {
			if strings.Contains(task, search) {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	tmpl, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Data{
		Tasks: filteredTasks,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


// // ---Register------
// func handleRegister(w http.ResponseWriter, r *http.Request) {
// 	if isLoggedIn(r) {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}


// 	if r.Method == "GET" {
// 		tmpl, err := template.ParseFiles("templates/register.html")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		err = tmpl.Execute(w, nil)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}


	// If the request method is POST, handle the registration form submission
// 	if r.Method == "POST" {
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")

// 		if _, ok := users[username]; ok {
// 			http.Error(w, "Username is already taken", http.StatusBadRequest)
// 			return
// 		}

// 		users[username] = password

// 		logged_In(w, username)
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
// }


// func logged_In(w http.ResponseWriter, username string) {
// 	http.SetCookie(w, &http.Cookie{
// 		Name:  "loggedIn",
// 		Value: username,
// 		Path:  "/",
// 	})
// }


func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		task := r.FormValue("task")


		if task == "" {
			http.Redirect(w, r, "/?error=Please enter a task", http.StatusSeeOther)
			return
		}


		tasks = append(tasks, task)


		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		indexStr := r.FormValue("index")


		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(tasks) {
			http.Redirect(w, r, "/?error=Invalid task index", http.StatusSeeOther)
			return
		}


		tasks = append(tasks[:index], tasks[index+1:]...)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		search := r.FormValue("search")


		http.Redirect(w, r, "/?search="+search, http.StatusSeeOther)
	}
}