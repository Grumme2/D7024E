package main

import (
	"github.com/Grumme2/D7024E/internal/d7024e"
)

func main() {
	d7024e.SendPingMessage(*d7024e.NewRandomKademliaID(), "localhost", "1337")
}
