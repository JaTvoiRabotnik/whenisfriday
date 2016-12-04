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

func header() string {
  return " <html>
                    <head>
                      <link rel="stylesheet" type="text/css" href="style.css">
                      <meta charset="UTF-8">
                        <title>When is Friday?</title>
                      <meta property="fb:admins" content="3500652">
                    </head>"
}

func body() string {
  return "<body>
            <h1>" + eccles() + "</h1>
          </body>"
}

func footer() string {
  return "</html>"
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, header())
    fmt.Fprint(w, body())
    fmt.Fprint(w, footer())
}
