package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type adventure struct {
	Description string
	Outdoor     bool
	Duration    int
	Distance    int
}
type preference struct {
	Outdoor     bool `json:"outdoor"`
	MaxDuration int  `json:"maxDuration"`
	MaxDistance int  `json:"maxDistance"`
}

func main() {
	adventureList := populateList()
	if isRand() {
		fmt.Print(randAdventure(adventureList))
	} else {
		fmt.Print(findAdventure(adventureList))
	}
}

func randAdventure(adventureList []adventure) string{
	rand.Seed(time.Now().UnixNano())
	return adventureList[rand.Intn(len(adventureList))].Description
}

func isRand() bool{
	var rand string
	fmt.Println("Pick at random? (y/n)")
	fmt.Scanln(&rand)
	switch rand {
	case "t", "true", "yes", "y":
		return true
	case "f", "false", "no", "n":
		return false
	default:
		fmt.Println("I have no idea what that meant... so I'm going to say you want something random...")
		return true
	}
}

func findAdventure(adventureList []adventure) []string {
	f := createBinding()
	var finalAdventure []string

	for _, adventure := range adventureList {
		if adventure.Duration <= f.MaxDuration && adventure.Distance <= f.MaxDistance && adventure.Outdoor == f.Outdoor {
			finalAdventure = append(finalAdventure, adventure.Description)
		}
	}
	if len(finalAdventure) == 0 {
		finalAdventure = append(finalAdventure, "No adventures meet that criteria")
	}
	return finalAdventure
}

func populateList() []adventure {
	finalAdventureList := []adventure{
		{"Crowder's Mtn", true, 5, 3},
		{"Pan's Labyrinth", false, 3, 0},
		{"V for Vendetta", false, 3, 0},
		{"Ceviche", false, 1, 0},
		{"Imurj", false, 2, 0},
		{"Lake Jordan", true, 1, 0},
		{"Zoo", true, 4, 1},
		{"Chess", false, 1, 0},
		{"Arboretum", true, 2, 0},
		{"Neomonde", false, 1, 0},
		{"Camping", true, 15, 0},
		{"Hmart", false, 1, 0},
		{"Quickly", false, 1, 0},
		{"Michael's Cool Things", false, 3, 0},
		{"Chalk things", true, 1, 0},
		{"Grandfather Mtn", true, 5, 3},
		{"Fantasia 2000", false, 3, 0},
		{"Improve List Features", false, 1, 0},

	}
	return finalAdventureList
}

func createBinding() preference {
	finalPref := preference{}
	var outdoor, duration, distance string
	fmt.Println("Do you want to go outdoors? (y/n)")
	fmt.Scanln(&outdoor)
	fmt.Println("Max duration? (hours)")
	fmt.Scanln(&duration)
	fmt.Println("Max distance? (hours)")
	fmt.Scanln(&distance)
	finalPref.MaxDuration, _ = strconv.Atoi(string(duration))
	finalPref.MaxDistance, _ = strconv.Atoi(string(distance))
	outdoor = strings.ToLower(outdoor)
	outdoor = strings.TrimSpace(outdoor)

	switch outdoor {
	case "t", "true", "yes", "y":
		finalPref.Outdoor = true
	case "f", "false", "no", "n":
		finalPref.Outdoor = false
	default:
		fmt.Println("I have no idea what that meant... so I'm going to say outdoor is true...")
		finalPref.Outdoor = true
	}
	return finalPref
}
