# Inertia.js Demo App - Ping CRM Echo Svelte  
A demo application to illustrate how [Inertia.js][2] works with [Echo (Go)][3] and [Svelte][6].

> This is a port of the [Svelte Ping CRM][5] written in Laravel and Svelte (but with reduced functionality).

An exemplary installation can be found at https://inertia-echo.herokuapp.com/

**This Demo is work in progress and may contain issues, errors and missing functionality!**

## Pre-requisites
[Download and install][7] Golang for your platform.

You need to have a [PostgreSQL][8] Database running.

## Build
Clone this repository, checkout the `demo` branch and build your own version:

```sh
git clone https://github.com/elipZis/inertia-echo.git
cd inertia-echo
git checkout demo
```

### Environment Variables
```sh
cp .env.example .env
```

Replace the `.env` variables with your setup.

### NPM:
```sh
npm install
```

### Assets
```sh
npm run dev
```

### Build and Start the Server
```sh
go build -o inertia-echo-demo elipzis.com/inertia-echo
./inertia-echo-demo
```

**After your first server start all tables will be created. Use the provided `repository/seed.sql` and seed the database with these.** 

## Usage
After you started the server open a browser and navigate to your configured `server:port`. 
If you seeded with the provided SQL-File you may login with
* johndoe@example.com
* secret

## License and Credits
This demo is released under the MIT license by [elipZis][1].

This program uses multiple other libraries. Credits and thanks to all the developers working on these great projects:
* [Inertia.js][2]
* [Echo][3]
* [Svelte Ping CRM][5]
* [Gorm][9]
* [Gorilla Sessions][10]
* [JWT-Go][11]
* [Go-Validator][12]
* [Go-DotEnv][13]

and many more.

## Disclaimer
This is an example app to illustrate the basic usage of different libraries. 
This demo, source, and the whole package comes without a warranty. It may or may not harm your computer. 
It is not a reference for best-practices or security concerns or any other application concept but an exemplary use-case on how to integrate multiple systems together.
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
