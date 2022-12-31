package main

import (
	route "lecture/1230/route"
)

func main() {
	router := route.GetRouter()

	router.Run(":8080")
}