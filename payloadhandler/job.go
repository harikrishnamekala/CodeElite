package payloadhandler

import (
	"os"
	"sync"
)

//Testcase is a particular test of the object
type Testcase struct {
	ID       string
	Filename string
	TestData string
	Path     string
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

//Writetestcase writes the testcases to the Secondarymemory
func (T *Testcase) Writetestcase() {

	File, err := os.Create(T.Path + T.Filename + T.ID)
	defer File.Close()
	if err != nil {
		panic(err)
	}
	_, err = File.Write([]byte(T.TestData))
	if err != nil {
		panic(err)
	}
}

//EncapsulateIOwritetestcases packages the Testcases in Objects
func (P *PayLoadT) EncapsulateIOwritetestcases(Path string) {
	var waitgroup *sync.WaitGroup
	waitgroup.Add(len(P.InputTestCases) * 2)
	go func() {
		for _, Inputv := range P.InputTestCases {
			var Newtestcase Testcase
			Newtestcase.Filename = "input"
			Newtestcase.ID = string(Inputv.Inputid)
			Newtestcase.Path = Path
			Newtestcase.TestData = Inputv.InputData
			Newtestcase.Writetestcase()
		}
		waitgroup.Done()
	}()
	go func() {
		for _, Outputv := range P.OutputTestCases {
			var Newtestcase Testcase
			Newtestcase.Filename = "output"
			Newtestcase.ID = string(Outputv.Outputid)
			Newtestcase.Path = Path
			Newtestcase.TestData = Outputv.Outputdata
		}
		waitgroup.Done()
	}()
	waitgroup.Wait()
}

//Startexecution Starts the Execution To be Called upon Organising the Data
func (P *PayLoadT) Startexecution() {

}
