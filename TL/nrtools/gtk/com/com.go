package com

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"unsafe"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian,tmp)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

////string to整形转换成字节
//func StringToByte(strSrc string) byte {
//	//str := strings.NewReader(strSrc)
//	//bf := bufio.NewReaderSize(str,0)
//	//rs,_ := bf.ReadSlice([]byte("G")[0])
//	fmt.Println("strSrc",strSrc,"rs[0]",strSrc[0])
//	return
//}

// 字符串转成Ascii码
func Chstrtoint(str string) ( int ){
	b := []byte(str)
	tmp := int(b[0])

	return tmp
}


func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// 字符串0-9转成uint8 0-9
func StringTouint8(str string)(b uint8) {
	switch str {

	case "0":
	b = 0
	case "1":
	b = 1
	case "2":
	b = 2
	case "3":
	b =3
	case "4":
	b = 4
	case "5":
	b =5
	case "6":
	b = 6
	case "7":
	b =7
	case "8":
	b =8
	case "9":
	b = 8
	}
	return b
}

// 字符串转成Ascii码
func Chstrtobyte(str string) ( byte ){
	b := []byte(str)
	tmp := int(b[0])

	return InttoByte(tmp)
}

// 字符串转成Ascii码,可见字符
func InttoByte(i int) (b byte ){
	switch i {
	case 33:
		b = '!'
	case 34:
		b= '"'
	case 35:
		b = '#'
	case 36:
		b= '$'
	case 37:
		b = '%'
	case 38:
		b= '&'
	case 39:
		b = ','
	case 40:
		b= '('
	case 41:
		b = ')'
	case 42:
		b= '*'
	case 43:
		b = '+'
	case 44:
		b= ','
	case 45:
		b = '-'
	case 46:
		b= '.'
	case 47:
		b = '/'
	case 48:
		b= '0'
	case 49:
		b = '1'
	case 50:
		b= '2'
	case 51:
		b = '3'
	case 52:
		b= '4'
	case 53:
		b = '5'
	case 54:
		b= '6'
	case 55:
		b = '7'
	case 56:
		b= '8'
	case 57:
		b = '9'
	case 58:
		b= ':'
	case 59:
		b = ';'
	case 60:
		b= '<'
	case 61:
		b = '='
	case 62:
		b= '&'
	case 63:
		b = '?'
	case 64:
		b= '@'
	case 65:
		b = 'A'
	case 66:
		b= 'B'
	case 67:
		b = 'C'
	case 68:
		b= 'D'
	case 69:
		b = 'E'
	case 70:
		b= 'F'
	case 71:
		b = 'G'
	case 72:
		b= 'H'
	case 73:
		b = 'I'
	case 74:
		b= 'J'
	case 75:
		b = 'K'
	case 76:
		b= 'L'
	case 77:
		b = 'M'
	case 78:
		b= 'N'
	case 79:
		b = 'O'
	case 80:
		b= 'P'
	case 81:
		b = 'Q'
	case 82:
		b= 'R'
	case 83:
		b = 'S'
	case 84:
		b= 'T'
	case 85:
		b = 'U'
	case 86:
		b= 'V'
	case 87:
		b = 'W'
	case 88:
		b= 'X'
	case 89:
		b = 'Y'
	case 90:
		b= 'Z'
	case 91:
		b = '['
	case 92:
		b= '/'
	case 93:
		b = ']'
	case 94:
		b= '^'
	case 95:
		b = '_'
	//case 96:
	//	b=  '、'
	case 97:
		b = 'a'
	case 98:
		b= 'b'
	case 99:
		b = 'c'
	case 100:
		b= 'd'
	case 101:
		b = 'e'
	case 102:
		b= 'f'
	case 103:
		b = 'g'
	case 104:
		b= 'h'
	case 105:
		b = 'i'
	case 106:
		b= 'j'
	case 107:
		b = 'k'
	case 108:
		b= 'l'
	case 109:
		b = 'm'
	case 110:
		b= 'n'
	case 111:
		b = 'o'
	case 112:
		b= 'p'
	case 113:
		b = 'q'
	case 114:
		b= 'r'
	case 115:
		b = 's'
	case 116:
		b= 't'
	case 117:
		b = 'u'
	case 118:
		b= 'v'
	case 119:
		b = 'w'
	case 120:
		b= 'x'
	case 121:
		b = 'y'
	case 122:
		b= 'z'
	case 123:
		b = '{'
	case 124:
		b= '|'
	case 125:
		b = '}'
	case 126:
		b= '`'
	//case 127:
	//	b = 'DEL'
	}
	return
}



func StrtoByte(str string) (b byte ){
	//switch str {
	//case "A":
	//	b = 'A'
	//}
	a := Chstrtoint("A")
	InttoByte(a)
	for i := 0; i<32;i++  {
		//if str == "A" + i {
		//	b = 'A' + i
		//}

	}
	return b
}