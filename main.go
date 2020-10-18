package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

//HighScore is the structure of the highscore data
type HighScore struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

func main() {
	host := os.Getenv("HOST")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	fmt.Println(host, username, password, dbname, port)
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest() {
	http.HandleFunc("/highscores", returnAllHighScores)
	http.HandleFunc("/addscores", addScores)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":7000", nil))
}

func returnAllHighScores(w http.ResponseWriter, r *http.Request) {
	HighScores := []HighScore{
		HighScore{ID: 1, Username: "ktrain", Score: 100},
		HighScore{ID: 2, Username: "nnpalmore", Score: 200},
	}
	fmt.Println("Endpoint hit")
	json.NewEncoder(w).Encode(HighScores)
}

func addScores(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add scores hit")
}
