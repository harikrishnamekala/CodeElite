package dboperations

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	//DATABASE Constant Coz Database name is variable
	DATABASE = "problemset"
)

type Input struct {
	Inputid   int
	InputData string
}

type Output struct {
	Outputid   int
	Outputdata string
}

type Testcaseprint struct {
	Problemid        string
	Problemstatement string
	TestcasesInput   []Input
	TestcasesOutput  []Output
}

/*
	To EstablishConnection with the MongoDB Database
*/

func EstablishConnection(DataCatagory string) *mgo.Collection {
	session, err := mgo.Dial("mongodb://testuser:password@localhost")
	if err != nil {
		panic(err)
	}
	c := session.DB(DATABASE).C(DataCatagory)

	var result Testcaseprint

	//datam := "{\"problemid\" : \"print-input-value\"}"

	err = c.Find(bson.M{"problemid": "print-input-value"}).One(&result)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	return c
}

func SelectData(data *mgo.Collection) {

}
