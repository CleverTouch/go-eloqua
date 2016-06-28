# go-eloqua

[![GoDoc](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua?status.svg)](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua)
[![license](https://img.shields.io/github/license/CleverTouch/go-eloqua.svg?maxAge=2592000)](https://github.com/CleverTouch/go-eloqua/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/CleverTouch/go-eloqua.svg?branch=master)](https://travis-ci.org/CleverTouch/go-eloqua)
[![Coverage Status](https://coveralls.io/repos/github/CleverTouch/go-eloqua/badge.svg?branch=master)](https://coveralls.io/github/CleverTouch/go-eloqua?branch=master)

go-eloqua is a golang library for accessing the Eloqua REST APIs.

## Rest API

The library is focused on using the 2.0 REST API endpoints.
The vast majority of Eloqua models have been created as golang structs.
All the endpoints in the official Eloqua docs, As of writing, are implemented in this library.

There are many other API endpoints that are not in the official documentation. Feel free to create a pull request or open an issue for these but be warned that they may not be very stable.

### Usage

Import the library.

```go
import "github.com/clevertouch/go-eloqua/eloqua"
```

Create a new client as per the example below, passing in your Eloqua base URL, Company Name, User Name & Password. The base URL can be found by logging into your Eloqua instance and copying the resulting URL up to the end of `.com`. This library uses basic authentication to access your Eloqua instance. **Ensure you always use a base URL starting with `https://` to ensure you login details are transferred encrypted**.

```go
client := eloqua.NewClient("https://secure.p01.eloqua.com", "CompanyName", "User.Name", "myPassWord")
```
You can then use this client to access all the services in this Library. Each of these services aligns with the API endpoints listed in the [Eloqua documentation](https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/REST-API.htm).  For example, To get an email with an ID of 5 you'd do the following:

```go
email, resp, err := client.Emails.Get(5)
```

When performing an update or create request all Eloqua-required properties are required by the method. Additional details can be sent by constructing an Eloqua entity and sending it as the last item in the method call. Here's an example of creating a landing page with additional non-required details.

```go
// Constuct a landing page with our additional data
landingPageInput = &eloqua.LandingPage{
	MicrositeId:  5,
	RelativePath: "example-url",
	FolderID: 25
}
// Send the create request with our landing page name & above input
landingPage, resp, err := client.LandingPages.Create("My new page", &landingPageInput)

```

For most listing requests you can pass through some listing options to control search, count & paging. Here's an example of listing out users:

```go
listOptions := eloqua.ListOptions{Count: 5, Search: "name=test*", Page: 2}
users, resp, err := client.Users.List(opts)
```


### Limitations

Listed below are some areas of the REST API that are known to not be fully implemented:

* Form processing steps only have generic struct representation.
* Campaign Elements (Or steps) only have generic representation.
* The dynamic content rules are very basic and all the different rules are not current supported.
* Segment filter rules have not been implemented.

## Bulk API

Support for the bulk API is planned once all the REST endpoints are supported.

## License

This library is distributed under the MIT license. This library is not the works of Oracle and is not an offical Eloqua library so no official Eloqua support is provided for its use.

The API endpoints used, data models & Oracle documentation links are works of Oracle and do not come under this project's licence.

## Attribution

This project has drawn a lot of inspiration, patterns and teachings from the https://github.com/google/go-github project.

The [Eloqua Developer Help Centre](https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/Welcome.htm) created by Oracle has provided a lot of assistance by having clear documentation and examples.
