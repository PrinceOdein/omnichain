package main

import (
	"fmt"
	"github.com/PrinceOdein/omnichain/api"
)

func main() {
	fmt.Println("Starting MyChain node...")
	api.StartServer()
}