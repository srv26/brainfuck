package brainfuck

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type ExecLoggedFunc func(program []byte, input io.Reader, output io.Writer) error

func ProvideLoggedFunc(run ExecLoggedFunc) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// r := strings.NewReader(",[.,]")

		// buf := make([]byte, 1)
		// var err error = nil
		// for {
		// 	output := new(bytes.Buffer)
		// 	n, err1 := r.Read(buf)
		// 	//fmt.Print(string(buf[:n]))
		// 	err = Run(buf[:n], os.Stdin, output)
		// 	got := output.Bytes()
		// 	fmt.Print(got)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	if err1 == io.EOF {
		// 		break
		// 	}
		// }
		// path := flag.String("f", "", "path to brainfuck program to run")
		// flag.Parse()

		// if *path == "" {
		// 	flag.Usage()
		// 	os.Exit(1)
		// }

		program, err1 := ioutil.ReadFile("/home/sagarwal/reconstruct/brainfuck/test.bf")
		//fmt.Println(program)
		if err1 != nil {
			log.Fatalln("Failed to read program:", err1)
		}
		fmt.Println("sourav")
		output := new(bytes.Buffer)
		err := Run(program, os.Stdin, output)
		fmt.Println(string(output.Bytes()))
		return err
		// let the defers cleanup logging resources
	}
}
