package payloadhandler

//OrganisePayloadforjob is used to construct problem with test-cases in an organised manner
func OrganisePayloadforjob(stringfiedjsondata string) PayLoadT {
	var req Request
	req.OrganiseLoad(stringfiedjsondata)

	var Payload PayLoad

	Payload.Deteriniloadtype(req)

	var PayloadT PayLoadT

	PayloadT.AttachtestcasesfromDB(Payload)

	return PayloadT
}

//Startexecution Starts the Execution To be Called upon Organising the Data
func Startexecution(Payload PayLoadT) {

}
