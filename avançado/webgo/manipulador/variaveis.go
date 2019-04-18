package manipulador

import "text/template"

var Modelos = template.Must(template.ParseFiles("html/ola.html"))
