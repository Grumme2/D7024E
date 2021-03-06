package d7024e

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"encoding/json"
)

type cli struct {
	kademlia *Kademlia
}

func NewCli(kademlia *Kademlia) *cli {
	cli := &cli{kademlia}
	return cli
}

func (cli *cli) AwaitCommand() {
	fmt.Println("Insert command:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	inputText := input.Text()

	space := regexp.MustCompile(` `)
	inputSplit := space.Split(inputText, 10)

	switch strings.ToUpper(inputSplit[0]) {
	case "EXIT":
		fmt.Println("EXIT command detected")
		return //Exits the function and terminates the node
	case "PUT":
		if len(inputSplit) == 2 {
			fileUpload := inputSplit[1]
			fmt.Println(fileUpload)
			//Uploads file
			cli.kademlia.Store(fileUpload) //File upload works (well atleast the RPC is sent and received properly)
			//See if file is uploaded

			time.Sleep(300 * time.Millisecond) //Sleep for 0.3s

			hashedUpload := cli.kademlia.network.MakeHash(fileUpload)

			hashedJson, err := json.Marshal(hashedUpload)
			if err != nil {
				fmt.Println(err)
			}

			dataFound, data, node := cli.kademlia.LookupData(string(hashedJson))
			_ = data //Prevent data declared and not used compilation error
			_ = node //Prevent data declared and not used compilation error
			if dataFound {
				fmt.Println("File upload successfully! Hashed upload: ")
				fmt.Println(hashedUpload)
			} else {
				fmt.Println(dataFound)
				fmt.Println(data)
				fmt.Println(node)
				fmt.Println("File upload unsuccessful")
			}

		} else {
			fmt.Println("Error! Invalid arguments!")
		}
	case "GET":
		if len(inputSplit) > 2 {
			inputSplit = space.Split(inputText, 2)
			hash := inputSplit[1]

			hashedJson, err := json.Marshal(hash)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Hashedjson: " + string(hashedJson))

			dataFound, data, node := cli.kademlia.LookupData(string(hashedJson))
			_ = data //Prevent data declared and not used compilation error
			if dataFound {
				//Also return which node it was retrieved from
				fmt.Println("File download successfully! Downloaded file: " + data + " from node with address: " + node.Address)
			} else {
				fmt.Println(dataFound)
				fmt.Println(data)
				fmt.Println(node)
				fmt.Println("File download unsuccessful")
			}
		} else {
			fmt.Println("Error! Invalid arguments!")
		}
	case "OK":
		fmt.Println("OK command detected")
	case "MYIP":
		ip := GetLocalIP()
		fmt.Println("Your IP is " + ip)
	case "PING":
		if len(inputSplit) == 2 {
			target := inputSplit[1]
			pingRPC := NewRPC(cli.kademlia.network.routingTable.me, target, "PING", "")
			cli.kademlia.network.SendMessage(pingRPC)
			fmt.Println("Sent ping to " + target)
		} else {
			fmt.Println("Error! Invalid arguments!")
		}
	case "SELFINBUCKET":
		buckets := cli.kademlia.network.routingTable.buckets
		myID := cli.kademlia.network.routingTable.me.ID
		myIndex := cli.kademlia.network.routingTable.GetBucketIndex(myID)
		isInBucket := buckets[myIndex].IsContactInBucket(cli.kademlia.network.routingTable.me)
		fmt.Println("Do you have yourself in your own buckets?")
		fmt.Println(isInBucket)
	case "MYCONTACT":
		fmt.Println("My ID and Address")
		fmt.Println(cli.kademlia.network.routingTable.me.ID)
		fmt.Println(GetLocalIP())
	case "BUCKETCONTENTS":
		buckets := cli.kademlia.network.routingTable.buckets
		fmt.Println(buckets)
	case "JN": //JOINNETWORK
		cli.kademlia.JoinNetwork()
	case "HELP":
		fmt.Println("Here are all available commands:")
		fmt.Println("HELP - Shows a list of all available commands.")
		fmt.Println("EXIT - Shuts down the node.")
		fmt.Println("PUT <FILE> - Uploads the given file. Outputs the resulting hash if successful.")
		fmt.Println("GET <HASH> - Returns the contents of the unhashed object.")
	default:
		fmt.Println("Command not recognised, type HELP for a list of commands.")
	}

	fmt.Println("")
	//Await another command
	cli.AwaitCommand()
}
