package main

import (
	"fmt"
	"log"

	"github.com/jcbondsii/aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	targetUser := "jcbondsii"
	err = cfg.SetUser(targetUser)
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	fmt.Printf("Current user set to: %s\n", targetUser)

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	fmt.Println("--- Current Config File Contents ---")
	fmt.Printf("Struct representation: %+v\n", updatedCfg)
	fmt.Printf("Current User Field   : %s\n", updatedCfg.CurrentUserName)
}
