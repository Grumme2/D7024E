package d7024e

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
)

type cli struct {

}

func NewCli() *cli {
	cli := &cli{}
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
			//TODO: Terminate node
			return //Exits the function
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
				//Upload the file
				//Output the hash of the object is upload was successful
				//If unsuccessful, output an error message
				fmt.Println("PUT")
			} else {
				fmt.Println("Error! Invalid arguments!")
			}
		case "GET":
			if (len(inputSplit) == 2) {
				hash := inputSplit[1]
				fmt.Println(hash)
				//Output the object returned from the hash if successful
				//If unsuccessful, output an error message
				fmt.Println("GET")
			} else {
				fmt.Println("Error! Invalid arguments!")
			}
		case "OK":
			fmt.Println("OK command detected")
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
