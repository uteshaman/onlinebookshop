package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Product struct {
	ID       int
	Name     string
	Price    int
	Rating   float32
	NumRater int
}

var products = []Product{
	Product{1, "Golang language", 12000, 5, 10},
	Product{2, "Java language", 11000, 4.2, 20},
	Product{3, "Python language", 9990, 4.8, 30},
	Product{4, "Php", 400, 3.9, 15},
	Product{5, "Js", 500, 4.5, 25},
}

var bookmarks = []Product{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, products)
}

func bookmarksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("bookmarks.html"))
	tmpl.Execute(w, bookmarks)
}

func addBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, p := range products {
		if fmt.Sprint(p.ID) == id {
			bookmarks = append(bookmarks, p)
			break
		}
	}
	http.Redirect(w, r, "/bookmarks", http.StatusFound)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for i, p := range bookmarks {
		if fmt.Sprint(p.ID) == id {
			bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/bookmarks", http.StatusFound)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	minPrice := r.FormValue("minPrice")
	maxPrice := r.FormValue("maxPrice")
	filtered := []Product{}
	for _, p := range products {
		if p.Price >= convertToInt(minPrice) && p.Price <= convertToInt(maxPrice) {
			filtered = append(filtered, p)
		}
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, filtered)
}

func convertToInt(str string) int {
	if str == "" {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func rateProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	rating := r.URL.Query().Get("rating")
	for i, p := range products {
		if fmt.Sprint(p.ID) == id {
			numRater := p.NumRater
			oldRating := p.Rating
			newRating := float32(convertToInt(rating))
			p.Rating = ((oldRating * float32(numRater)) + newRating) / float32(numRater+1)
			p.NumRater++
			products[i] = p
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bookmarks", bookmarksHandler)
	http.HandleFunc("/addBookmark", addBookmarkHandler)
	http.HandleFunc("/deleteBookmark", deleteBookmarkHandler)
	http.HandleFunc("/filter", filterHandler)
	http.HandleFunc("/rateProduct", rateProductHandler)
	fmt.Println("Server started on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
