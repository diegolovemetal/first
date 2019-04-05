package prototype

import "testing"

func TestPrototype(t *testing.T) {
	resume := NewResume()

	resume.setPersonalInfo("diego", "male", "21")
	resume.setWorkExperience("重庆", "睿智融科")
	resume.Display()

	cloneresume := resume.Clone()
	cloneresume.setPersonalInfo("sasa", "female", "20")
	cloneresume.setWorkExperience("南京", "???" )
	cloneresume.Display()
}
