package delim

import "os"
import "fmt"

func OpenFile(file_name string, mode rune) *os.File {
	var f *os.File

	if file_name == "-" {
		switch mode {
		case 'r':
			f = os.Stdin
		case 'w':
			f = os.Stdout
		}

	} else {
		var err error
		switch mode {
		case 'r':
			f, err = os.Open(file_name)
		case 'w':
			f, err = os.Create(file_name)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return f
}
