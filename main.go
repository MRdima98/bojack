package main

import(
    "log"
    "net/http"
    "html/template"
)

type Reference struct {
    Season int
    Episode int
    Quote string
    Character string
    Image string
    Tags []string
}

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.html")
    references := []Reference{ {
        1, 1,
        "I'm not a hero. I'm a high-functioning sociopath. Merry Christmas!",
        "Sherlock Holmes",
        "sherlock.jpg",
        []string{"Sherlock",
        "Christmas"}},
    }
    t.Execute(w, references)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
