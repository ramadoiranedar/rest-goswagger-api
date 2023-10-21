package rest

import (
	"context"
	"fmt"
	"net/http"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
)

type AuthorizeApi struct {
	rt *rest_goswagger_api.Runtime
}

func NewAuthorizationApi(rt *rest_goswagger_api.Runtime) *AuthorizeApi {
	return &AuthorizeApi{
		rt: rt,
	}
}

func Authorization(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI) {
	api.APIKeyAuthenticator = func(name, in string, ta security.TokenAuthentication) runtime.Authenticator {
		authFunc := security.APIKeyAuthCtx(name, in, func(ctx context.Context, token string) (context.Context, interface{}, error) {
			type principalKey string
			const principalContextKey principalKey = "principal"
			existPrincipal := ctx.Value(principalContextKey)
			if existPrincipal == nil {
				existPrincipal = &models.Principal{}
				ctx = context.WithValue(ctx, principalContextKey, existPrincipal)
			}

			p, err := ta(token)
			if err != nil {
				return ctx, p, err
			}

			authPrincipal := p.(*models.Principal)
			if authPrincipal.Role != "" {
				existPrincipal.(*models.Principal).Role = authPrincipal.Role
			}

			if authPrincipal.Provider != "" {
				existPrincipal.(*models.Principal).Provider = authPrincipal.Provider
			}

			return ctx, existPrincipal, nil
		})

		return HttpAuthenticator(name, func(r *http.Request) (accept bool, principle interface{}, authError error) {
			return authFunc.Authenticate(r)
		})
	}

	api.KeyAuth = func(token string) (*models.Principal, error) {
		if len(token) > 0 && token == rt.Config().GetString("x_app_key") {
			return &models.Principal{}, nil
		}

		return nil, rt.SetError(http.StatusUnauthorized, fmt.Sprintf("invalid key: %v", token))
	}

	api.HasRoleAuth = func(role string) (*models.Principal, error) {
		if role == "" {
			return nil, rt.SetError(http.StatusForbidden, "access denied for this role")
		}

		return &models.Principal{
			Role: role,
		}, nil
	}

	api.IsPeopleAuth = func(people string) (*models.Principal, error) {
		if people == "" {
			return nil, rt.SetError(http.StatusForbidden, "access denied for this people")
		}

		return &models.Principal{
			People: people,
		}, nil
	}

	api.APIAuthorizer = NewAuthorizationApi(rt)
}

func HttpAuthenticator(name string, handler func(*http.Request) (bool, interface{}, error)) runtime.Authenticator {
	return runtime.AuthenticatorFunc(func(params interface{}) (bool, interface{}, error) {
		if request, ok := params.(*http.Request); ok {
			return handler(request)
		}

		if scoped, ok := params.(*security.ScopedAuthRequest); ok {
			accept, p, err := handler(scoped.Request)
			if name == "x-access-role" && err == nil && p != nil {
				accept = false
				auth := p.(*models.Principal)
				for _, role := range scoped.RequiredScopes {
					if role == auth.Role {
						accept = true
						break
					}
				}
			}
			return accept, p, err
		}
		return false, nil, nil
	})
}

func (a *AuthorizeApi) Authorize(r *http.Request, p interface{}) error {
	authPrincipal := p.(*models.Principal)
	route := middleware.MatchedRouteFrom(r)

	if scopes, ok := route.Authenticator.Scopes["isPeople"]; ok && len(scopes) > 0 {
		if route.Params.Get(scopes[0]) != authPrincipal.People {
			return a.rt.SetError(http.StatusForbidden, "access denied for this PEOPLE resource")
		}
	}

	return nil
}
