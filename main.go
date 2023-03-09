package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/scawand/Go/E14_TP2/module"
)

func main() {
	var LesJoueurs [4]module.Joueur

	LesJoueurs[0].Nom = "John Doe"
	LesJoueurs[0].SesDe[0].ID = 0
	LesJoueurs[0].SesDe[0].NbFace = 20

	LesJoueurs[1].Nom = "John Doe jr"
	LesJoueurs[1].SesDe[0].ID, LesJoueurs[1].SesDe[0].NbFace = 1, 20

	LesJoueurs[2].Nom = "Jane Doe"
	LesJoueurs[2].SesDe[0].ID, LesJoueurs[2].SesDe[0].NbFace = 2, 20

	LesJoueurs[3].Nom = "Jane Doe jr"
	LesJoueurs[3].SesDe[0].ID, LesJoueurs[3].SesDe[0].NbFace = 2, 20

	//fmt.Println(LesJoueurs)
	start := time.Now()
	for cle := range LesJoueurs {
		fmt.Println("\n\n" + LesJoueurs[cle].Nom + " lance les dés")
		for i := 0; i < 10; i++ {
			result := LesJoueurs[cle].LancerDe(LesJoueurs[cle].SesDe[0].NbFace) + 1
			fmt.Printf("%d ", result)
			//time.Sleep(1 * time.Second)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTerminer en %v secondes!\n\n", elapsed.Seconds())

	start = time.Now()
	apis := []string{
		"https://www.dejete.com/api?nbde=10&tpde=20",
		"https://www.dejete.com/api?nbde=10&tpde=4",
		"https://www.dejete.com/api?nbde=10&tpde=6",
		"https://www.dejete.com/api?nbde=10&tpde=8",
	}
	for _, api := range apis {
		deClient := http.Client{
			Timeout: time.Second * 2, // Timeout after 2 seconds
		}

		req, err := http.NewRequest(http.MethodGet, api, nil)
		if err != nil {
			println(err)
		}
		res, getErr := deClient.Do(req)
		if getErr != nil {
			println(getErr)
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			println(readErr)
		}

		var apiResultat []module.ApiDe

		jsonErr := json.Unmarshal(body, &apiResultat)
		if jsonErr != nil {
			fmt.Printf("Parse value: %q, error: %s", string(body), jsonErr.Error())
		}
		fmt.Println(apiResultat)

	}

	elapsed = time.Since(start)
	fmt.Printf("\nTerminer en %v secondes!\n\n", elapsed.Seconds())

	start = time.Now()

	ch := make(chan string, 1)

	for _, api := range apis {
		go GetAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTerminer en %v secondes!\n", elapsed.Seconds())

}

func GetAPI(api string, ch chan string) {
	deClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", err)
	}

	res, getErr := deClient.Do(req)
	if getErr != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		println(readErr)
	}

	var apiResultat []module.ApiDe

	jsonErr := json.Unmarshal(body, &apiResultat)
	if jsonErr != nil {
		println("Parse value: %q, error: %s", string(body), jsonErr.Error())
	}
	ch <- fmt.Sprintf(" %d \n", apiResultat)

}
