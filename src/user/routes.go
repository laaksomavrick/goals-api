package user

import "github.com/laaksomavrick/goals-api/src/core"

// Routes defines the shape of all the routes for the user package
var Routes = core.Routes{
	core.Route{
		Name:         "Create",
		Method:       "POST",
		Pattern:      "/users",
		AuthRequired: false,
		HandlerFunc:  Create,
	},
	core.Route{
		Name:         "Me",
		Method:       "GET",
		Pattern:      "/users/me",
		AuthRequired: true,
		HandlerFunc:  Show,
	},
}
