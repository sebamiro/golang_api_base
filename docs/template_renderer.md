# TemplateRenderer

## Struct

- templateCache sync.Map
    sync stdlib, stores a cache of parsed page templates

- funcMap template.FuncMap
    template html/template, stores the template function map

- config *config.Config
    stores app configuration

## Methods

### New *config.Config -> *TemplateRenderer

recives the app config as parameter
returns a new instance of TemplateRenderer
    templateCache: sync.Map{}
    funcMap: funcmap.GetFuncMap(),
    config: param

### Parse *TemplateRenderer -> *templateBuilder

returns new templateBuilder with self and new templateBuild

### getCacheKey *Self string string -> string

gets a cache key for a given group and ID

recives group and key as parameters
returns "<group>:<key>" if group is not empty else "<key>"

### parse *Self *templateBuild -> (*TemplateParsed, error)

parses a set of templates and caches them for quick execution
If the application environment i set to local, the cahe will be
bypassed and templates will be parsed upon each request so hot-reloading
is possible without restarts

recives build
returns the parsed and cached template or an error if the parsing fails

### Load *Self string string -> (*TemplateParsed, error)

recives group and key
returns cached template or an error if its not cached or if it
cant be casted to *TemplateParsed

# templateBuilder

handles chaining a template parse operation

## Struct

- build *templateBuild
    data of the teplate
- renderer *TemplateRenderer
    instance of the TemplateRenderer

## Methods

### Group *Self string -> *templateBuilder

sets the cache group for the template being built

recives group and stores it in .build
returns itself

### Key *Self string -> *templateBuilder

sets the cache key for the template being built

recives key and stores it in .build
returns itself

### Base *Self string -> *templateBuilder

sets the name of the base template to be used during template parsing and
execution.

recives the basename
returns itself

### Files *Self ...string -> *templateBuilder

sets a list of template files to be included in the parse
Should not include file execution and paths should be relative to template dir

recives list of files
returns itself

### Directories *Self ...string -> *templateBuilder

sets a list of directories that all template files within will be parsed
The paths should be relative to template dir

recives list of directories
returns itself

### Store *Self -> (*TemplateParsed, error)

parses the templates and stores the in the cache

return .renderer.parse(.build)

### Execute *Self any -> (*bytes.Buffer, error)

execues template with the given data
If its not cached, it will parse and cache it

recives the data
returns TemplateParsed.Execute(data)

# templateBuild

stores the build data used to parse a template

## Struct

- group string
- key string
- base string
- files []string
- directories []string

# TemplateParsed

is a wrapper around parsed templates wich are stored in the TemplateRenderer cache

## Struct

- Template *template.Template
- build *templateBuild

## Methods

### Execute *Self any -> (*bytes.Buffer, error)

excutes a template with the given data and provides the output

recives data
returns the output or an error if the Template is nil or if the execution fails
