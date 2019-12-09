package main

func main() {

}

type Pet struct {
	name  string
	color string
}

func newPetByName(name string) *Pet {
	return &Pet{
		name: "cat",
	}
}

func newPetByColor(color string) *Pet {
	return &Pet{
		color: "white",
	}
}

type BlackPet struct {
	Pet // 嵌入Pet, 类似于派生
}

// “构造基类”
func newPet(name string) *Pet {
	return &Pet{
		name: name,
	}
}

func newBlackPet(color string) *BlackPet {
	/*return &BlackPet{Pet{
		color: color,
	}}*/
	pet := &BlackPet{}
	pet.color = color
	return pet
}
