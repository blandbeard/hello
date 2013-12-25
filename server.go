package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "github.com/codegangsta/martini-contrib/binding"
)

type Registration struct {
  Uuid  string  `form:"uuid" json:"uuid" binding:"required"`
}


func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Layout: "layout",
  }))

  m.Get("/", func(r render.Render) {
    r.HTML(200, "index", "blandbeard")
  })

  m.Post("/registrations", binding.Bind(Registration{}), func(registration Registration, r render.Render) string {
    r.HTML(200, "registrations", registration.Uuid)
    return registration.Uuid
  })

  m.Run()
}