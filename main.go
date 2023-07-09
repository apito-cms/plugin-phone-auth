package main

import (
	"fmt"
	"github.com/apito-cms/buffers/extensions"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

// plugin internal state and implementation
type Authentication struct {
}

// Init system Function Implementation
func (g Authentication) Init(cred *protobuff.ThirdPartyCredential) error {
	return nil
}

// Migration system Function Implementation
func (g Authentication) Migration() error {
	fmt.Println(fmt.Sprintf("Running Migration with"))
	return nil
}

// SchemaRegister system Function Implementation
func (g Authentication) SchemaRegister() (*extensions.ThirdPartyGraphQLSchemas, error) {
	fmt.Println(fmt.Sprintf("Running Schema Register with"))
	queries := graphql.Fields{}
	mutations := graphql.Fields{}

	args := graphql.FieldConfigArgument{
		"secret": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	}
	args["phone"] = &graphql.ArgumentConfig{
		Type: graphql.String,
	}

	queries["loggedInUser"] = &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "LoggedInUser",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"phone": &graphql.Field{
					Type: graphql.String,
				},
				"avatar": &graphql.Field{
					Type: graphql.String,
				},
			},
		}),
		Resolve: g.GetLoggedInUser,
	}

	queries["login"] = &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "UserLogin",
			Fields: graphql.Fields{
				"id_token": &graphql.Field{
					Type: graphql.String,
				},
				"refresh_token": &graphql.Field{
					Type: graphql.String,
				},
			},
		}),
		Args:    args,
		Resolve: g.Login,
	}

	mutations["register"] = &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "UserRegister",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"id_token": &graphql.Field{
					Type: graphql.String,
				},
				"refresh_token": &graphql.Field{
					Type: graphql.String,
				},
			},
		}),
		Args:    args,
		Resolve: g.Register,
	}

	return &extensions.ThirdPartyGraphQLSchemas{
		Queries:   queries,
		Mutations: mutations,
	}, nil
}

// RESTApiRegister system Function Implementation
func (g Authentication) RESTApiRegister() ([]*extensions.ThirdPartyRESTApi, error) {

	fmt.Println(fmt.Sprintf("Running REST Api Register with"))

	return []*extensions.ThirdPartyRESTApi{
		{
			Path:       "/test",
			Method:     "GET",
			Controller: g.ProviderRegister,
		},
	}, nil
}

// system Function Implementation
func (g Authentication) ProviderRegister(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"message": "provider register",
		"code":    200,
	})
}

// Register system Function Implementation
func (g Authentication) LoadConfiguration() (map[string]*protobuff.PluginDetails, error) {
	fmt.Println("Running Load Configuration")
	return map[string]*protobuff.PluginDetails{}, nil
}

// Register system Function Implementation
func (g Authentication) Login(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println(fmt.Sprintf("Running Login"))
	return map[string]interface{}{
		"id_token":      "id token",
		"refresh_token": "refresh token",
	}, nil
}

// Register system Function Implementation
func (g Authentication) Register(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("Running Email Register")
	return nil, nil
}

// ForgetPassword system Function Implementation
func (g Authentication) ForgetPassword() {
	fmt.Println("Running Register")
}

// SendEmail system Function Implementation
func (g Authentication) SendEmail() {
	fmt.Println("Running Register")
}

// SendOTP system Function Implementation
func (g Authentication) SendOTP() {
	fmt.Println("Running Register")
}

// GetRegisterUser system Function Implementation
func (g Authentication) GetRegisterUser() {
	fmt.Println("Running Register")
}

// GetLoggedInUser system Function Implementation
func (g Authentication) GetLoggedInUser(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println(fmt.Sprintf("Running Register"))
	return nil, nil
}

// EmailAuth because plugin Name is email-auth exported
// This exported Name is case-sensitive for the extension to load
var PhoneAuth Authentication
