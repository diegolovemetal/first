package prototype

import "fmt"

type Resume struct {
	name string
	sex string
	age string
	timeArea string
	company string
}

func (r *Resume) setPersonalInfo(name, sex, age string) {
	if r == nil {
		return
	}
	r.name = name
	r.sex = sex
	r.age = age
}

func (r *Resume) setWorkExperience(timeArea, company string) {
	if r == nil {
		return
	}
	r.company = company
	r.timeArea = timeArea
}

func (r *Resume) Display()  {
	if r == nil {
		return
	}
	fmt.Println("个人信息：", r.name, r.sex, r.age)
	fmt.Println("工作经历：", r.timeArea, r.company)

}

func (r *Resume) Clone() *Resume {
	if r == nil {
		return nil
	}
	new_obj := (*r)
	return &new_obj
}

func NewResume() *Resume {
	return &Resume{}
}

