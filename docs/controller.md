# Controller

Provides base functionality and dependencies to routes.
The pattern is to embed a Controller in each individual route struct
and to use the router to inject the Container so the routes have access
to the services within the container.

## Struct

- Container *services.Container

## Methods

### NewController *services.Container -> Controller

recives a pointer to the container instance
returns an instance of Controller

### RenderPage *Self echo.Context Page -> error

Renders a Page as HTTP response
Error if the page has no name
or if fails the parse or execution of the templates
Checks the AppName
Checks if its an HTMX request, wich indicatess that only partial content should
be rendered
Then the Page is rendered by .Container.TemplateRenderer
Sets response status = Page.StatusCode
Sets headers
Applies HTMX response if one
Caches the page if its enabled
returns Echo's HTMLBlob

### cachePage NOT IMPLEMENTED YET

### Redirect *Self echo.Context string ...any -> error

Redirects to a given route name with optional parameters
recives context, route and routeParams

### Fail *Self error string -> error

Helper to fail a request by returning 5000 error and logging

# Page

Consist of all data that will be used to render a page response for a given
controller.
Methods also become available in the templates.

## Struct

- AppName string
- Title string // Title of the page
- Context echo.Context // request context
- ToURL func(name string, params ...any) string
    // function to convert route name an optional params to a URL
- Path string // path of the current request
- ULR string // URL of the current request
- Data any // Stores whatever additional data that needs to be pased to the templates
- Form any // Stores a struct that represents a form on the page
- IsAuth bool
- AuthUser // TODO USER
- StatusCode int
- Metatags struct { Description string, Keywords []string }
- Pager Pager
