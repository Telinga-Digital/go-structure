package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyJson(v interface{}) string {
	prettyJson, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	return string(prettyJson)
}
