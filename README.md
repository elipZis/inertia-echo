# Inertia-Echo - An Echo Server-Side Adapter for Inertia.js 
[![GitHub license](https://img.shields.io/github/license/elipzis/inertia-echo.svg)](https://github.com/elipzis/inertia-echo/blob/master/LICENSE.md) [![GitHub (pre-)release](https://img.shields.io/badge/release-0.1.0-yellow.svg)](https://github.com/elipzis/inertia-echo/releases/tag/0.1.0) [![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.me/elipzis)

Inertia-Echo is an [Echo][3] (Go) server-side adapter for [Inertia.js][2] to build modern monolithic single-page apps. 
Based on [inertia-laravel][4] and [zgabievi's PingCRM Demo][5].

> A demo app can be found in the `demo` branch at https://github.com/elipZis/inertia-echo/tree/demo

## Pre-requisites
[Download and install][7] Golang for your platform.

## Notes
This module serves as middleware to the [Echo][3] server system and has to be registered accordingly.
It is not intended to be used without [Echo][3] and a client-side [Inertia.js][2].

> For usage instructions about Echo please refer to the [official documentation][14]

## Setup
Create a copy of the example environment variables
```sh
cp .env.example .env
```
Setup your own properties accordingly.

In the configured `resources` directory you need to create a `views` folder and a configured `INERTIA_ROOT_VIEW` file.
You may use the example file provided in this repository which requires [Webpack][15] and [Mix][16].

## Usage
Create a new [Echo][3] instance and register the Inertia middleware with it
```golang
e := echo.New()
e.Use(inertia.Middleware(e))
```
Import the module into your project.

The middleware hooks into the [Echo][3] error and template rendering with a dedicated Inertia instance by itself. 
Therefore, to render a client-side [Inertia.js][2] view you can register a route and render a component
```
// Handler
func hello(c echo.Context) error {
    // Status, Component Name, Data to pass on
    return c.Render(http.StatusOK, "Index", map[string]interface{}{})
}

// Route
e.GET("/hello", hello)
```
The internal template renderer of Inertia-Echo checks whether a fresh full base-site has to be returned or only the reduced Inertia response.

> For more examples refer to the `demo` branch at https://github.com/elipZis/inertia-echo/tree/demo

## Configuration
You can leverage several `...WithConfig` functions to configure this module to your needs.

For example, you may create your own `Inertia-Echo` instance via `NewInertia(...)` and pass the instance to the middleware via 
```
e.Use(inertia.MiddlewareWithConfig(inertia.MiddlewareConfig{
    Inertia: MyInertia,
}))
```
By that you enable yourself to use functionality such as `Share(...)` in your own e.g. handlers.

## License and Credits
This module is released under the MIT license by [elipZis][1].

This program uses multiple other libraries. Credits and thanks to all the developers working on these great projects:
* [Inertia.js][2]
* [Echo][3]
* [Svelte Ping CRM][5]

and many more.

## Disclaimer
This source and the whole package comes without a warranty. It may or may not harm your computer. 
It is not a reference for best-practices or security concerns or any other application concept.
Please use with care and not as absolute reference.  
Any damage cannot be related back to the author. 

  [1]: https://elipZis.com
  [2]: https://inertiajs.com/
  [3]: https://echo.labstack.com/
  [4]: https://github.com/inertiajs/inertia-laravel
  [5]: https://github.com/zgabievi/pingcrm-svelte
  [6]: https://svelte.dev/
  [7]: https://golang.org/dl/
  [8]: https://www.postgresql.org/download/
  [9]: https://github.com/go-gorm/gorm/
  [10]: https://github.com/gorilla/sessions
  [11]: https://github.com/dgrijalva/jwt-go
  [12]: https://github.com/go-playground/validator
  [13]: https://github.com/joho/godotenv
  [14]: https://echo.labstack.com/guide
  [15]: https://webpack.js.org/
  [16]: https://laravel.com/docs/8.x/mix
