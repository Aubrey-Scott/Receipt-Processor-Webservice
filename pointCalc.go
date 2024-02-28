package main

import (
	"math"
	"regexp"
	"strconv"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func calcPoints(rec receipt) int {
	var score int = 0

	score += scoreItems(rec)
	score += scoreTotal(rec)
	score += scoreDate(rec)
	score += scoreTime(rec)

	return score
}

// get score from items
func scoreItems(rec receipt) int {
	var score int = 0
	var itemList []string
	var itemPrices []string

	//make a list of item descriptions
	for i := range rec.Items {
		itemList = append(itemList, rec.Items[i].ShortDescription)
	}

	//make a list of item prices
	for i := range rec.Items {
		itemPrices = append(itemPrices, rec.Items[i].Price)
	}

	//Add 5 score for every 2 items on the receipt
	if len(itemPrices) > 1 {
		score += (len(itemPrices) / 2) * 5
	}

	//trims item descriptions and checks if remaining length is evenly divided into 3, adds to score if so
	for i := range itemList {
		var ruleCheck string = clearString(itemList[i])

		if len(ruleCheck)%3 == 0 {
			if s, err := strconv.ParseFloat(itemPrices[i], 64); err == nil {
				score += int(math.Round(s))
			}
		}
	}

	return score
}

// get score for total
func scoreTotal(rec receipt) int {
	var score int = 0

	//Check if total is evenly divided into .25, adds to score if so, additioanlly adds 50 score if total is truncated
	if s, err := strconv.ParseFloat(rec.Total, 64); err == nil {
		var totalFloat float64 = s

		if totalFloat == math.Trunc(totalFloat) {
			score += 50
		}

		if (math.Mod(totalFloat, .25)) == 0 {
			score += 25
		}
	}

	return score
}

// get score for date
func scoreDate(rec receipt) int {
	var score int = 0

	//turn date into int and check if odd, if so, add 6 to score
	var last2 string = rec.PurchaseDate[len(rec.PurchaseDate)-2:]

	i, err := strconv.Atoi(last2)
	if err != nil {
		panic(err)
	}

	if (i % 2) != 0 {
		score += 6
	}

	return score
}

// get score for time
func scoreTime(rec receipt) int {
	var score int = 0

	//turn purchase time into int and check if between 2 and 4, if so, add 10 to score
	var first2 string = rec.PurchaseTime[0:2]

	j, err := strconv.Atoi(first2)
	if err != nil {
		panic(err)
	}

	if j > 14 && j < 16 {
		score += 10
	}

	return score
}

// helper function that clears all non alphanumeric characters from a string for scoring
func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
