package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:6767")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	connectedClients := 0

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("client %v connected from %s\n", connectedClients, conn.RemoteAddr().String())
		connectedClients++

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	// buf := make([]byte, 4096)
	// messageBuffer := []byte{}
	messageId := 0

	type PrimeCheckRequest struct {
		Method *string  `json:"method"`
		Number *float64 `json:"number"`
	}

	type PrimeCheckResponse struct {
		Method string `json:"method"`
		Prime  bool   `json:"prime"`
	}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			errMessage := err.Error()
			fmt.Printf("%v | %v | %s \n", conn.RemoteAddr().String(), messageId, errMessage)

			conn.Write([]byte(errMessage))
			return
		}

		request := scanner.Text()
		fmt.Printf("%v | %v | processing: %v\n", conn.RemoteAddr().String(), messageId, request)

		data := PrimeCheckRequest{}
		err := json.Unmarshal([]byte(request), &data)
		if err != nil || data.Number == nil || data.Method == nil || *data.Method != "isPrime" {
			errMessage := "malformed JSON for request"
			fmt.Printf("%v | %v | %s \n", conn.RemoteAddr().String(), messageId, errMessage)

			conn.Write([]byte(errMessage))
			return
		}

		isPrime := IsPrime(*data.Number)
		res := PrimeCheckResponse{
			Method: *data.Method,
			Prime:  isPrime,
		}

		reponseBuf, err := json.Marshal(res)
		if err != nil {
			return
		}

		conn.Write(append(reponseBuf, byte('\n')))
		fmt.Printf("%v | %v | %s\n", conn.RemoteAddr().String(), messageId, reponseBuf)

		messageId++
		// messageBuffer = []byte{}
		// buf = nil
		// buf = make([]byte, 4096)
	}
}

func IsPrime(number float64) (isPrime bool) {
	bigFloat := big.NewFloat(number)
	bigInt, accuracy := bigFloat.Int(nil)
	if accuracy == big.Exact {
		isPrime = bigInt.ProbablyPrime(20)
	}
	return
}
