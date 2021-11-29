package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// once you give the empty route, every request will be handled by the handle func
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)

		w.Write([]byte("Hello " + name)) // Since we are using ResponseWriter which implements writer interface, we need
		// to give the string as byte array - or with other name, slice of byte

	})
	// we give nil in handler, so it will use default handler
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
