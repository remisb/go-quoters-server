package main

type Quote struct {
	Id int `sql:"id"`
	Quote string `sql:"quote"`
}

func getQuoteRandom() (Quote, error) {
	dbx, err := dbconn.Open()
	if err != nil {
		return Quote{}, err
	}

	defer dbx.Close()
	var quotes []Quote
	err = dbx.Select(&quotes,"SELECT * FROM quote ORDER BY RAND() LIMIT 1;")
	return quotes[0], err
}

func getQuotes() ([]Quote, error) {
	dbx, err := dbconn.Open()
	if err != nil {
		return nil, err
	}

	defer dbx.Close()
	var quotes []Quote
	err = dbx.Select(&quotes,"SELECT * FROM quote")
	return quotes, err
}

func getQuoteById(id int) (Quote, error) {
	dbx, err := dbconn.Open()
	if err != nil {
		return Quote{}, err
	}

	defer dbx.Close()
	selDB, err := dbx.Query("SELECT * FROM quote WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	quoteIem := Quote{}
	for selDB.Next() {
		var id int
		var quote string
		err = selDB.Scan(&id, &quote)
		if err != nil {
			panic(err.Error())
		}
		quoteIem.Id = id
		quoteIem.Quote = quote
	}
	return quoteIem, nil
}
