package auth

import (
	"fmt"
	"testing"
)

func TestNewToken(t *testing.T) {
	to := &Token{Secret: "I am a hero"}
	str, err := to.GenToken("cloud")
	fmt.Println(str, err)
	str, err = to.GenToken("cloud1")
	fmt.Println(str, err)
}

func TestValidateToken(t *testing.T) {
	to := &Token{Secret: "I am a hero"}
	str, err := to.GenToken("cloud")
	fmt.Println(str, err)

	name, valid := to.ValidateToken(str)
	fmt.Println(name, valid)

	str = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ1.eyJ1c2VybmFtZSI6ImNsb3VkIn0.FeD2yXXGQe290fFO7_FP_XoRubukCiboa2PYID3bAcc"
	name, valid = to.ValidateToken(str)
	fmt.Println(name, valid)
	str = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNsb3VkIn1.FeD2yXXGQe290fFO7_FP_XoRubukCiboa2PYID3bAcc"
	name, valid = to.ValidateToken(str)
	fmt.Println(name, valid)
	str = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNsb3VkIn0.FeD2yXXGQe290fFO7_FP_XoRubukCiboa2PYID3bAc1"
	name, valid = to.ValidateToken(str)
	fmt.Println(name, valid)
}
