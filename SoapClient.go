package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	SendArray(username, password, RecepientNumber, SpecialNumber, MessageBody, IsFlashMessage)
}

func makeRequestToSend(Data string) {

	dataAsByte := []byte(Data)
	response, err := http.Post("https://sms.sunwaysms.com/SMSWS/SOAP.asmx?wsdl", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func GetCredit(username string, password string) {

	_func := "GetCredit"

	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + username + "</UserName><Password>" + password + "</Password></" + _func + "></soap:Body></soap:Envelope>"

	makeRequestToSend(wsReq)
}

func SendArray(username string, password string, RecipientNumber []string, SpecialNumber string, MessageBody string, IsFlashMessage string) {

	_to := "<string>" + strings.Join(RecipientNumber, "</string><string>") + "</string>"

	_func := "SendArray"

	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + username + "</UserName><Password>" + password + "</Password><RecipientNumber>" + _to + "</RecipientNumber><MessageBody>" + MessageBody + "</MessageBody><SpecialNumber>" + SpecialNumber + "</SpecialNumber><IsFlashMessage>" + IsFlashMessage + "</IsFlashMessage></" + _func + "></soap:Body></soap:Envelope>"

	makeRequestToSend(wsReq)
}

func GetMessageStatus(username string, password string, MessageID []string) {

	_func := "GetMessageStatus"

	_to := "<string>" + strings.Join(MessageID, "</string><string>") + "</string>"

	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + username + "</UserName><Password>" + password + "</Password><MessageID>" + _to + "</MessageID></" + _func + "></soap:Body></soap:Envelope>"

	makeRequestToSend(wsReq)

}

func GetMessageStatus1(username string, password string, MessageID string) {

	_func := "GetMessageStatus"

	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + username + "</UserName><Password>" + password + "</Password><MessageID>" + MessageID + "</MessageID></" + _func + "></soap:Body></soap:Envelope>"

	makeRequestToSend(wsReq)

}

func SendArraySchedule(username string, password string, RecepientNumber []string, MessageBody string, SpecialNumber string, Year string, Month string, Day string, Hour string, Minute string, IsFlashMessage string) {

	_func := "SendArraySchedule"

	_to := "<string>" + strings.Join(RecepientNumber, "</string><string>") + "</string>"

	wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><UserName>" + username + "</UserName><Password>" + password + "</Password><RecipientNumber>" + _to + "</RecipientNumber><MessageBody>" + MessageBody + "</MessageBody><SpecialNumber>" + SpecialNumber + "</SpecialNumber><IsFlashMessage>" + IsFlashMessage + "</IsFlashMessage><Year>" + Year + "</Year><Month>" + Month + "</Month><Day>" + Day + "</Day><Hour>" + Hour + "</Hour><Minute>" + Minute + "</Minute></" + _func + "></soap:Body></soap:Envelope>"

	makeRequestToSend(wsReq)

}
