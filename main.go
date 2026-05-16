package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name         string `json:"name"`
	MobileNumber []MobileN
	Card         []Card
}

type MobileN struct {
	Prefix   string
	Country  int
	Number   int
}

type Card struct {
	CardNumber   int
	CardType     string
	Transactions []transactions
}

type transactions struct {
	ID       int
	Amount   float64
	Currency string
	Type     string
}

var users = make(map[string]User)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", UsersList).Methods("POST")
	r.HandleFunc("/", DroperHandler).Methods("GET")
	r.HandleFunc("/mobile", MobileNumber).Methods("GET")
	r.HandleFunc("/card", UCard).Methods("GET")
	http.HandleFunc("/transfer", TransferCriptoToCard)
	r.HandleFunc("/droper", Dropers).Methods("GET")
	fmt.Println("List")
	http.ListenAndServe(":8080", r)
}

func DroperHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DList"))
}

func UsersList(w http.ResponseWriter, r *http.Request) {
	users := [5]User{
		{Name: "Muhammadjon S", MobileNumber: +992000111111, Card: 9876543212345678},
		{Name: "Parviz H", MobileNumber: +992000222222, Card: 1234567898765432},
		{Name: "Zarina A", MobileNumber: +998000333333, Card: 1111222233334444},
		{Name: "Bezhan Sh", MobileNumber: +992000444444, Card: 5555666677778888},
		{Name: "Guldofarin Kh", MobileNumber: +7000555555, Card: 0000111122223333},
	}
	fmt.Println(users)
}

func MobileNumber(w http.ResponseWriter, r *http.Request) {
	var number MobileN
	
	MList := []mobileN{
		{Prefix: "+", Country: 992, Number: 000111111},
		{Prefix: "+", Country: 992, Number: 000222222},
		{Prefix: "+", Country: 998, Number: 000333333},
		{Prefix: "+", Country: 992, Number: 000444444},
		{Prefix: "+", Country: 7, Number: 000555555},
	}
		fmt.Println("Mobile Number:", number)
}

func UCard(w http.ResponseWriter, r *http.Request) {
	var visa Card
	VList := []Card{
		{CardNumber: 1234567898765432, Format: "visa", User},
		{CardNumber: 1111222233334444, Format: "visa", User},
		{CardNumber: 5555666677778888, Format: "visa", User},
		{CardNumber: 0000111122223333, Format: "visa", User},
	}

	VTList := []transactions{
		{ID: 001, Amount: 150.75, Currency: "USD", Type: "in", Format:},
		{ID: 002, Amount: 320.00, Currency: "TJS", Type: "in", Format:},
		{ID: 003, Amount: 89.99, Currency: "TJS", Type: "out", Format:},
	}
	visa.Transactions = VTList

	var km Card
	KMTList := []Card{
		{CardNumber: 9876543212345678, Format: "km", User},
	}

	KMTList := []transactions{
		{ID: 001, Amount: 500.00, Currency: "TJS", Type: "in"},
		{ID: 002, Amount: 35.00, Currency: "TJS", Type: "out"},
		{ID: 003, Amount: 1500.00, Currency: "TJS", Type: "in"}
	}
	km.Transactions = KMTList

	fmt.Println("Card", visa, km)

}

func TransferCriptoToCard(w http.ResponseWriter, r *http.Request) {
	var cripto Coin

	type Coin struct {
	Name       string  //BTC
	Amount     float64  //
	WalletAdress string //bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh
}

Transfer := []transactions{
{ID: 001, Amount: 500.00, Currency: "TJS", Type: "in"}
}

type km struct{
	PrefixBIN int
	KMNumberNumber int
}

kmNumber := []km{
	{6278, KMNumber},
	{5152, KMNumber},
	{5440, KMNumber},
}

//if 6278, 5152, 5440

if km(PrefixBIN) {
		http.Error(w, "Wrong Card Number", http.StatusBadRequest)
		return
	}
}


func Dropers(w http.ResponseWriter, r *http.Request) {

	users := [2]Droper{
		{CardNumber: 1234567898765432},
		{CardNumber: 5555666677778888},
	}
	return Dropers
}