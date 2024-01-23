package main

import "fmt"

// pattern for having a struct and ways to configure this.

type Admin struct {
	Name string
}

type ConfigFunc func(admin *Admin)

func newAdmin(funcs ...ConfigFunc) *Admin {
	admin := &Admin{
		Name: "default name",
	}
	for _, modifier := range funcs {
		modifier(admin)
	}

	return admin
}

func withName(name string) ConfigFunc {
	return func(admin *Admin) {
		admin.Name = name
	}
}

func main() {
	defaultAdmin := newAdmin()
	admin := newAdmin(withName("rakesh"))

	fmt.Println(defaultAdmin)
	fmt.Println(admin)
}
