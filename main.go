package main

import (
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

//go:embed data/*.csv
var content embed.FS

type country struct {
	Name    string `json:"country"`
	Capital string `json:"capital"`
}

func parseFile() [][]string {
	fileContent, err := content.ReadFile("data/country-list.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return [][]string{}
	}

	reader := csv.NewReader(strings.NewReader(string(fileContent)))
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return data
}

func createCountries(data [][]string) []country {
	var dataMutex sync.RWMutex
	dat := data
	var list []country
	for _, row := range dat {
		dataMutex.Lock()
		ctr := country{row[0], row[1]}
		list = append(list, ctr)
		dataMutex.Unlock()
	}
	return list
}

// func getInput() (input string) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	input = scanner.Text()
// 	input = strings.ReplaceAll(input, " ", "")
// 	return
// }

func createRandomQuestion(list []country) country {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	max := len(list) - 1
	number := r1.Intn(max)
	country := list[number]
	return country
}

// func checkAnswer(expec, ans string) {
// 	if strings.EqualFold(expec, ans) {
// 		fmt.Println("*Ding!Ding!Ding!* You are right!")
// 	} else {
// 		fmt.Println("*Buzzz* Sorry. Wrong answer.")
// 		fmt.Printf("It is %s. Better luck next time.\n", expec)
// 		fmt.Printf("Your answer is %s.\n", ans)
// 	}
// }

func main() {
	data := parseFile()
	countriesList := createCountries(data)

	// ans := getInput()
	// checkAnswer(expec, ans)
	mux := http.NewServeMux()
	mux.HandleFunc("/question", func(w http.ResponseWriter, r *http.Request) {
		cont := createRandomQuestion(countriesList)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cont)
	})

	fmt.Println("Starting server on Port 4444.")
	http.ListenAndServe(":4444", mux)
}
