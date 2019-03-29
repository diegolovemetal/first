package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Age int
	BirthDay string
	Sex string
	Email string
	Phone string
}
/* map 转Json
 */
func testMap() (ret string, err error){

	myMap := make(map[string]interface{})
	myMap["username"] = "diego"
	myMap["age"] = 21
	myMap["sex"] = "male"

	data ,err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("json marshal failed,err :", err)
		return
	}
	//fmt.Printf("%s\n", string(data))
	ret = string(data)
	return
}

func testInt() {
	age := 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("json.marshal failed,error:",err)
		return
	}
	fmt.Printf("%s\n",data)
}
func testSlice() {
	m := make(map[string]interface{})
	var s []map[string]interface{}
	m["username"] = "diego"
	m["age"] = 21
	m["gender"] = "male"
	m["country"] = "China"
	m["religion"] = "Han"
	m["hobby"] = "Sports"
	s = append(s,m)

	m =make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "female"
	s = append(s, m)
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.marshal failed,err:",err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func testStruct() (ret string, err error){
	user1 := &User{
		UserName: "diego",
		NickName: "阙二",
		Age: 21,
		BirthDay: "1997/08/15",
		Sex: "Male",
		Email: "cqupt1984@gamil.com",
		Phone: "15340526395",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json marshal failed,err:",err)
		return
	}
	ret = string(data)
	return 
}

func test()  {
	data, err := testStruct()
	if err != nil {
		fmt.Println("test struct failed,err:",err)
		return
	}
	var user1 User
	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		fmt.Println("unmarshal failed,err:",err)
		return
	}
	fmt.Println(user1)

}

func test2()  {
	data, err := testMap()
	if err != nil {
		fmt.Println("test map failed,err:",err)
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("unmarshal failed,err:",err)
		return
	}
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m))

}
func main() {
	test2()
	fmt.Println("-----")
}