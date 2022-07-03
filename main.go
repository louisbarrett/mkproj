package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Enter project name: ")
	var input string
	fmt.Scanln(&input)

	os.MkdirAll(input, 0700)
	os.MkdirAll(filepath.Join(input, "src"), 0700)
	os.MkdirAll(filepath.Join(input, "bin"), 0700)
	os.MkdirAll(filepath.Join(input, "pkg"), 0700)
        os.MkdirAll(filepath.Join(input, "cmd",input), 0700)

	ioutil.WriteFile(filepath.Join(input, "src", "main.go"), []byte("package main\n\nfunc main() {\n\t//QED\n}\n"), 0644)
	ioutil.WriteFile(filepath.Join(input, "go.mod"), []byte("module "+input+"\n\n//QED"), 0644)
}
