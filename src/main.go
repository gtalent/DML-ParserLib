package main
import(
    "os"
    "flag"
    "fmt"
)

func processInput(input string) string {
	file, err := os.Open(input,  os.O_RDONLY, 0)
	if file == nil {
		fmt.Println("Cannot access file: ", err.String())
		return ""
	}
	file.Close()
	return ""
}

func main() {
	fmt.Println(flag.Arg(0))
}
