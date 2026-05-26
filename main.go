package main

import (

	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {

	http.HandleFunc("/", DropCardHandler)
	http.HandleFunc("/users", UsersList)
	http.HandleFunc("/mobile", MobileNumber)
	http.HandleFunc("/card", UCard)
	http.HandleFunc("/transaction", TransferCriptoToCard)
	http.HandleFunc("/transfer", TransactionsRecord)
	http.HandleFunc("/dropcard", DropCard)
	fmt.Println("AllCardList 8080")
	http.ListenAndServe(":8080", nil)

/////////////////////////////////////////////
		r := mux.NewRouter()
	r.HandleFunc("/", DropCardHandler).Methods("GET")
	r.HandleFunc("/user", UsersList).Methods("POST")
	r.HandleFunc("/mobile", MobileNumber).Methods("GET")
	r.HandleFunc("/card", UCard).Methods("GET")
	r.HandleFunc("/transactions", TransferCriptoToCard).Methods("GET")
	r.HandleFunc("/transfer", TransactionsRecord).Methods("GET")
	r.HandleFunc("/dropcard", DropCard).Methods("GET")

	fmt.Println("List")
	http.ListenAndServe(":8080", r)
}

//////////
type User struct {
	Name         string `json:"name"`
	MobileNumber []MobileN
	Card         []Card
}

type MobileN struct {
	Prefix   string `json:"prefix"`
	Country  int `json:"country"`
	Number   int `json:"number"`
}

type Card struct {
	CardNumber   string `json:"card"`
	CardFormat     string `json:"type"`
	Transactions []transactions
}

type CardFormat struct {
	Format string `json:"card_format"`
	PrifixBIN string `json:"bin"`
	Number string `json:"card_number"`
}

type Transactions struct {
	ID       int     `json:"id"`            //serial
	Amount   float64 `json:"money_amount"`  //number
	Currency string  `json:"currency"`      //text (TJS, USD)
	Receiver string  `json:"user_receiver"` //text (Card PAN)
	Sender   string  `json:"sender"`        //text (Name)
	Provider string  `json:"sender_host"`   //text (Binance)
}


func DropCardHandler(w http.ResponseWriter, r *http.Request) {
	
	user := User{
		Name: "name",
}
	w.Write([]byte("UList"))
}


func UsersList(w http.ResponseWriter, r *http.Request) {
	users := [5]User{
		{Name: "Muhammadjon S", Prefix: "+", Country: 992, Number: 000111111, CardNumber: "9876543212345678"},
		{Name: "Parviz H", Prefix: "+", Country: 992, Number: 000222222, CardNumber: "1234567898765432"},
		{Name: "Zarina A", Prefix: "+", Country: 998, Number:000333333, CardNumber: "1111222233334444"},
		{Name: "Bezhan Sh", Prefix: "+", Country: 992, Number:000444444, CardNumber: "5555666677778888"},
		{Name: "Guldofarin Kh", Prefix: "+", Country: 7, Number:000555555, CardNumber: "0000111122223333"},
	}
	
	fmt.Println(users)

	}

//////////////////
func MobileNumber(w http.ResponseWriter, r *http.Request) {
	var number MobileN
	
	MList := []MobileN{
		{Prefix: "+", Country: 992, Number: 000111111},
		{Prefix: "+", Country: 992, Number: 000222222},
		{Prefix: "+", Country: 998, Number: 000333333},
		{Prefix: "+", Country: 992, Number: 000444444},
		{Prefix: "+", Country: 7, Number: 000555555},
	}
		fmt.Println("Mobile Number:", number)
}
///////////


func UCard(w http.ResponseWriter, r *http.Request) {

		VList := []VCard{
		{CardNumber: "1234567898765432"},
		{CardNumber: "1111222233334444"},
		{CardNumber: "5555666677778888"},
	}
		
	VTList := []transactions{
		{ID: 1, Amount: 150.75, Currency: "USD"},
		{ID: 2, Amount: 320.00, Currency: "TJS"},
		{ID: 3, Amount: 89.99, Currency: "TJS"},
	}
	visa.Transactions = VTList


	KMTList := []KMCard{
		{CardNumber: "9876543212345678"},
		
		KMTList := []transactions{
		{ID: 4, Amount: 500.00, Currency: "TJS"},
		{ID: 5, Amount: 35.00, Currency: "TJS"},
		{ID: 6, Amount: 1500.00, Currency: "TJS"},
	}
	km.Transactions = KMTList
}

	UPTList := []UPCard{
		{CardNumber: "0000111122223333"},

	UPList :=[]transactions{
		{ID: 7, Amount: 300.00, Currency: "TJS"}
	}
	up.Transactions = KMTList
}

	fmt.Println("Cards", visa, km, up)
}


///////////////////////////

func TransferCriptoToCard(w http.ResponseWriter, r *http.Request) {

	var cripto Coin

	type Coin struct {
	Name       string  //BTC
	Amount     float64
	WalletAdress string //bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh
}


Transfer := []transactions{
{ID: 1, Amount: 500.00, Currency: "TJS"},
{ID: 2, Amount: 320.00, Currency: "TJS"},
}

/////////

func CardFormat(w http.ResponseWriter, r *http.Request) {

	var visa CardFormat
	var km CardFormat
	var up CardFormat

kmNumber := []KMCard{
	{PrefixBIN: 6278, Number},
	{PrefixBIN: 5152, Number},
	{PrefixBIN: 5440, Number},
}
}

/////////////////////////////////////

func TransactionsRecord(w http.ResponseWriter, r *http.Request) {

	AllCardsList := []Transactions{

		{ID: 1, Amount: 150.75, Currency: "USD", Receiver: "1234********5432", Sender: "Van Li", Provider: "Binance"},
		{ID: 2, Amount: 320.00, Currency: "TJS", Receiver: "0000********3333", Sender: "Alijon R.", Provider: "Wallet"},
		{ID: 3, Amount: 89.99, Currency: "TJS", Receiver: "5555********8888", Sender: "John Tomson", Provider: "American Express"},
	}

	w.Write([]byte("Transactions List"))

	json.NewEncoder(w).Encode(AllCardsList)

}

/////////////////////

func TokenizeVisaCard(cardNumber string) string {
	visacard := "****************"

	token := TokenizeVisaCard(visacard)

	fmt.Println("VisaCard:", visacard)
	fmt.Println("Token:", token)
}
////////

if km(6278, 5152, 5440) {
		http.Error(w, "Wrong Card Number", http.StatusBadRequest)
		return
	}


if visa(992) {
		http.Error(w, "Wrong Card Number", http.StatusBadRequest)
		return
	}

//////////////////////////

func DropCard(w http.ResponseWriter, r *http.Request) {

	users := [2]DropCard{
		{CardNumber: 1234567898765432},
		{CardNumber: 5555666677778888},
	}
	return DropCard()
}