package dboperations

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	//DATABASE Constant Coz Database name is Should be specified at the Start of App
	DATABASE = "testdatabase"
	//USERNAME store the username to the Database
	USERNAME = "testuser"
	//PASSWORD has the required passphrase to the Database
	PASSWORD = "password"
	//HOST of the Database wheathere it maybe on a seperate server or a local server
	HOST = "localhost"
)

//ProblemID is a String used as key to fetch the details from the MongoDB
type ProblemID bson.M

//DB interface is the print that connects the Struct DBconn and it's methods to make it a member
type DB interface {
	Connect() error
	CloseConnection(Connection *mgo.Session)
	SelectCollection(Collection string)
	Storeproblemdata(ID ProblemID) error
	GetstoredproblemData() Problem
	GetSession() *mgo.Session
}

//Input is the sub-system of the Entire Problem Struct
type Input struct {
	Inputid   int
	InputData string
}

//Output is similiar to I can use the Input also as a mask to output but the Content is no getting Unmarshalled
type Output struct {
	Outputid   int
	Outputdata string
}

//Problem the BluePrint of the Document Stored in the MongoDB
type Problem struct {
	ProblemID        string
	Problemstatement string
	TestcasesInput   []Input
	TestcasesOutput  []Output
}

//DBconn holds the Database Collection and Session of the MongoDB
type DBconn struct {
	Collection  *mgo.Collection
	Conn        *mgo.Session
	ProblemData Problem
}

//Connect is used to Establish a Connection to the Database
func (P *DBconn) Connect() error {
	Sess, err := mgo.Dial("mongodb://" + USERNAME + ":" + PASSWORD + "@" + HOST)
	if err != nil {
		return err
	}

	P.Conn = Sess

	return nil
}

//CloseConnection Closes the Established MongoDB Connection
func (P *DBconn) CloseConnection(Connection *mgo.Session) {
	Connection.Close()
}

//SelectCollection selects the Collection that is to be extracted
func (P *DBconn) SelectCollection(Collection string) {
	Colle := P.Conn.DB(DATABASE).C(Collection)
	P.Collection = Colle
}

//Storeproblemdata Stores the Extracted Mongo Document in the Object
func (P *DBconn) Storeproblemdata(ID ProblemID) error {

	var Problemdetails Problem

	err := P.Collection.Find(ID).One(&Problemdetails)

	P.ProblemData = Problemdetails

	if err != nil {
		return err
	}
	return nil
}

//GetSession is used to get the Current Session in the MongoDB
func (P *DBconn) GetSession() *mgo.Session {
	return P.Conn
}

//GetstoredproblemData is used to get the Extracted Document in the MongoDB
func (P *DBconn) GetstoredproblemData() Problem {
	return P.ProblemData
}

//Driver : This is just a regualr Driver Function to show how our interface binding works in Golang
func Driver(Argument DB) {
	if err := Argument.Connect(); err != nil {
		panic(err)
	}
	Argument.SelectCollection("problemset")
	ID := ProblemID{"problemid": "print-input-value"}
	Argument.Storeproblemdata(ID)
	problem := Argument.GetstoredproblemData()
	fmt.Println(problem)
	Argument.CloseConnection(Argument.GetSession())
}

/*
	To EstablishConnection with the MongoDB Database
*/

/*func main() {
	var data = new(DBconn)
	Driver(data)
}*/
