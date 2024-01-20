package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// package level var since the source has many currencies
var currency = "USD"

// then we need a struct to store the information - this is the outermsot json
type Gold struct {
	Prices []Price      `json:"items"` // tag to the exact json info we need
	Client *http.Client // this is for testing later without a network conn - this now
	// holds our custom client from our testing file.
}

// this is for the nested json inside our items array, again we only get the fields we need
type Price struct {
	Currency      string    `json:"currency"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

func (g *Gold) GetPrices() (*Price, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client                                                         // we have our client
	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency) // we have our url

	// now we need a request
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error contacting goldprice.org", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading json")
		return nil, err
	}

	gold := Gold{}
	var previous, current, change float64
	err = json.Unmarshal(body, &gold) // This puts the whole json into our gold struct
	if err != nil {
		log.Println("error converting json - Unmarshaling", err)
		return nil, err
	}
	// then after we have the whole json, we go into the nexted json and pick out what we want
	previous, current, change = gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change

	// then we need to put what we want into our Price struct that we created to hold our info.
	var currentInfo = Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}
