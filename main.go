package main

import "os"

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
	serverCommand.Host = "0.0.0.0"
	serverCommand.Port = "8080"

	if port != "" {
		serverCommand.Port = port
	}

	serverCommand.Serve([]string{"tree.gif"})
}
