# Spacey Bois

Spacey Bois is a text-based adventure game. 

# Table Of Contents 
- [Description](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#description)
- [How It Works](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#how-it-works)
- [Example Code](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#example-code)
- [Technology Used](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#technology-used)
- [Setting up for the Application](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#setting-up-for-the-application)
- [Main Features](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#main-features)
- [Features in Progress](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#features-in-progress)
- [Contact Information](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#contact-information)
- [Link to Backend Repo](https://github.com/ktrahan2/spacey-bois-frontend/blob/main/README.md#link-to-backend-repo)

## Description

Spacey Bois frontend is made with React-Redux and vanilla CSS. The game is based around a role-playing game called Scum and Villainy. It allows the user to create a character with a name and class type. Then the user is prompted with different options in order to advance the storyline. The storyline prompts are mostly generic but there is an example of one custom prompt showing what happens when a user makes an attack roll. At the end the player is given a score which is then add into the database and rendered on the /highscores route on the frontend. 

## How It Works

[Spacey Bois](https://youtu.be/mLFRHayBuT4)

## Example Code 
```
   const prompt8and9AttackResponse = () => {
        if (diceResult === 6) {
            let optionValue1 = "Move the bodies"
            let optionValue2 = "Help Grips"
            if (playerWeapon === "Blaster" || playerWeapon === "Krieger Blaster Pistol") {
                let text = `You fire off a few precision rounds silencing the inspectors before they even move."\n Vapor: 'Wow check out the itchy finger on ${playerName}. I knew there was something about them I liked.'\n Nines: 'We need to get out of here before the station realizes what happened and closes the Jump Gate.\n ${playerName} can you put these bodies back on the Malklaith shuttle or help Grips clean up the Science Bay.'`
                let response = {
                    text: text,
                    optionValues: [optionValue1, optionValue2],
                    nextPrompt: [12, 13]
                }
                return response
            } else if (playerWeapon === "Psyblade") {
                let text = `You reveal your Psyblade and cut through the inspectors before they can even let out a sound."\n Vapor: 'Wow check out the itchy finger on ${playerName}. I knew there was something about them I liked.'\n Nines: 'We need to get out of here before the station realizes what happened and closes the Jump Gate.'\n ${playerName} can you put these bodies back on the Malklaith shuttle or help Grips clean up the Science Bay.'`
                let response = {
                    text: text,
                    optionValues: [optionValue1, optionValue2],
                    nextPrompt: [12, 13]
                }
                return response
            }.....
    }    
```
   
```
   determineHarmTaken = () => {
        let roll = this.roll6SidedDie()
        this.props.diceResult(roll)

        if (roll === 6) {
        } else if (roll === 5 || roll === 4) {
            this.setState((prevState) => {
                return {currentHarm: {
                    ...prevState.currentHarm,  
                    levelOne: prevState.currentHarm.levelOne + 1, 
                }}
            })
        } else {
            this.setState((prevState) => {
                return {currentHarm: {
                    ...prevState.currentHarm,  
                    levelTwo: prevState.currentHarm.levelTwo + 1, 
                }}
            })
        }
        this.setState({
            optionType: ""
        })
    }
```

## Technology Used

- React-Redux
- CSS


## Setting up for the application

To start the server run

``` 
  go run main.go
```

## Main Features

- Get fetch to retrieve all highscores from database
- Post fetch to add highscores to the database

## Features in Progress

- Upgrade backend to create tables through go rather than manually.

## Contact Information

[Kyle Trahan](https://www.linkedin.com/in/kyle-trahan-8384678b/)

## Link to Frontend Repo
  https://github.com/ktrahan2/spacey-bois-frontend



