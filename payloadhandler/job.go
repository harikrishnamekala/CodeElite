package payloadhandler

import (
	"os"
)

//Testcase is a particular test of the object
type Testcase struct {
	ID       string
	Filename string
	TestData string
}

//OrganisePayloadforjob is used to construct problem with test-cases in an organised manner
func OrganisePayloadforjob(stringfiedjsondata string) PayLoadT {
	var req Request
	req.OrganiseLoad(stringfiedjsondata)

	var Payload PayLoad

	Payload.Deteriniloadtype(req)

	var PayloadT PayLoadT

	PayloadT.Payload = Payload

	PayloadT.AttachtestcasesfromDB(Payload)

	return PayloadT
}

//Writetestcase writes the testcases to the Harddrive
func (T *Testcase) Writetestcase(Path string) {

	File, err := os.Create(Path + string(T.Filename) + string(T.ID))
	defer File.Close()
	if err != nil {
		panic(err)
	}
	File.Write([]byte(T.TestData))
}

//EncapsulateIOtotestcases packages the Testcases in Objects
func (T *Testcase) EncapsulateIOtotestcases(Testcasetype string) {

	go func() {
		for Inputv := range IOcase.InputTestCases {

		}
	}()
	go func() {
		for Outputv := range IOcase.OutputTestCases {

		}
	}()
}

//Startexecution Starts the Execution To be Called upon Organising the Data
func (P *PayLoadT) Startexecution() {

}
