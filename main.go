package main

import "log"

func main() {
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggingService(service)

	api := NewApiServer(service)

	log.Fatal(api.Start(":3000"))
}
