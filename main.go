package main

import (
	"fmt"
	"log"
	"net/http"
)

func ussdCallback(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	sessionId := r.FormValue("sessionId")
	serviceCode := r.FormValue("serviceCode")
	phoneNumber := r.FormValue("phoneNumber")


	_ = fmt.Sprintf("%s,%s,%s", sessionId, serviceCode, phoneNumber)
	response := r.FormValue("response")

	if len(response) == 0{
		w.Write([]byte("What would you want to check \n1. My Account \n2. My Phone Number"))
		return
	}else{
		switch response{
		case "1":
			w.Write([]byte("Choose account information you want to view \n1. Account Number\n2. Account Balance"))
			return
		case "2":
			w.Write([]byte(fmt.Sprintf("END Your Phone Number is %s", phoneNumber)))
			return
		case "1*1":
			w.Write([]byte(fmt.Sprintf("END Your Account Number is %s", phoneNumber)))
			return
		case "1*2":
			w.Write([]byte("END Your Balance is TK 100"))
			return
		default:
			w.Write([]byte("END Invalid input"))
			return
		}
	}
}
func test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	w.Write([]byte("This App Works In Under Production"))
}
func main(){

	fmt.Println("this is a ussd application for testing ....")
	http.HandleFunc("/", ussdCallback)
	http.HandleFunc("/test",test)

	log.Fatal(http.ListenAndServe(":8080",nil))
}
