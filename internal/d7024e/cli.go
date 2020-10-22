package d7024e

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
)

type cli struct {
	kademlia *Kademlia
}

func NewCli(kademlia *Kademlia) *cli {
	cli := &cli{kademlia}
	return cli
}

func (cli *cli) AwaitCommand(){
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
		case "PRINT":
			if (len(inputSplit) > 1) {
				fmt.Println("Printing test: " + inputSplit[1])
			} else {
				fmt.Println("Error! Invalid arguments!")
			}
		case "PUT":
			if (len(inputSplit) == 2) {
				fileUpload := inputSplit[1]
				fmt.Println(fileUpload)
				//Uploads file
				cli.kademlia.Store(fileUpload)
				//See if file is uploaded
				//sleep maybe? to make sure it has time to upload
				hashedUpload := cli.kademlia.network.MakeHash(fileUpload)
				dataFound, data, node := cli.kademlia.LookupData(hashedUpload)
				_ = data //Prevent data declared and not used compilation error
				_ = node //Prevent data declared and not used compilation error
				if (dataFound){
					fmt.Println("File upload successfully! Hashed upload: " + hashedUpload)
				} else {
					fmt.Println("File upload unsuccessful")
				}

			} else {
				fmt.Println("Error! Invalid arguments!")
			}
		case "GET":
			if (len(inputSplit) == 2) {
				hash := inputSplit[1]
				dataFound, data, node := cli.kademlia.LookupData(hash)
				_ = data //Prevent data declared and not used compilation error
				if (dataFound){
					//Also return which node it was retrieved from
					fmt.Println("File download successfully! Downloaded file: " + data + " from node with address: " + node.Address)
				} else {
					fmt.Println("File download unsuccessful")
				}
			} else {
				fmt.Println("Error! Invalid arguments!")
			}
		case "OK":
			fmt.Println("OK command detected")
		case "WHATISMYIP":
			ip := cli.kademlia.network.GetLocalIP()
			fmt.Println("Your IP is " + ip)
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
			fmt.Println(cli.kademlia.network.GetLocalIP())
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