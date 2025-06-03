package main

import (
	"awesomeProject/repl"
	"fmt"
	"os"
	"os/user"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!\n", usr.Username)
	fmt.Println("Feel free to type in commands...")
	repl.Start(os.Stdin, os.Stdout)
}
