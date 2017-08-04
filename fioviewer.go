// Copyright 2017 Richard Hillmann.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"flag"

	gin "gopkg.in/gin-gonic/gin.v1"

	"github.com/project0/fioviewer/pkg/api"
	"github.com/project0/fioviewer/pkg/fiolog"
)

func main() {

	dir := flag.String("dir", ".", "Path to the fio log files")
	listen := flag.String("listen", ":8080", "Listen to this ip:port")
	staticDir := flag.String("static", ".", "Path to the static (dist) files of the WebUI")

	flag.Parse()

	fiolog.Register(dir)

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// Serve static files for UI components.
	router.StaticFile("/index.html", *staticDir+"/index.html")
	router.StaticFile("/", *staticDir+"/index.html")
	router.Static("/static/", *staticDir+"/static/")

	// add api routes
	api.Register(router)

	router.Run(*listen)
}
