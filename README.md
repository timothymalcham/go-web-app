# go-web-app

Learning Go by building the app from the book "Let's Go"


## Three Essential Aspects of a Go Webapp

1. Handler - think a controller in traditional MVC. Execute app logic and handle HTTP response header/bodies.
2. Router (servemux) - maps url patterns to handlers. Usually one servemux per app.
3. Web Server - listen to incoming requests from the web app itself. No need for a third party like Nginx or Apache.

