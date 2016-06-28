/*
Package eloqua provides a client for accessing the Eloqua API.

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
*/
package eloqua
