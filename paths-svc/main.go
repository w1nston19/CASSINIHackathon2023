package main

func main() {
	c := Coordinator{}
	c.Initialize()

	c.Coordinates("Sofia, Bulgaria")
	Paths(Object{Coords: Coordinates{23.090646, 42.011108}},
		Object{Coords: Coordinates{22.68902, 42.280996}})
}
