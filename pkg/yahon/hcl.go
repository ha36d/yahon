package yahon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl"
)

func (Holder) HclToJson(ctx context.Context, source string) {
	input, err := os.ReadFile(source)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	var v interface{}
	err = hcl.Unmarshal(input, &v)
	if err != nil {
		log.Printf("unable to parse HCL: %s", err)
	}
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("unable to marshal json: %s", err)
	}

	fmt.Println(string(json))
}
