package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	fmt.Println(`Starting fisheye-api`)
	fmt.Println(`
 ______   __     ______     __  __     ______     __  __     ______
/\  ___\ /\ \   /\  ___\   /\ \_\ \   /\  ___\   /\ \_\ \   /\  ___\
\ \  __\ \ \ \  \ \___  \  \ \  __ \  \ \  __\   \ \____ \  \ \  __\
 \ \_\    \ \_\  \/\_____\  \ \_\ \_\  \ \_____\  \/\_____\  \ \_____\
  \/_/     \/_/   \/_____/   \/_/\/_/   \/_____/   \/_____/   \/_____/`)

	http.ListenAndServe(":8080", r)
}
