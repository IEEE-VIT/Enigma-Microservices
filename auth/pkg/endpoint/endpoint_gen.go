// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "evento_microservices/auth/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SignInEndpoint endpoint.Endpoint
	SignUpEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.AuthService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		SignInEndpoint: MakeSignInEndpoint(s),
		SignUpEndpoint: MakeSignUpEndpoint(s),
	}
	for _, m := range mdw["SignIn"] {
		eps.SignInEndpoint = m(eps.SignInEndpoint)
	}
	for _, m := range mdw["SignUp"] {
		eps.SignUpEndpoint = m(eps.SignUpEndpoint)
	}
	return eps
}
