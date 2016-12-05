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
  		return `<h1 class="yes">Today.</h1>`
  	case today + 1:
  		return `<h1 class="yes">Tomorrow.</h1>`
  	case today + 2:
  		return `<h1 class="maybe">In two days.</h1>`
  	default:
  		return `<h1 class="no">Too far away.</h1>`
    }
}

func header() string {
  return `<html>
            <head>
              <link rel="stylesheet" type="text/css" href="style.css">
              <meta charset="UTF-8">
                <title>When is Friday?</title>
              <meta property="fb:admins" content="3500652">
            </head>`
}

func body() string {
  return "<body>" + eccles() + "</body>"
}

func footer() string {
  return "</html>"
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, header())
    fmt.Fprint(w, body())
    fmt.Fprint(w, footer())
}
