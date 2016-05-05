package eloqua

import "testing"

func TestAuthHeader(t *testing.T) {
	c := NewClient("", "TestCompany", "John.Smith", "mysecret")
	expectedString := "Basic VGVzdENvbXBhbnlcSm9obi5TbWl0aDpteXNlY3JldA=="

	if c.authHeader != expectedString {
		t.Errorf("Auth header is not as expected \nExpected: %s \nRecieved: %s", expectedString, c.authHeader)
	}
}
