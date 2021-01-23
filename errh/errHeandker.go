package errh

import "fmt"

// Panic ...
func Panic(e error) {
	if e != nil {
		fmt.Printf("%v \n", e)
		panic(e)
	}
}

// IsFile ...
func IsFile(e error) error {
	if e != nil {
		if e.Error() == "readdirent: not a directory" {
			return nil
		}
	}
	return e
}
