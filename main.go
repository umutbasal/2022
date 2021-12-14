package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	serverCommand := ServerCommand{}
	serverCommand.Ratio = 1.0
	serverCommand.FixedWidth = -1
	serverCommand.FixedHeight = -1
	serverCommand.StretchedScreen = false
	serverCommand.Colored = true
	serverCommand.Reversed = false
	serverCommand.FitScreen = true
	serverCommand.Delay = 0.15
	serverCommand.Host = ""
	serverCommand.Port = "8080"

	if port != "" {
		serverCommand.Port = port
	}

	log.Println("config initialized: " + serverCommand.Host + ":" + serverCommand.Port)

	err := serverCommand.Serve([]string{"tree.gif"})
	if err != nil {
		log.Fatal(err)
	}
}
