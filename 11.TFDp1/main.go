package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	dataMap map[int]int
)

func init() {
	dataMap = make(map[int]int)
}

func main() {
	http.HandleFunc("/factorial", HandlerFactorial)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	if result, ok := dataMap[num]; ok {
		return result
	}
	return num * factorial(num-1)
}

//func factorial(num int) int {
//	if num <= 1 {
//		return 1
//	}
//	if result, ok := dataMap[num]; ok {
//		return result
//	}
//	ans := 1
//	for i := 1; i <= num; i++ {
//		ans *= i
//	}
//	dataMap[num] = ans
//	return ans
//}

func HandlerFactorial(writer http.ResponseWriter, req *http.Request) {
	//io.WriteString(writer, "test test test")

	num := req.FormValue("num")
	n, err := strconv.Atoi(num)
	if err != nil {
		http.Error(writer, err.Error(), 404)
		return
	}
	io.WriteString(writer, strconv.Itoa(factorial(n)))
}
