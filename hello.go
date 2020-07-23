package main

import "fmt"
import "os"
import "io/ioutil"
import "encoding/json"
import "process-messages/message"

func main() {
	// holds the data 
	messageTally := make(map[string]int)

	for count, arg := range os.Args {
		if count == 0 { 
			// ignore first arg
			continue
		}
		jsonFile, err := os.Open(arg)
		defer jsonFile.Close()

		// if we os.Open returns an error then handle it
		if (err != nil) {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened %s", arg)
		
		processedJSON := message.MessageJSON{}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		// process json into object
		err = json.Unmarshal(byteValue, &processedJSON)
		if err != nil {
			fmt.Println(err)
		}

		// tally up the number of messages by each sender
		for _, message := range processedJSON.Messages {
			current := messageTally[message.SenderName]
			messageTally[message.SenderName] = current + 1
		}
	}

	for name, count := range messageTally {
		fmt.Println(count, ",", name)
	}
}
