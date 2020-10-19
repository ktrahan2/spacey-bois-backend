package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	//defer closes the database whenever this function ends
	defer db.Close()
}

func connectToDatabase() {

	var err error
	host := os.Getenv("HOST")
	databaseUsername := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, databaseUsername, password, database)
	//doesn't like this after I closed my VS code down. Had to use 'unset PGSYSCONFDIR'
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//can probably delete this func when finished
func homePage(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	fmt.Fprintf(w, "Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllHighScores(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM highscores")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var highscores []HighScore
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
		//making multiple slices of a single object instead of one single slice,
		//can probably refactor to iterate through and push to Highscores variable, instead of making multiple
		highscore := HighScore{ID: id, Username: username, Score: score}
		highscores = append(highscores, highscore)
	}
	json.NewEncoder(w).Encode(highscores)
}

func addScores(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%v", string(reqBody))
	var highscore HighScore
	json.Unmarshal(reqBody, &highscore)

	username := highscore.Username
	score := highscore.Score

	var err error
	sqlStatement := `
	INSERT INTO highscores (username, score)
	VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, username, score)
	if err != nil {
		panic(err)
	}

	// fmt.Fprintf(w, "Add scores hit")
}
