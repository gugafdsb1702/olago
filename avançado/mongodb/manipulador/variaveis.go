package manipulador

import "text/template"

var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
