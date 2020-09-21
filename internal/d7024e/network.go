package d7024e

import (
	"fmt"
	"net"
)

type Network struct {
	routingTable *RoutingTable
}

func Listen(ip string, port int) {
	PORT := ":8000" //Predefined port

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("RECEIVED: ", string(buffer[0:n-1]))

		data := []byte("Alive")
		fmt.Printf("REPONSE: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//func (network *Network) SendPingMessage(contact *Contact) {
func SendPingMessage(contact *Contact) {
	CONNECT := contact.Address + ":8000"

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	for {
		data := []byte("Ping \n")
		_, err = c.Write(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("REPLY: %s\n", string(buffer[0:n]))
		return
	}
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
