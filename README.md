
Sourcegraph GoDoc Go Report Card Build Status Codecov Join the chat at https://gitter.im/labstack/echo Forum Twitter License

Supported Go versions
As of version 4.0.0, Echo is available as a Go module. Therefore a Go version capable of understanding /vN suffixed imports is required:

1.9.7+
1.10.3+
1.14+
Any of these versions will allow you to import Echo as github.com/labstack/echo/v4 which is the recommended way of using Echo going forward.

For older versions, please use the latest v3 tag.

Feature Overview
Optimized HTTP router which smartly prioritize routes
Build robust and scalable RESTful APIs
Group APIs
Extensible middleware framework
Define middleware at root, group or route level
Data binding for JSON, XML and form payload
Handy functions to send variety of HTTP responses
Centralized HTTP error handling
Template rendering with any template engine
Define your format for the logger
Highly customizable
Automatic TLS via Let’s Encrypt
HTTP/2 support
Benchmarks
Date: 2020/11/11
Source: https://github.com/vishr/web-framework-benchmark
Lower is better!








The benchmarks above were run on an Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz

Guide
Installation
// go get github.com/labstack/echo/{version}
go get github.com/labstack/echo/v4
Example
package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
Help
Forum
Chat
Contribute
Use issues for everything

For a small change, just send a PR.
For bigger changes open an issue for discussion before sending a PR.
PR should have:
Test case
Documentation
Example (If it makes sense)
You can also contribute by:
Reporting issues
Suggesting new features or enhancements
Improve/fix documentation
Credits
Vishal Rana - Author
Nitin Rana - Consultant
Contributors
License
MIT



GORM
The fantastic ORM library for Golang, aims to be developer friendly.

go report card test status Join the chat at https://gitter.im/jinzhu/gorm Open Collective Backer Open Collective Sponsor MIT license Go.Dev reference

Overview
Full-Featured ORM
Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
Hooks (Before/After Create/Save/Update/Delete/Find)
Eager loading with Preload, Joins
Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
Context, Prepared Statement Mode, DryRun Mode
Batch Insert, FindInBatches, Find To Map
SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, NamedArg, Search/Update/Create with SQL Expr
Composite Primary Key
Auto Migrations
Logger
Extendable, flexible plugin API: Database Resolver (Multiple Databases, Read/Write Splitting) / Prometheus…
Every feature comes with tests
Developer Friendly
Getting Started
GORM Guides https://gorm.io
Contributing
You can help to deliver a better GORM, check out things you can do

License
© Jinzhu, 2013~time.Now

Released under the MIT License
