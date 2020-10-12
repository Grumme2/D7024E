package main

import (
	"github.com/Grumme2/D7024E/internal/d7024e"
	//"fmt"
)

func main() {
	cli := d7024e.NewCli()
	cli.AwaitCommand()
}
