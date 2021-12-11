package main

import (
	"encoding/json"
	"fmt"
	"github.com/evt/blockchain-api/internal/app/handlers/blockhandler"
	"github.com/evt/blockchain-api/internal/app/handlers/indexhandler"
	"github.com/evt/blockchain-api/internal/app/services/blockservice"
	"github.com/evt/blockchain-api/internal/app/services/indexservice"
	"github.com/evt/blockchain-api/internal/pkg/contract"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evt/blockchain-api/config"
	"github.com/evt/blockchain-api/internal/app/handlers/grouphandler"
	"github.com/evt/blockchain-api/internal/app/services/groupservice"
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
	groupService := groupservice.New(contract)
	indexService := indexservice.New(contract)
	blockService := blockservice.New(ethClient)

	// handler init
	groupHandler := grouphandler.New(groupService)
	indexHandler := indexhandler.New(indexService)
	blockHandler := blockhandler.New(blockService)

	app := fiber.New()
	app.Use(logger.New())

	// routes
	app.Get("/groups", groupHandler.GetAll)
	app.Get("/groups/:id", groupHandler.Get)
	app.Get("/indexes/:id", indexHandler.Get)
	app.Get("/blocks/:id", blockHandler.Get)

	log.Printf("Running HTTP server on %s\n", cfg.HTTPAddr)

	go func() { _ = app.Listen(cfg.HTTPAddr) }()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("closing")

	return nil
}
