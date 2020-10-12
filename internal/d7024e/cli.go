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
			return //Exits the function
		case "PONG":
			fmt.Println("PONG command detected")
		case "OK":
			fmt.Println("OK command detected")
		case "HELP":
			fmt.Println("There is no one that can help you")
		default:
			fmt.Println("Command not recognised, type HELP for a list of commands")
	}

	fmt.Println("")
	//Await another command
	cli.AwaitCommand()
}

func (cli *cli) CliPrint(print string){
	fmt.Println(print)
}