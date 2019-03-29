package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

//func main() {
//	file, err := os.Create("./test.xls")
//	if err != nil {
//		panic(err)
//	}
//	defer func(file *os.File) {file.Close()}(file)
//
//	file.WriteString("\xEF\xBB\xBF")
//
//	w := csv.NewWriter(file)
//	w.Write([]string{"编号","姓名","年龄"})
//	for i:=1; i<11; i++ {
//		num := strconv.FormatInt(int64(i), 10)
//		age := strconv.FormatInt(int64(rand.Intn(100)), 10)
//		w.Write([]string{num, "name" + num, age})
//	}
//	w.Flush()
//
//	records := [][]string{}
//	for i := 11; i < 21; i++ {
//		num := strconv.FormatInt(int64(i), 10)
//		age := strconv.FormatInt(int64(rand.Intn(100)), 10)
//		str := []string{num, "name" + num, age}
//		records = append(records, str)
//	}
//	w.WriteAll(records)
//}
func main() {
	file, err := os.Create("./test.xls")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) { file.Close() }(file)

	file.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(file)
	w.Write([]string{"No", "Name", "Age"})
	for i := 1; i < 11; i++ {
		num := strconv.FormatInt(int64(i), 10)
		age := strconv.FormatInt(int64(rand.Intn(100)), 10)
		w.Write([]string{num, "name" + num, age})
	}

	w.Flush()

	records := [][]string{}
	for i := 11; i < 21; i++ {
		num := strconv.FormatInt(int64(i), 10)
		age := strconv.FormatInt(int64(rand.Intn(100)), 10)
		str := []string{num, "name" + num, age}
		records = append(records, str)
	}
	w.WriteAll(records)

}