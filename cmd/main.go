package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dharmjit/paytab-transfer/domain"
	"github.com/dharmjit/paytab-transfer/handler"
	"github.com/dharmjit/paytab-transfer/repository"
	"github.com/dharmjit/paytab-transfer/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "accounts.json", "provide the filename")
	flag.Parse()
	log.Printf("Reading Account File: %s ", fileName)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("ReadFile: %s\n", err)
	}
	var data []domain.Account
	log.Printf("json unmarshal started")
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatalf("Unmarshal: %s\n", err)
	}
	log.Printf("In Memory Account Repository creation started")
	// to hold the accounts map
	accountMap := make(map[uuid.UUID]domain.Account)
	for _, acc := range data {
		accountMap[acc.ID] = acc
	}
	repo := repository.NewAccountRepository(accountMap)
	service := service.NewAccountService(repo)

	// we use gorilla mux for easier syntax for route declaration
	router := mux.NewRouter()
	router.Handle("/api/v1/transfer", handler.TransferFunds(service)).Methods("POST")
	router.Handle("/api/v1/accounts", handler.ListAccounts(service)).Methods("GET")
	router.Handle("/api/v1/accounts/{accountID}", handler.GetAccount(service)).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("HTTP Server is ready to recieve requests")

	<-done
	// Graceful shut down of the server
	log.Print("HTTP Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP Server Shutdown Failed:%+v", err)
	}
	log.Print("HTTP Server Exited Properly")
}
