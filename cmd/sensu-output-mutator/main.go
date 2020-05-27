package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {
	var outJSON, okayText, failText string

	reader := bufio.NewReader(os.Stdin)
	inJSON, _ := reader.ReadString('\n')

	if gjson.Get(inJSON, "check.metadata.labels.okay_text").Exists() {
		okayText = gjson.Get(inJSON, "check.metadata.labels.okay_text").String()
	} else {
		okayText = "Service is healthy"
	}

	if gjson.Get(inJSON, "check.metadata.labels.fail_text").Exists() {
		failText = gjson.Get(inJSON, "check.metadata.labels.fail_text").String()
	} else {
		failText = "Service is unhealthy"
	}

	if gjson.Get(inJSON, "check.status").Int() == 0 {
		outJSON, _ = sjson.Set(inJSON, "check.output", okayText)
	} else {
		outJSON, _ = sjson.Set(inJSON, "check.output", failText)
	}

	fmt.Println(outJSON)
}
