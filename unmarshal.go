package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	input_file *string = flag.String("input", "", "use -input=<>")
)

func main() {
	flag.Parse()
	if len(*input_file) == 0 {
		fmt.Println("input file should be set")
		return
	}
	file, err := os.Open(*input_file)
	if err != nil {
		fmt.Println("open input file error ", err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		org, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		tmp := make(map[string]interface{})
		Err := json.Unmarshal([]byte(org), &tmp)
		if Err != nil {
			fmt.Println("Unmarshal error: ", Err)
			continue
		}
		if tmp2, ok := tmp["features"]; ok {
			if v, exist := tmp2.(map[string]interface{})["data.org_token"]; exist {
				if _, ok := v.(string); ok && len(v.(string)) > 0 {
					fmt.Println(v.(string))

				}
			}
			if v, exist := tmp2.(map[string]interface{})["data.phone"]; exist {
				if _, ok := v.(string); ok && len(v.(string)) > 0 {
					fmt.Println(v.(string))

				}
			}
			if v, exist := tmp2.(map[string]interface{})["data.deviceId"]; exist {
				if _, ok := v.(string); ok && len(v.(string)) > 0 {
					fmt.Println(v.(string))
				}
			}

		}
	}
}
