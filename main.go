package forum

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()

	m.Get("/hello", func() string {
		return "hello world"
	})

	m.Run()
}
