package rest

import (
	"context"
	"fmt"
	"net/http"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"github.com/ramadoiranedar/rest-goswagger-api/gen/restapi/operations"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/jwt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"

	openapi_errors "github.com/go-openapi/errors"
)

type AuthorizeApi struct {
	rt *rest_goswagger_api.Runtime
}

func NewAuthorizationApi(rt *rest_goswagger_api.Runtime) *AuthorizeApi {
	return &AuthorizeApi{
		rt: rt,
	}
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

func hasRequiredRole(principal interface{}, requiredRoles []string) bool {
	auth, isPrincipal := principal.(*models.Principal)
	if !isPrincipal {
		return false
	}

	for _, role := range requiredRoles {
		if role == auth.Role {
			return true
		}
	}

	return false
}

func HttpAuthenticator(name string, handler func(*http.Request) (bool, interface{}, error)) runtime.Authenticator {
	return runtime.AuthenticatorFunc(func(params interface{}) (bool, interface{}, error) {
		request, isRequest := params.(*http.Request)
		scopedRequest, isScopedRequest := params.(*security.ScopedAuthRequest)

		if isRequest {
			return handler(request)
		}

		if isScopedRequest {
			accept, principal, err := handler(scopedRequest.Request)

			if name == "x-access-role" && err == nil && principal != nil {
				accept = hasRequiredRole(principal, scopedRequest.RequiredScopes)
			}

			return accept, principal, err
		}

		return false, nil, nil
	})
}

func updatePrincipal(existing, incoming *models.Principal) {
	if incoming.Role != "" {
		existing.Role = incoming.Role
	}

	if incoming.Provider != "" {
		existing.Provider = incoming.Provider
	}
}

func authenticateAPIKey(name, in string, ta security.TokenAuthentication, r *http.Request) (accept bool, principle interface{}, authError error) {
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
		updatePrincipal(existPrincipal.(*models.Principal), authPrincipal)

		return ctx, existPrincipal, nil
	})

	return authFunc.Authenticate(r)
}

func setupAPIKeyAuth(api *operations.RestGoswaggerAPIServerAPI) {
	api.APIKeyAuthenticator = func(name, in string, ta security.TokenAuthentication) runtime.Authenticator {
		return HttpAuthenticator(name, func(r *http.Request) (accept bool, principle interface{}, authError error) {
			return authenticateAPIKey(name, in, ta, r)
		})
	}
}

func setupKeyAuth(api *operations.RestGoswaggerAPIServerAPI, rt *rest_goswagger_api.Runtime) {
	api.KeyAuth = func(token string) (*models.Principal, error) {
		if len(token) > 0 && token == rt.Config().GetString("x_app_key") {
			return &models.Principal{}, nil
		}
		return nil, rt.SetError(http.StatusUnauthorized, fmt.Sprintf("invalid key: %v", token))
	}
}

func setupRoleAuth(api *operations.RestGoswaggerAPIServerAPI, rt *rest_goswagger_api.Runtime) {
	// TODO: uncomment when some endpoint has implement that security
	// api.HasRoleAuth = func(role string) (*models.Principal, error) {
	// 	if role == "" {
	// 		return nil, rt.SetError(http.StatusForbidden, "access denied for this role")
	// 	}
	// 	return &models.Principal{Role: role}, nil
	// }
}

func setupRoleAuthPayloadJWT(api *operations.RestGoswaggerAPIServerAPI, rt *rest_goswagger_api.Runtime) {
	// TODO: uncomment when some endpoint has implement that security
	api.HasRoleAdminAuth = func(token string) (*models.Principal, error) {
		return checkHasRoleAuthPayloadJWT(rt, constants.ROLE_SLUG_ADMIN, token)
	}

	// TODO: uncomment when some endpoint has implement that security
	// api.HasRoleUserAuth = func(token string) (*models.Principal, error) {
	// 	return checkHasRoleAuthPayloadJWT(rt, constants.SLUG_ROLE_USER, token)
	// }

}

func setupPeopleAuth(api *operations.RestGoswaggerAPIServerAPI, rt *rest_goswagger_api.Runtime) {
	// TODO: uncomment when some endpoint has implement that security
	// api.IsPeopleAuth = func(people string) (*models.Principal, error) {
	// 	if people == "" {
	// 		return nil, rt.SetError(http.StatusForbidden, "access denied for this people")
	// 	}
	// 	return &models.Principal{People: people}, nil
	// }
}

func setupAPIAuthorizer(api *operations.RestGoswaggerAPIServerAPI, rt *rest_goswagger_api.Runtime) {
	api.APIAuthorizer = NewAuthorizationApi(rt)
}

func Authorization(rt *rest_goswagger_api.Runtime, api *operations.RestGoswaggerAPIServerAPI) {
	setupAPIKeyAuth(api)
	setupKeyAuth(api, rt)
	setupRoleAuth(api, rt)
	setupRoleAuthPayloadJWT(api, rt)
	setupPeopleAuth(api, rt)
	setupAPIAuthorizer(api, rt)
}

func parseToken(rt *rest_goswagger_api.Runtime, token string) (*jwt.PayloadJWT, error) {
	secret := rt.Conf.GetString("jwt.secret")
	maker, err := jwt.NewJWTMaker(secret)
	if err != nil {
		return nil, err
	}

	payload, err := maker.VerifyToken(token)
	if err != nil {
		return nil, rt.SetError(http.StatusUnauthorized, "invalid API key token: %v", err)
	}

	return payload, nil
}

func verifySingleRole(payload *jwt.PayloadJWT, role string) (*models.Principal, error) {
	if payload.RoleSlug != role {
		return nil, openapi_errors.New(http.StatusForbidden, "insufficient API key privileges")
	}

	return &models.Principal{
		UserID:    payload.UserID,
		ExpiredAt: payload.ExpiredAt.Unix(),
		Email:     payload.Email,
		Role:      payload.RoleSlug,
	}, nil
}

func checkHasRoleAuthPayloadJWT(rt *rest_goswagger_api.Runtime, role, token string) (*models.Principal, error) {
	payload, err := parseToken(rt, token)
	if err != nil {
		return nil, err
	}

	p, err := verifySingleRole(payload, role)
	if err != nil {
		return nil, err
	}

	return p, nil
}
