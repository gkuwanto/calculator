package main

import (
	"fmt"
	"net/http"
	"strconv"
	"math"
)

func calculatorHandler(w http.ResponseWriter, req *http.Request) {

	numA, err := strconv.Atoi(req.URL.Query().Get("num_a"))
	numB, err := strconv.Atoi(req.URL.Query().Get("num_b"))
	if err != nil {
		fmt.Fprintf(w,"Invalid number")
		return
	}
	operator := req.URL.Query().Get("operator")
	var result int
	switch operator {
	case "plus":
		result = numA + numB	
	case "minus":
		result = numA - numB
	case "times":
		result = numA * numB
	case "divide":
		result = numA / numB
	case "power":
		resultf := math.Pow(float64(numA),float64(numB))
		fmt.Fprintf(w, "%v\n", resultf)
		return
	default:
		fmt.Fprintf(w, "invalid operator")	
		return
	}


	fmt.Fprintf(w,"%v\n",result) 
}

func main() {

	http.HandleFunc("/calculator", calculatorHandler)

	fmt.Println("serving HTTP...")
	http.ListenAndServe(":8989", nil)
}
