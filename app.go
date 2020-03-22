package main

import (
//	"fmt"
	"createdata/server"
	"createdata/router"
	"createdata/middleware"
	"createdata/controller"

)

func setup() server.Info {
    return server.Info{
        Hostname: "",
        UseHTTP: true,
        UseHTTPS: false,
        RedirectToHTTPS: false,
        HTTPPort: 80,
        HTTPSPort: 443,
//        CertFile: "tls/server.crt",
//        KeyFile: "tls/server.key"
    }
}

func main() {
    handler := middleware.SetUp(router.Instance())

	controller.Load()
    // Start the HTTP and HTTPS listeners
    server.Run(
        handler,       // HTTP handler
        handler,       // HTTPS handler
        setup(), // Server settings
    )

}