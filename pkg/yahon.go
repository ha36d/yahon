package yahon

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/json/parser"
)

func Yahon(ctx context.Context, source string) {
	input, err := os.ReadFile(source)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	ast, err := parser.Parse([]byte(input))
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
