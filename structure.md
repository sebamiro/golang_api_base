# Backend

- Echo: Go web framework.
- Ent: Orm

# Frontend

- HTMX
- Bulma: Style

# Storage

- TODO: Define DB (?Postgers)
- TODO: ?Redis

---

## Sessions

Gorilla sessions, stored in a cookie, as a middleware

## Authentication

Service `AuthClient`
`bcypt` to encrypt sensible values

### Login / Logout

`AuthClient` has `Login()` and `Logout()` methods

Prior logging a user `CheckPasswoird()`, checks pass in `User` Entity

routes user/login user/logout

### Forgot Password

`GenereatePasswordResetToken()` creates new `PasswordToken` Entity belonging
to the user. Token is stored encrypted

`GetValidPasswordToken()` will load a matching, non-expired password token
entity belonging to the user and verify the token sent in the url

Once is claimed `DeletePasswordTokens()`

routes user/password user/password/reset/:user/:password_token/:token

### Registration

`AuthClient`, `HashPassword()`, and creates `User` entity

routes user/register

### Authenticated user

`AuthClient`, has `GetAuthenticatedUser()` and `GetAuthenticatedUserID()`

### Middleware

For all routes is a middleware that loads currently logged in user entity
and store it within the requst context. `middleware.LoadAuthenticatedUser()`

`middleware.RequireAuthentication()` or `middleware.RequireNoAuthentication()`

## Email Verification

Not now, but is good idea

## Routes

Router functionality is provided by Echo
`BuildRouter`

### Controller

`Controller` provides base functionality for each route

### Patterns

To create a route, create a new route that embeds the `Controller`

    type home struct {
        controller.Controller
    }

    func (*home) Get(c echo.Context) error {}
    func (*home) Post(c echo.Context) error {}

Then create route and add to the router

    home := home{Controller: controller.NewController()}
    g.GET("/", home.Get).Name = "home"
    g.POST("/", home.Post).Name = "home.post"


### Errors

`ErrorHandler` whould be a clever way to handle errors

### Testing

Wise
