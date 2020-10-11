package d7024e

import (
	"encoding/json"
	"fmt"
)

type RPC struct {
	Sender Contact
	TargetAddress string
	MessageType string
	Content string
}

//Declares the RPCProcedure type
type RPCProcedureType string

const (
	Ping = RPCProcedureType("PING")
	Store = RPCProcedureType("STORE")
	FindNode = RPCProcedureType("FINDNODE")
	FindValue = RPCProcedureType("FINDVALUE")
	NodeLookup = RPCProcedureType("NODELOOKUP")
	Ok = RPCProcedureType("OK")
)

func NewRPC(sender Contact, targetAddress string, messageType string, content string) RPC {
	return RPC{Sender: sender, TargetAddress: targetAddress, MessageType: messageType, Content: content}
}

func JSONEncode(message RPC) []byte {
	var jsonData []byte
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}
	return jsonData
}

func JSONDecode(message []byte) RPC {
	var jsonData RPC
	err := json.Unmarshal(message, &jsonData)
	if err != nil {
		fmt.Println(err)
	}
	return jsonData
}