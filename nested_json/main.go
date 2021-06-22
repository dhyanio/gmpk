package main
import (
	"encoding/json"
	"fmt"
	"strings"
)
const (
	dataObject       = `{"a":{"b":{"c":"d"}}}`
	dataObjectSecond = `{"x":{"y":{"z":"a"}}}`
	dataKey          = "a/b/c"
	dataKeySecond    = "x/y/z"
)
func findValue(data string, key string) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		panic(err)
	}
	t := strings.Replace(key, "/", "", -1)
	var keyList = strings.Split(t, "")
	birds := result
	con := 0
	double := false
	for _, value := range keyList {
		if double == false {
			birds = birds[value].(map[string]interface{})
		}
		if con >= len(keyList)-2 {
			if double == true {
				fmt.Print(birds[value])
				break
			}
			double = true
			continue
		}
		con++
	}
}
func main() {
	// for testing, just change const variable here...
	findValue(dataObjectSecond, dataKeySecond)
}