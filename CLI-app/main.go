package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	fmt.Println("Enter the string:")
	i := bufio.NewReader(os.Stdin)
	s, _:= i.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}
