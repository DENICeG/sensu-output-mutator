package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var serviceName string
	if len(os.Args) > 1 {
		serviceName = os.Args[1]
	} else {
		serviceName = "Service"
	}

	reader := bufio.NewReader(os.Stdin)
	inJSON, _ := reader.ReadString('\n')

	var event interface{}
	_ = json.Unmarshal([]byte(inJSON), &event)

	//status := event.(map[string]interface{})["spec"].(map[string]interface{})["check"].(map[string]interface{})["status"].(float64)
	status := event.(map[string]interface{})["check"].(map[string]interface{})["status"].(float64)

	var newOutput string
	if status == 0 {
		newOutput = serviceName + " is healthy"
	} else {
		newOutput = serviceName + " is unhealthy"
	}

	//event.(map[string]interface{})["spec"].(map[string]interface{})["check"].(map[string]interface{})["output"] = newOutput
	event.(map[string]interface{})["check"].(map[string]interface{})["output"] = newOutput

	updatedJSON, _ := json.Marshal(event)
	fmt.Printf("%s", updatedJSON)
}
