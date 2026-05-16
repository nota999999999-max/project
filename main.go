package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile"`
	Card         int    `json:"card"`
}

var users = make(map[string]User)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", UsersList).Methods("POST")
	r.HandleFunc("/", DroperHandler).Methods("GET")
	r.HandleFunc("/mobile", MobileNumber).Methods("GET")
	r.HandleFunc("/card", Card).Methods("GET")
	r.HandleFunc("/droper", Dropers).Methods("GET")
	fmt.Println("List")
	http.ListenAndServe(":8080", r)
}

func DroperHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DList"))
}

func UsersList(w http.ResponseWriter, r *http.Request) {
	users := [5]User{
		{Name: "Muhammadjon S", MobileNumber: "+992", Card: 9876543212345678},
		{Name: "Parviz H", MobileNumber: "+992", Card: 1234567898765432},
		{Name: "Zarina A", MobileNumber: "+998", Card: 1111222233334444},
		{Name: "Bezhan Sh", MobileNumber: "+992", Card: 5555666677778888},
		{Name: "Guldofarin Kh", MobileNumber: "+7", Card: 0000111122223333},
	}
	fmt.Println(users)
}

type mobileN struct {
	Prefix   string
	Country  int
	MCompany string
	Number   int
}

func MobileNumber(w http.ResponseWriter, r *http.Request) {

	var number Mobile

	user.Mobile

	user := User{

		Country: CHN,
		Mobile:  "+",
	}

	number := MobileNumber(user)

	fmt.Println("Mobile Number:", number)
}

type Card struct {
	CardNumber   int
	Name         string
	ExpiryDate   string
	Currencies   string
	CVV          int
	Transactions []transactions
}

type transactions struct {
	ID       int
	Amount   float64
	Currency string
	Type     string
}

func Card(w http.ResponseWriter, r *http.Request) {

	var visa Card

	trnList := []transactions{
		{
			ID:       001,
			Amount:   150.75,
			Currency: "USD",
			Type:     "in",
		},
		{
			ID:       002,
			Amount:   320.00,
			Currency: "TJS",
			Type:     "in",
		},
		{
			ID:       003,
			Amount:   89.99,
			Currency: "TJS",
			Type:     "out",
		},
	}

	visa.Transactions = trnList

	fmt.Println(visa)

}

func Dropers(w http.ResponseWriter, r *http.Request) {
	users := [2]Droper{
		{Name: "Parviz H", Mobile_number: "+992", Card_number: 1234567898765432},
		{Name: "Bezhan Sh", Mobile_number: "+992", Card_number: 5555666677778888},
	}
	return Dropers
}
