package main

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
)

//定义想要注入的结构体
type User struct {
	Name string
	Age  int64
	Male bool
}

func (u *User) GetNum(i int64) int64 {
	return i
}

func (u *User) Print(s string) {
	fmt.Println(s)
}

func (u *User) Say() {
	fmt.Println("hello world")
}

//定义规则
const rule1 = `
rule "name test" "desc"  salience 0
begin
        if 7 == User.GetNum(7){
            User.Age = User.GetNum(89767) + 10000000
			User.Print(User.Name)
			User.Say()
        }
end
`

func StructDemo() {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := context.NewDataContext()
	//注入初始化的结构体
	dataContext.Add("User", user)
	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	err := ruleBuilder.BuildRuleFromString(rule1) //string(bs)
	if err != nil {
		panic(err)
	}
	eng := engine.NewGengine()
	eng.Execute(ruleBuilder, true)

}

func SliceDemo() {
	var aslice []string
	aslice = make([]string, 10)
	aslice[1] = "language"

	dc := context.NewDataContext()
	dc.Add("aslice", aslice)
	f := func(m []string) int {
		return len(m)
	}
	dc.Add("len", f)
	dc.Add("println", fmt.Println)

	fset5 := func(arr []string, value string) {
		arr[4] = value
	}
	dc.Add("Set5", fset5)

	ruleBuilder := builder.NewRuleBuilder(dc)
	err := ruleBuilder.BuildRuleFromString(mrule) //string(bs)
	if err != nil {
		panic(err)
	}
	eng := engine.NewGengine()
	eng.Execute(ruleBuilder, true)
}

var mrule = `
rule "len"   salience 0
begin
	  l=len(aslice)
	  println(l)
	
	  println(aslice)
	  Set(2,aslice,"qqqq")
	  println(aslice[4])
end
`

func main() {
	SliceDemo()
}
