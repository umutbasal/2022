package main

import (
	"log"
	"os"
	"path/filepath"
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

	log.Println("config initialized: `" + serverCommand.Host + ":" + serverCommand.Port + "`")

	p, err := os.Getwd()
	dir := filepath.Base(p)

	if err != nil {
		log.Println(err)
	}

	kodata := os.Getenv("KO_DATA_PATH")
	if kodata == "" {
		kodata = "kodata"
	}

	gif := kodata + "/tree.gif"

	if dir == "bin" {
		gif = "../" + gif
	}

	err = serverCommand.Serve([]string{gif})
	if err != nil {
		log.Println(err)
	}
}
