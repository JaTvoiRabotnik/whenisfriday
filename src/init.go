package hello

import (
    "fmt"
    "time"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func eccles() string {
	today := time.Now().Weekday()
	switch time.Friday {
  	case today + 0:
  		return "Today."
  	case today + 1:
  		return "Tomorrow."
  	case today + 2:
  		return "In two days."
  	default:
  		return "Too far away."
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "When's Friday? ")
    fmt.Fprint(w, eccles())
}
