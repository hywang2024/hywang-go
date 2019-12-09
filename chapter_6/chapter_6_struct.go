package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {
	people := People{"hywang", 22}
	fmt.Println(people.name, people.age)

	//new
	p := new(People)
	p.name = "suan"
	p.age = 23
	fmt.Println(p.name, p.age)
	//&取地址操作 视为对该类型进行一次 new 的实例化操作
	p1 := &People{"hywang", 22}
	fmt.Println(p1.name, p1.age)

	var version int = 1
	command := &Command{}
	command.Comment = "version"
	command.Var = &version
	command.Comment = "show version"
	fmt.Println(command)
	i := newCommand("version", &version, "show version")
	fmt.Println(i)

	i2 := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}
	printMsgType(i2)
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{name, varref, comment}
}

func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}
