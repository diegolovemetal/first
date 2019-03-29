package main

import (
	"fmt"
	"sync"
)



var m sync.Map

type UserInfo struct {
	Name string
	Age int
}
func main() {
	//vv, ok := m.LoadOrStore("1","one")
	//fmt.Println(vv, ok)
	//
	//vv, ok = m.Load("1")
	//fmt.Println(vv, ok)
	//
	//vv, ok = m.LoadOrStore("1", "oen")
	//fmt.Println(vv, ok)
	//
	//vv, ok = m.Load("1")
	//fmt.Println(vv, ok) //one true
	//
	//m.Store("1", "oneone")
	//vv, ok = m.Load("1")
	//fmt.Println(vv, ok)
	//
	//m.Store("2", "twotwo")
	//m.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})
	//
	//m.Delete("1")
	//m.Range(func(k, v interface{}) bool {
	//	fmt.Println(k, v)
	//	return true
	//})
	//
	//map1 := make(map[string]UserInfo)
	//var usr1 UserInfo
	//usr1.Name = "diego"
	//usr1.Age = 21
	//map1["usr1"] = usr1
	//
	//var usr2 UserInfo
	//usr2.Name = "sasa"
	//usr2.Age = 20
	//m.Store("map_test", map1)
	//
	//mapValue, _ := m.Load("map_test")
	//
	//for k, v := range mapValue.(interface{}).(map[string]UserInfo) {
	//	fmt.Println(k, v)
	//	fmt.Println("name:", v.Name)
	//}
	//增
	m.Store("1", []string{"hi"})
	m.Store(1, 1111)

	//查
	if val, ok := m.Load(1); ok {
		fmt.Println("查", val)
	}
	//改
	m.Store("1","掉泪老妈")
	if val, ok := m.Load("1"); ok {
		fmt.Println("改", val)
	}
	//loadorstore
	if v, ok := m.LoadOrStore("22","333"); ok {
		fmt.Println("LoadOrStore", v)
	}
	//删除
	m.Delete("1")
	//遍历
	f := func(key, value interface{}) bool{
		fmt.Println("遍历",key ,value)
		return true
	}
	m.Range(f)
}

