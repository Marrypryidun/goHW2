package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type worker struct {
	Person people 		`json:"person"`
	Experience float64	`json:"experience,omitempty"`
	Skills string		`json:"skills"`
}
type people struct {
	FirstName string 	`json:"firstName"`
	SecondName string 	`json:"secondName"`
	Age int				`json:"age"`
	Sex string			`json:"sex"`
	HigherEducation bool`json:"higherEducation"`
}

func main() {
	//initialize an worker
	person:= worker{
		people{
			"Maria",
			"Pryidun",
			18,
			"woman",
			false},
		0.5,
		"medicine"}

	// 1-create and initialize the profession of a doctor
	doctor:=  struct {
		w worker
		specialization string

	}{person, "surgeon"}

	//change skills of worker
	person.Skills="programming"

	//2-create and initialize the profession of a doctor
	programer:= struct {
		w worker
		languages [] string
		prject [] string
	}{person,
		[]string{"C++","C#","JS","Go"},
		[]string{"Web-app"}}

	//3-create and initialize the profession of a gardener
	person.Skills="gardening"
	gardener:=struct{
		w worker
		tools[] string
	}{person,[]string{"scissors"}}

	person.Skills="сourage"

	//4-create and initialize the profession of a rescuer
	rescuer:=struct{
		w worker
		isSwim bool
	}{person,true}
	person.Skills="blogging"

	//5-create and initialize the profession of a blogger
	blogger:=struct {
		Worker worker 			`json:"worker"`
		SubscribersNumber int	`json:"subscribersNumber"`
	}{person,15000}

	//tags for bloger in json format
	out, err := json.MarshalIndent( blogger,""," ")
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))

	fmt.Println(doctor,"\n",programer,"\n",gardener,"\n",rescuer,"\n",blogger)
}
