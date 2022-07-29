package main

import (
	//"bytes"
	"encoding/json"
	"log"
	"os"
)

//Json islemleri
type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	King    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAdress(p *Person, i int) Email {
	return p.Email[i]
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

}

func saveJson(filename string, key interface{}) {
	outfile, err := os.Create(filename)
	checkError(err)
	encoder := json.NewEncoder(outfile)
	err2 := encoder.Encode(&key)
	checkError(err2)
	defer outfile.Close()
}

func main() {

	person := Person{
		ID:        1,
		FirstName: "Sadagat",
		LastName:  "Asgarov",
		Username:  "sada",
		Gender:    "male",
		Name: Name{
			Family:   "Asgarovlar",
			Personal: "ll",
		},
		Email: []Email{
			Email{ID: 1, King: "work", Address: "sada@gmail.com"},
			Email{ID: 2, King: "ev", Address: "sadaev@gmail.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Kodlama"},
		},
	}

	/* res:=GetPerson(&p)
	fmt.Println(res)
	fmt.Println(p) */
	saveJson("aaa.json", person)

}
