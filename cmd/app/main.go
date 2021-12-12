//go:generate swagger generate spec -m -o ../../swagger/swagger.yml

// Package classification Blockchain API.
//
// API-server for blockchain indexes.
//
// Group and index data are taken from a smart contract, which is deployed on Ropsten (ETH test network) using go-ethereum.
// For connecting to a smart contract you have to create your own Infura API key (see INFURA_ENDPOINT env var)
//
//
//     Schemes: http
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//     License: Apache 2.0 https://www.apache.org/licenses/LICENSE-2.0
//     Contact: Eugene Toropov<eugene@go-masters.co>
//
//     Consumes:
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"fmt"
	"github.com/evt/blockchain-api/internal/app/handlers/handlerblock"
	"github.com/evt/blockchain-api/internal/app/handlers/handlerindex"
	"github.com/evt/blockchain-api/internal/app/services/serviceblock"
	"github.com/evt/blockchain-api/internal/app/services/serviceindex"
	"github.com/evt/blockchain-api/internal/pkg/contract"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evt/blockchain-api/config"
	"github.com/evt/blockchain-api/internal/app/handlers/handlergroup"
	"github.com/evt/blockchain-api/internal/app/services/servicegroup"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// config
	cfg, err := config.Get()
	if err != nil {
		return fmt.Errorf("config.Get failed: %w", err)
	}

	configBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("config MarshalIndent failed: %w", err)
	}

	fmt.Println("Configuration:", string(configBytes))

	// clean architecture: handler -> service -> repository

	// init ethereum client
	ethClient, err := ethclient.Dial(cfg.InfuraEndpoint)
	if err != nil {
		return fmt.Errorf("ethclient.Dial failed: %w", err)
	}

	// Bind to an already deployed contract
	contract, err := contract.Bind(cfg.ContractAddress, ethClient)
	if err != nil {
		return fmt.Errorf("contract.New failed: %w", err)
	}

	// service init
	groupService := servicegroup.New(contract)
	indexService := serviceindex.New(contract)
	blockService := serviceblock.New(ethClient)

	// handler init
	groupHandler := handlergroup.New(groupService)
	indexHandler := handlerindex.New(indexService)
	blockHandler := handlerblock.New(blockService)

	app := fiber.New()
	app.Use(logger.New())

	// routes
	v1 := app.Group("/v1")
	v1.Get("/groups", groupHandler.GetAll)
	v1.Get("/groups/:id", groupHandler.Get)
	v1.Get("/indexes/:id", indexHandler.Get)
	v1.Get("/blocks/:id", blockHandler.Get)
	v1.Get("/blocks/:id/header", blockHandler.GetHeader)

	log.Printf("Running HTTP server on %s\n", cfg.HTTPAddr)

	go func() {
		log.Fatal(app.Listen(cfg.HTTPAddr))
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("closing")

	return nil
}
