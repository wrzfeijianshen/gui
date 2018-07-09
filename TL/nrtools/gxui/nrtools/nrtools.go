package nrtools

import (
	"github.com/tarm/serial"
)

type ECINS struct {
	Name string
	Baud int
}

func OpenSerial(SerialNum string) (*serial.Port, bool) {
	c := &serial.Config{Name: SerialNum, Baud: 4800}
	s, err := serial.OpenPort(c)
	if err != nil {
		//		log.Fatal(err)
		return s, false
	}
	return s, true
}

func WriteSerial(n *serial.Port, str string) (bool, error) {
	_, err := n.Write([]byte(str))
	if err != nil {
		//		log.Fatal(err)
		return false, err
	} else {
		return true, nil
	}
	/*
		buf := make([]byte, 128)
		ss, err = ss.Read(buf)
		if err != nil {
			//		log.Fatal(err)
			return false,err
		}else{
			// log.Printf("%q", buf[:n])
			return true,nil
		}
	*/
}
