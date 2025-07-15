package job

import "fmt"

func Test(data []byte) error {
	fmt.Println("Test function called with data:", string(data))
	return nil
}

func HelloWorld(data []byte) error {
	fmt.Println("Hello World job executed")
	return nil
}
