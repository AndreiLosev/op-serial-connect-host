package main

import (
	"io/ioutil"
	"op-serial-connect-client/errh"
	"op-serial-connect-client/p2p"
	"os"

	"github.com/cheggaaa/pb/v3"
	"go.bug.st/serial"
)

func main() {
	if len(os.Args) != 3 {
		println("недостаточно аргументов")
		return
	}
	path := os.Args[1]
	target := os.Args[2]
	fileThree := p2p.ShowFileTree(path)
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/pts/4", mode)
	defer port.Close()
	errh.Panic(err)
	targetFilesPath := p2p.CreateHostTree(path, target, fileThree)
	bar := pb.StartNew(len(targetFilesPath))
	for i, x := range fileThree {
		file, err := ioutil.ReadFile(x)
		errh.Panic(err)
		err = p2p.SendData(port, []byte(targetFilesPath[i]), "__FILE_PATH__")
		err = p2p.SendData(port, file, "__FILE__")
		errh.Panic(err)
		bar.Increment()
	}
	err = p2p.SendData(port, []byte("__RETURN__"), "__RETURN__")
	errh.Panic(err)
	bar.Finish()
}
