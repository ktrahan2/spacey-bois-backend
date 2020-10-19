package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

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
	connectToDatabase()
	handleRequest()
}

func connectToDatabase() {
	host := os.Getenv("HOST")
	databaseUsername := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, databaseUsername, password, database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer closes the database whenever this function ends
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func handleRequest() {
	http.HandleFunc("/highscores", returnAllHighScores)
	http.HandleFunc("/addscores", addScores)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":7000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllHighScores(w http.ResponseWriter, r *http.Request) {

	//needs sql.Open() to be a global variable
	rows, err := db.Query("SELECT * FROM highscores")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var (
			id       int
			username string
			score    int
		)
		if err := rows.Scan(&id, &username, &score); err != nil {
			log.Fatal(err)
		}
		log.Printf("id: %d, username: %s, score: %d", id, username, score)
		HighScores := []HighScore{
			HighScore{ID: id, Username: username, Score: score},
		}
		fmt.Println(HighScores)
		// json.NewEncoder(w).Encode(HighScores)
	}
	fmt.Println("Endpoint hit")
}

func addScores(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add scores hit")
}
