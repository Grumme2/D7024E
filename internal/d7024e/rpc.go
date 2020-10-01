package d7024e

type RPC struct {
	Sender KademliaID
	TargetAddress string
	TargetPort string
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

func NewRPC(sender KademliaID, targetAddress string, targetPort string, content string) RPC {
	return RPC{Sender: sender, TargetAddress: targetAddress, TargetPort: targetPort, Content: content}
}

//Add serialization support
