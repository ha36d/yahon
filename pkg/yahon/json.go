package yahon

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/json/parser"
)

func (Holder) JsonToHcl(ctx context.Context, source string) {
	input, err := os.ReadFile(source)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	ast, err := parser.Parse([]byte(input))
	if err != nil {
		log.Printf("Error: %s", err)
	}
	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		log.Printf("Error: %s", err)
	}
}
