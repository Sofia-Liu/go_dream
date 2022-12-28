package guess

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// for循环，每次先输入，再判断
//	func main() {
//		randNum := GenerateRandInt(0, 100)
//		for {
//			fmt.Println("Enter your number:")
//			reader := bufio.NewReader(os.Stdin)
//			input_str, err := reader.ReadString('\n')
//			input_str = strings.Trim(input_str, "\n")
//			if err != nil {
//				log.Fatal(err)
//			}
//			//字符串转为 int：strconv.Atoi()    反过来：strconv.Itoa()
//			input, err := strconv.Atoi(input_str)
//			if err != nil {
//				log.Fatal(err)
//			}
//			if input == randNum {
//				fmt.Println("You are right!")
//				break
//			} else if input < randNum {
//				fmt.Println("too small")
//			} else if input > randNum {
//				fmt.Println("too big")
//			} else {
//				fmt.Println("error input, please input an int")
//			}
//		}
//	}

func getnum(randNum_in int) {
	fmt.Println("Enter your number:")
	reader := bufio.NewReader(os.Stdin)
	input_str, err := reader.ReadString('\n')
	input_str = strings.Trim(input_str, "\n")
	if err != nil {
		log.Fatal(err)
	}
	//字符串转为 int：strconv.Atoi()    反过来：strconv.Itoa()
	input, err := strconv.Atoi(input_str)
	if err != nil {
		log.Fatal(err)
	}
	if input == randNum_in {
		fmt.Println("You are right!")
	} else if input < randNum_in {
		fmt.Println("too small")
		getnum(randNum_in)
	} else if input > randNum_in {
		fmt.Println("too big")
		getnum(randNum_in)
	} else {
		fmt.Println("error input, please input an int")
	}
}

func main() {
	randNum := GenerateRandInt(0, 100)
	getnum(randNum)
	//fmt.Println("Enter your number:")
	//reader := bufio.NewReader(os.Stdin)
	//input_str, err := reader.ReadString('\n')
	//input_str = strings.Trim(input_str, "\n")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////字符串转为 int：strconv.Atoi()    反过来：strconv.Itoa()
	//input, err := strconv.Atoi(input_str)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

//import (
//	"fmt"
//	"log"
//	"os"
//)
//func main() {
//	fileInfo, err := os.Stat("my.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(fileInfo.Size())
//}

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
