# go-eloqua

[![GoDoc](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua?status.svg)](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua)
[![license](https://img.shields.io/github/license/CleverTouch/go-eloqua.svg?maxAge=2592000)](https://github.com/CleverTouch/go-eloqua/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/CleverTouch/go-eloqua.svg?branch=master)](https://travis-ci.org/CleverTouch/go-eloqua)
[![Coverage Status](https://coveralls.io/repos/github/CleverTouch/go-eloqua/badge.svg?branch=master)](https://coveralls.io/github/CleverTouch/go-eloqua?branch=master)

go-eloqua is a golang library for accessing the Eloqua REST APIs.

**The library is currently in rapid development so production use is not advised until the library has matured.**

Better documentation and example code is planned once the library is more complete.

Feel free to create a pull request to contribute to the project or open an issue to request a feature or report a bug.

## Rest API

The library is focused on using the 2.0 REST API endpoints.
The vast majority of Eloqua models have been created as golang structs. 

- [x] Accounts
- [x] Activities
- [x] Contacts
- [x] Contact fields
- [x] Contact lists
- [x] Contact segments
- [x] Content sections
- [x] Emails - *Mostly done, Related objects to be added*
- [x] Email folders
- [x] Email groups
- [x] Email headers
- [x] Email footers
- [x] Forms - *FormSteps only have basic representation*
- [x] Form data
- [x] Images
- [x] Landing pages - *Mostly done, Related objects to be added*
- [x] Microsites
- [x] Option lists
- [x] Users
- [x] Campaigns - *Campaign Elements only have basic representation*
- [x] Custom objects
- [x] Custom object data
- [x] External activities
- [x] External assets
- [x] External asset types
- [x] Visitors

## Bulk API Implementation Status

Support for the bulk API is planned once all the REST endpoints are supported.

## License

This library is distributed under the MIT license. This library is not the works of Oracle and is not an offical Eloqua library so no official Eloqua support is provided for its use.

The API endpoints used, data models & Oracle documentation links are works of Oracle and do not come under this project's licence.

## Attribution

This project has drawn a lot of inspiration, patterns and teachings from the https://github.com/google/go-github project.

The [Eloqua Developer Help Centre](https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/Welcome.htm) created by Oracle has provided a lot of assistance by having clear documentation and examples.
