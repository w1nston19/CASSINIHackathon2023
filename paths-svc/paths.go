package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const req = "https://api.mapbox.com/directions/v5/mapbox/driving/%f,%f;%f,%f?alternatives=true&continue_straight=false&geometries=geojson&overview=simplified&steps=false&access_token=pk.eyJ1Ijoiem9uZXYiLCJhIjoiY2xmbzRhYjZ6MHJ1MzNzbnI1ZHBkazh5ZSJ9.q6dxvrNYWsN_aPQtclEP3Q"

type Routes struct {
	Paths []Path `json:"routes"`
}

func Paths(from, to Object) {
	endpoint := fmt.Sprintf(req, from.Coords.first, from.Coords.Second,
		to.Coords.first, to.Coords.Second)
	req, _ := http.Get(endpoint)
	bytes, _ := io.ReadAll(req.Body)
	var routes Routes
	err := json.Unmarshal(bytes, &routes)

	for i, _ := range routes.Paths {
		routes.Paths[i].From = from
		routes.Paths[i].To = to
	}
	from.Paths = append(from.Paths, routes.Paths...)
	to.Paths = append(to.Paths, routes.Paths...)

	fmt.Println(err)
	fmt.Println(routes)
}
