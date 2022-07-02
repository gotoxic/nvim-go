package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/neovim/go-client/nvim"
)

var vsp bool

func init() {
  flag.BoolVar(&vsp, "v", false, "vertical split")
  flag.Parse()
}

func main() {
	// Get address from environment variable set by Nvim.
	addr := os.Getenv("NVIM_LISTEN_ADDRESS")
	if addr == "" {
		log.Fatal("NVIM_LISTEN_ADDRESS not set")
	}

	// Dial with default options.
	v, err := nvim.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup on return.
	defer v.Close()

	path, err := filepath.Abs(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}

	cmd := "sp " + path
  if vsp {
    cmd = "v"+cmd
  }
	if err := v.Command(cmd); err != nil {
		log.Fatal(err)
	}

}
