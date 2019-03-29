package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./file.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close()}(file)
	}
	fmt.Println("open file success")

	write1 := bufio.NewWriter(file)
	space1 := write1.Available()
	fmt.Printf("默认有%d字节\n", space1)

	insertByte, err := write1.Write([]byte("默认大小缓冲写入!\n"))
	if err != nil {
		fmt.Printf("write err:%v\n", err)
	}
	fmt.Printf("写入成功，写入%d字节\n", insertByte)

	useSpace1 := write1.Buffered()
	fmt.Printf("缓冲区已经使用%d字节\n",useSpace1)
	//丢弃缓冲区数据
	write1.Reset(file)

	write2 := bufio.NewWriterSize(file, 10000)
	insertByte2 , err := write2.WriteString("this is write string test!\n")
	if err != nil {
		fmt.Printf("write byte err :%v\n", err)
	}
	fmt.Printf("write string success,写入%d字节\n", insertByte2)

	err = write2.WriteByte('B')
	if err != nil {
		fmt.Printf("write byte err:%v\n", err)
	}

	insertByte3, err := write2.WriteRune('操')
	if err != nil {
		fmt.Printf("write rune err:%v\n", err)
	}
	fmt.Printf("write rune success , 写入 %d 字节\n", insertByte3)

	err = write2.Flush()
	if err != nil {
		fmt.Printf("write2 flush error:%v\n", err)
	}
	fmt.Println("write2 flush success")


}
