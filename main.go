package main

import (
	"net/http"
)

func main() {
	fileServer := http.fileServer(http.Dir("./static"))
}
