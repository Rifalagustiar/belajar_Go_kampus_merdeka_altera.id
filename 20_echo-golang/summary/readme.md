Framework Echo adalah salah satu framework web populer di bahasa Go. Berikut adalah ringkasan materi yang perlu diketahui tentang framework Echo:

Routing di Echo dilakukan dengan menggunakan objek echo.Echo. Anda dapat menambahkan rute menggunakan metode GET, POST, PUT, DELETE, dan lain-lain. Berikut adalah contoh kode untuk menambahkan rute GET:

import (
"net/http"
"github.com/labstack/echo/v4"
)

func main() {
e := echo.New()
e.GET("/", func(c echo.Context) error {
return c.String(http.StatusOK, "Hello, World!")
})
e.Start(":8080")
}

Handler adalah fungsi yang digunakan untuk menangani permintaan HTTP. Echo menyediakan berbagai macam jenis handler, termasuk JSON, HTML, dan XML. Berikut adalah contoh kode untuk menambahkan handler JSON:

type User struct {
Name string `json:"name"`
Email string `json:"email"`
}

func main() {
e := echo.New()
e.GET("/users/:id", func(c echo.Context) error {
id := c.Param("id")
user := User{Name: "John", Email: "john@example.com"}
return c.JSON(http.StatusOK, user)
})
e.Start(":8080")
}
Echo mendukung berbagai macam jenis template, termasuk Go template, HTML template, dan Mustache. Berikut adalah contoh kode untuk menggunakan HTML template:

func main() {
e := echo.New()
e.Renderer = &Template{
templates: template.Must(template.ParseGlob("\*.html")),
}
e.GET("/", func(c echo.Context) error {
return c.Render(http.StatusOK, "index.html", map[string]interface{}{
"title": "Homepage",
"body": "Welcome to the homepage!",
})
})
e.Start(":8080")
}

type Template struct {
templates \*template.Template
}

func (t \*Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
return t.templates.ExecuteTemplate(w, name, data)
}
