package brainfuck

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ExecuteIncrease(cmd *cobra.Command, args []string) error {
	cmd.Use = "fgf"
	fmt.Println("hey")
	return nil
}

func ExecuteDecrease(cmd *cobra.Command, args []string) error {

	fmt.Println("hi")
	return nil
}
