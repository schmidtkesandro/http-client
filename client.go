package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Currency struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// type USDBRL struct {
// 	Currency Currency `json:"USDBRL"`
// }

func main() {
	// faz a solicitação da cotação para o servidor
	ctx, cancel := context.WithTimeout(context.Background(), 1300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Decodifique a resposta JSON em uma struct
	//var data map[string]json.RawMessage
	var currency Currency
	if err := json.NewDecoder(resp.Body).Decode(&currency); err != nil {
		log.Fatal(err)
	}

	// err = json.Unmarshal([]byte(resp.Body), &currency)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// parseia o JSON da resposta HTTP para obter o valor da cotação atual

	//usdbrl := data["USDBRL"]
	// Acesse o valor desejado pelo índice da estrutura
	//currency.Bid = data["bid"].(map[string]interface{})["currency"].(map[string]interface{})["bid"].(string)

	//err = json.Unmarshal(usdbrl, &currency)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(usdbrl))
	// fmt.Println("Valor do dólar : ", currency.Bid)
	// valorDoDolar, _ := strconv.ParseFloat(currency.Bid, 64)
	fmt.Println("Valor do dólar f : ", currency.Bid)
	// salva o valor da cotação em um arquivo
	// err = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %.2f", currency.Bid)), 0644)
	err = os.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %.2f", currency.Bid)), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Cotação atual do dólar: ", currency.Bid)
}
