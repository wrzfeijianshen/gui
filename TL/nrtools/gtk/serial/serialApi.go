package serial

var(
	SerialPort        *Port
	BoolSentSerialFor bool
	strSet1           string
	strSet2           string
	strSet3           string
)

type ECINS struct {
	Name string
	Baud int
}

func OpenSerial(c *Config) (*Port, bool) {
	//c := &Config{Name: SerialNum, Baud: 4800}
	s, err := OpenPort(c)
	if err != nil {
		//		log.Fatal(err)
		return s, false
	}
	return s, true
}

func WriteSerial(n *Port, str string) (bool, error) {
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
