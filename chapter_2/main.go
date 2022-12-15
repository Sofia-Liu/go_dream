package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}

//func main() {
//	fmt.Print("Enter a grade:")
//	reader := bufio.NewReader(os.Stdin)
//	input, err := reader.ReadString('\n')
//	/*ReadString treats '\n' as nil, misunderstandly as error,which
//	  make the code shut down and exit.if we don't want to take it
//	  as an error,add "if".
//	*/
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(input)
//}
