package main

import (
	//"fmt"
	"github.com/Grumme2/D7024E/internal/test"
	"github.com/Grumme2/D7024E/internal/d7024e"
	"os"
)

func main() {
	test.TestPrint()
	d7024e.Listen("hardcoded ip", 1234)

	arguments := os.Args
	pingIP := arguments[1]
	testContact := d7024e.NewContact(d7024e.NewRandomKademliaID(), pingIP)
	d7024e.SendPingMessage(&testContact)
}
