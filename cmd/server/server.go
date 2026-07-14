package main

import (
	"client-server/cmd/database"
	"client-server/cmd/models"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CurrencyHandler)
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Print("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
			log.Print("Error: ", err)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error: ", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
			log.Print("Error: ", err)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Error: ", err)
		return
	}

	USDToBRL := models.CurrencyResponse{}
	err = json.Unmarshal(body, &USDToBRL)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = database.Insert(&USDToBRL.USDBRL)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
			log.Print("Timeout saving to database: ", err)
			return
		}

		log.Println("Error saving to database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"bid": strconv.FormatFloat(USDToBRL.USDBRL.Bid, 'f', -1, 64),
	})
}
