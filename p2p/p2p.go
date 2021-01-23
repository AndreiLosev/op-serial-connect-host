package p2p

import (
	"errors"
	"hash/crc32"
	"op-serial-connect-client/errh"

	"go.bug.st/serial"
)

//SendData ...
func SendData(port serial.Port, data []byte, metaData string) error {
	sendData := createPackages(data, metaData)
	buff := make([]byte, 64)
	for _, pack := range sendData {
		_, err := port.Write(pack)
		errh.Panic(err)
		n, err := port.Read(buff)
		errh.Panic(err)
		if string(buff[:n]) == "___invalid data___" {
			return errors.New("___invalid data___, sendData func")
		}
	}
	return nil
}

func createPackages(data []byte, metaData string) [][]byte {
	buff := createPack(make([][]byte, 0), []byte(metaData))
	const delta = 1000
	dataLen := len(data)
	for i := 0; i < dataLen; i += delta {
		if dataLen <= i+delta {
			buff = createPack(buff, data[i:dataLen])
			buff = createPack(buff, []byte("eOFF"))
		} else {
			buff = createPack(buff, data[i:i+delta])
		}
	}
	return buff
}

func createPack(buff [][]byte, data []byte) [][]byte {
	crc := crc32.ChecksumIEEE(data)
	pack := append([]byte{}, data...)
	pack = append(pack, i32tob(crc)...)
	return append(buff, pack)
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}
