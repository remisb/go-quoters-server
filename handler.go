package main

import (
	"github.com/remisb/go-quoters-server/web"
	"math/rand"
	"net/http"
	"time"
)

var randSource rand.Source
var randomized *rand.Rand

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
	randomized = rand.New(randSource)
}

func getQuotesListHandler(w http.ResponseWriter, r *http.Request) {
	quotes, err := getQuotes()
	if err != nil {
		web.RespondError(w, r, http.StatusInternalServerError, err)
		return
	}
	web.Respond(w, r, http.StatusOK, quotes)
}

func getRandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	quote, err := getQuoteRandom()
	if err != nil {
		web.RespondError(w, r, http.StatusInternalServerError, err)
		return
	}
	web.Respond(w, r, http.StatusOK, quote)
}

func returnQuoteById(w http.ResponseWriter, r *http.Request, quoteId int) {
	quoteItem, err := getQuoteById(quoteId)
	if err != nil {
		web.RespondError(w, r, http.StatusInternalServerError, err)
		return
	}

	web.Respond(w, r, http.StatusOK, quoteItem)
}

func getQuoteHandler(w http.ResponseWriter, r *http.Request) {
	id := web.UrlParamInt(r, "id", 0)
	if id == 0 {
		web.RespondError(w, r, http.StatusBadRequest, "passed id value is undefined or invalid")
		return
	}

	returnQuoteById(w, r, id)
}
