package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func knockKnock(w http.ResponseWriter, r *http.Request) {
	r.WithContext(context.WithValue(r.Context(), "joke", "Yes")) //key should not be a stringbut I'm ignoring thaton purpose
	fmt.Fprint(w, "Knock Knock - (hint:-type \"/whoIsThere\")")
}

func whoIsThere(w http.ResponseWriter, r *http.Request) {
	testContext(r.Context(), "joke")
	if true {
		fmt.Fprint(w, "A programmer (not a comedian)")
		return
	}
	fmt.Fprint(w, "(hint:-type \"/knockKnock\")")
}

func testContext(ctx context.Context, key interface{}) (bool, interface{}) {
	if val := ctx.Value(key); val != nil {
		fmt.Println(val)
		return true, val
	}
	return false, nil
}

func main() {

	http.HandleFunc("/knockKnock", knockKnock)
	http.HandleFunc("/whoIsThere", whoIsThere)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
