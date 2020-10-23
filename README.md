# Spacey Bois

Spacey Bois is a text-based adventure game. 

# Table Of Contents 
- [Description](https://github.com/ktrahan2/spacey-bois-backend/tree/main#description)
- [Example Code](https://github.com/ktrahan2/spacey-bois-backend/tree/main#example-code)
- [Technology Used](https://github.com/ktrahan2/spacey-bois-backend/tree/main#technology-used)
- [Setting up for the Application](https://github.com/ktrahan2/spacey-bois-backend/tree/main#setting-up-for-the-application)
- [Main Features](https://github.com/ktrahan2/spacey-bois-backend/tree/main#main-features)
- [Features in Progress](https://github.com/ktrahan2/spacey-bois-backend/tree/main#features-in-progress)
- [Contact Information](https://github.com/ktrahan2/spacey-bois-backend/tree/main#contact-information)
- [Link to Frontend Repo](https://github.com/ktrahan2/spacey-bois-backend/tree/main#link-to-frontend-repo)

## Description

Spacey Bois backend is made with Go as a very simple api. The api will allow highscores to POST into a PSQL database and also GET these entries.

## Example Code 
```
   func addScores(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	switch r.Method {
	case "POST":
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
	default:
		//method not available error.
		http.Error(w, http.StatusText(405), 405)
	}
} 
```
   
```
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
```

## Technology Used

- Go
- PSQL



## Setting up for the application

To start the server run

``` 
  go run main.go
```

## Main Features

- GET fetch to retrieve all highscores from database
- POST fetch to add highscores to the database

## Features in Progress

- Upgrade backend to create tables through go rather than manually making them in PSQL

## Contact Information

[Kyle Trahan](https://www.linkedin.com/in/kyle-trahan-8384678b/)

## Link to Frontend Repo
https://github.com/ktrahan2/spacey-bois-frontend



