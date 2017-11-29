package payloadhandler

import (
	"encoding/json"

	"temp.com/dboperations"
)

//Request : Example request subbmitted by the client
type Request struct {
	Problemid  string
	Language   string
	Sourcecode string
	Collection string
}

//PayLoad is work to be done
type PayLoad struct {
	LoadType string
	Data     Request
}

//PayLoadT is work with imported TestCase Dataset
type PayLoadT struct {
	InputTestCases  []dboperations.Input
	OutputTestCases []dboperations.Output
	Payload         PayLoad
}

//OrganiseLoad gives a structure to unstructured data
func (R *Request) OrganiseLoad(Stringfiedjsondata string) {
	if err := json.Unmarshal([]byte(Stringfiedjsondata), R); err != nil {
		panic(err)
	}
}

//Deteriniloadtype has to be initlized weather we want to use the docker or not TODO: Set Env Variable in the OS
func (P *PayLoad) Deteriniloadtype(Req Request) {
	P.Data = Req
	P.LoadType = "non-docker"

	//P.LoadType = string(os.Getenv("EXECUTION_TYPE"))
}

//AttachtestcasesfromDB attaches the Testcases to PayLoadT
func (PT *PayLoadT) AttachtestcasesfromDB(Payload PayLoad) {
	var DB dboperations.DBconn

	ID := dboperations.ProblemID{"problemid": Payload.Data.Problemid}

	PT.InputTestCases, PT.OutputTestCases = dboperations.FetchTestCases(DB, Payload.Data.Collection, ID)

}
