# go-eloqua

[![GoDoc](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua?status.svg)](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua)
[![license](https://img.shields.io/github/license/CleverTouch/go-eloqua.svg?maxAge=2592000)](https://github.com/CleverTouch/go-eloqua/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/CleverTouch/go-eloqua.svg?branch=master)](https://travis-ci.org/CleverTouch/go-eloqua)
[![Coverage Status](https://coveralls.io/repos/github/CleverTouch/go-eloqua/badge.svg?branch=master)](https://coveralls.io/github/CleverTouch/go-eloqua?branch=master)

go-eloqua is a golang library for accessing the Eloqua REST APIs.

Better documentation and example code is planned once the library is more complete.

## Rest API

The library is focused on using the 2.0 REST API endpoints.
The vast majority of Eloqua models have been created as golang structs.
All the endpoints in the official Eloqua docs, As of writing, are implemented in this library.

There are many other API endpoints that are not in the official documentation. Feel free to create a pull request or open an issue for these but be warned that they may not be very stable. 

#### Limitations

Listed below are some areas of the API that are known to not be fully implemented:

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
