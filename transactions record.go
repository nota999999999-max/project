package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Transactions struct {
	ID       int     `json:"id"`            //serial
	Amount   float64 `json:"money_amount"`  //number
	Currency string  `json:"currency"`      //text (TJS, USD)
	Receiver string  `json:"user_receiver"` //text (Card PAN)
	Sender   string  `json:"sender"`        //text (Name)
	Provider string  `json:"sender_host"`   //text (Binance)
}

func main() {

	http.HandleFunc("/transfer", TransactionsRecord)

	fmt.Println("AllCardList 8080")
	http.ListenAndServe(":8080", nil)

}

func TransactionsRecord(w http.ResponseWriter, r *http.Request) {

	AllCardsList := []Transactions{

		{ID: 1, Amount: 150.75, Currency: "USD", Receiver: "1234********5432", Sender: "Van Li", Provider: "Binance"},
		{ID: 2, Amount: 320.00, Currency: "TJS", Receiver: "0000********3333", Sender: "Alijon R.", Provider: "Wallet"},
		{ID: 3, Amount: 89.99, Currency: "TJS", Receiver: "5555********8888", Sender: "John Tomson", Provider: "American Express"},
	}

	w.Write([]byte("Transactions List"))

	json.NewEncoder(w).Encode(AllCardsList)

}
