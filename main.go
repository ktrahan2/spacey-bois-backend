package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

//HighScore is the structure of the highscore data
type HighScore struct {
	gorm.Model
	//gorm assumes my ID is the primary id field
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

var (
	highscores = []HighScore{
		{Username: "me", Score: 900},
	}
)

func main() {
	connectToDatabase()
	db.AutoMigrate(&HighScore{})
	handleRequest()

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

	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Successfully connected!")
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/highscores", returnAllHighScores).Methods("GET")
	router.HandleFunc("/addscores", addScores).Methods("POST")
	router.HandleFunc("/addscores", addScores).Methods("OPTIONS")
	http.ListenAndServe(":7000", router)
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func returnAllHighScores(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	var highscores []HighScore

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM highscores")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var highscore HighScore
		if err := rows.Scan(&highscore.ID, &highscore.Username, &highscore.Score); err != nil {
			log.Fatal(err)
		}
		highscores = append(highscores, highscore)
	}
	json.NewEncoder(w).Encode(highscores)
}

func addScores(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		log.Println(r)
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
	default:
		http.Error(w, http.StatusText(405), 405)
	}

}
