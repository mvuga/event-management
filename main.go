package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rest-api/database"
	"rest-api/handlers"
	"rest-api/routes"
	"rest-api/utils"
	"rest-api/vault"

	"github.com/gin-gonic/gin"
)

/* const vaultHost = "http://192.168.1.106:8200"
const secretsPath = "rest-api"
const mountPath = "kv"
const secretId = "0c7c5de2-41c8-f595-9288-cefa789eff9f"
const roleId = "0461b7b9-25b6-939f-ad41-26115223e59a" */

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	vaultHost, err := utils.GetEnvironmentVariables("VAULT_HOST")
	if err != nil {
		log.Fatal(err)
	}
	roleId, err := utils.GetEnvironmentVariables("ROLE_ID")
	if err != nil {
		log.Fatal(err)
	}
	secretId, err := utils.GetEnvironmentVariables("SECRET_ID")
	if err != nil {
		log.Fatal(err)
	}
	mountPath, err := utils.GetEnvironmentVariables("MOUNT_PATH")
	if err != nil {
		log.Fatal(err)
	}
	secretsPath, err := utils.GetEnvironmentVariables("SECRETS_PATH")
	if err != nil {
		log.Fatal(err)
	}
	//Fetch DB params
	databaseData, err := vault.GetDBParams(ctx, vaultHost, roleId, secretId, secretsPath, mountPath)
	if err != nil {
		log.Fatalf("Unable to fetch DB params from Vault: %v\n", err)
	}
	//Create DB connections
	dbPool, err := database.CreateConnectionPool(ctx, databaseData.DBUser, databaseData.DBPassword, databaseData.DBHost, databaseData.DBPort, databaseData.DBName)
	if err != nil {
		log.Fatalf("Unable to create database connection pool: %v\n", err)
	}
	defer dbPool.Close()
	// Create gin server
	eventHandler := handlers.NewEventHandler(dbPool)
	server := gin.Default()
	routes.RegisterRoutes(server, eventHandler)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped cleanly")
}
