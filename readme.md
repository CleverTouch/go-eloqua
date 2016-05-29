# go-eloqua

[![GoDoc](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua?status.svg)](https://godoc.org/github.com/CleverTouch/go-eloqua/eloqua) [![Build Status](https://travis-ci.org/CleverTouch/go-eloqua.svg?branch=master)](https://travis-ci.org/CleverTouch/go-eloqua) [![Coverage Status](https://coveralls.io/repos/github/CleverTouch/go-eloqua/badge.svg?branch=master)](https://coveralls.io/github/CleverTouch/go-eloqua?branch=master)

go-eloqua is a golang based library for accessing the Eloqua REST API's.

**The library is currently in rapid development so production use is not advised until the library has matured.**

Better documentation and example code is planned once the library is more complete.

Feel free to create a pull request to contribute to the project or open an issue to request a feature or report a bug.

## Rest API Implementation Status

- [x] Accounts
- [ ] Activities
- [x] Contacts
- [x] Contact fields
- [x] Contact lists
- [x] Contact segments
- [x] Content sections - *Mostly done, Related objects to be added*
- [x] Emails - *Mostly done, Related objects to be added*
- [x] Email folders
- [x] Email groups
- [ ] Email headers
- [ ] Email footers
- [ ] Forms
- [ ] Form data
- [ ] Images
- [ ] Landing pages
- [ ] Microsites
- [ ] Option lists
- [x] Users - *Mostly done, Related objects to be added*
- [ ] Campaigns
- [x] Custom objects
- [x] Custom object data
- [ ] External activities
- [ ] External assets
- [ ] External asset types
- [ ] Visitors

## Bulk API Implementation Status

Support for the bulk API is planned once all the REST endpoints are supported.

## License

This library is distributed under the MIT license. 

## Attribution

This project has drawn a lot of inspiration, patterns and teachings from the https://github.com/google/go-github project.

The [Eloqua Developer Help Centre](https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/Welcome.htm) created by Oracle has provided a lot of assistance by having clear documentation and examples.
