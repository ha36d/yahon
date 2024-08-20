package yahon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func (Holder) YamlToJson(ctx context.Context, source string) {
	input, err := os.ReadFile(source)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	var v interface{}
	err = yaml.Unmarshal(input, &v)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	fmt.Print(string(json))
}
