package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Source struct {
	Name string `json:"name"`
}

type Article struct {
	Source      Source `json:"source"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

type NewsResponse struct {
	Articles []Article `json:"articles"`
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/search", searchHandler)

	log.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query", http.StatusBadRequest)
		return
	}

	apiKey := "4f9990accd43464c99a20ea63e86ff66"
	url := "https://newsapi.org/v2/everything?q=" + query + "&apiKey=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to get news", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var news NewsResponse
	json.Unmarshal(body, &news)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}
