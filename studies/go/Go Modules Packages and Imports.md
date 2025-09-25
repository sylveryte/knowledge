# Modules Packages and Imports

## Repository

- a git repository
- 🐥 while it is possible you should _not_ keep
  more than 1 [[#Module]] in a repository

## Module

- A package in javascript
- Every module has a globally unique identifier it is
  called `module path` usually where a module is stored

  - it is case sensitive

- Simple [module structure](https://go.dev/doc/modules/layout)
- Are generally of two types
- Single application
- A Library
  - Root of module should have package name that maches
    the repository name
  - Uncommon for lib modules to have one or more applicatoins
    included with them as utils. In case you have create a
    `cmd` dir at root and create 1dir for each app.
  - Use `main` as the package name within each dir

### Using Go mod

- To initialize go mod [[Go Init#Initialize project/module]]
- `go mod init example/hello`

#### Go build

- Since go `1.21` the default behavior is to download new
  version of go to build
- Can control this behavior by using `GOTOOLCHAIN` env var
  - auto: download new and build
  * local: pre 1.21 behavior
  * [ specific ]: specify go version
- More `GOTOOLCHAIN` on docs

#### Require directive

- Only present if you have dependency
- They list required modules with min versions required
- First require section is _direct_ dependencies
- Second is _indirect_ dependencies (deps of directs)
  - Each line ends with `//` to indicate,
  - there is no significance other than docs 🐦

### Third party modules

- Import same as `fmt`
- The `go get` command downloads modules and updates
  `go.mod` file
- If it doesn't have a version tag,
  go makes a pseudoversion
- In addition of updating `go.mod`, a `go.sum` file is created.
  - If you remove or change mod, an entry in `go.sum` might
    remains /no issue/ though
  - `go.sum` as one or two entries for each modules
    - mod, version and hash of module
    - hash of `go.mod` file for the module
- Go maintains a module cache on local
- To delete local cache `go clean -modcache`
- `go mod tidy` fixes some comments
- `go list` to list packages used in your mod
  - -m to list modules
  - -versions flag to specified module

#### Versions

- Semantic v`major.minor.patch`
- You'll always get the `minimal version selection` deps
  that is declared to work in go.mod files across all deps
- To upgrate to a bug patch release
  - use command `go get -u=patch`

### Vendoring

- To esure that a module always builds with identical
  deps, you can keep copies of the deps inside their
  module it's called Vendoring
- `go mod vendor` creates dir called vendor at the top
  level of your module that contains all your modules deps
- If new deps are added to `go.mod` you have to run `go mod vendor`
  to update else there will be errors ⭕

### Docs

- use [pkg.go.dev](pkg.go.dev) for module docs

### Overriding deps

- use `replace [old mod namae] => [new mod name]`
- local replace possible but avoid it, use workspaces

### Workspaces

- to modify modules simultaneously
- don't point local dirs with go.mod ☢️
- `go work init ./workspace_app`
- `go work use ./workspace_lib`

### Module Proxy Servers

- Every module is somewhere like github or gitlab but by default,
  go get fetches from a `proxy server` run by Google ⏰
- You can disable by using `GOPROXY=direct` env
- You can run your own proxy server

### Package

- A module in javascript
- First line is package clause
- Go doesn't allow `circular dependency` 🚧
- Hyphen is not a valid character in a package name

### Importing Exporting

- Go's import statement allows to use exported const,
  funcs, types, vars
- 🍏 Uses /first/ letter _Capitalization_ to indicate _exported_
- Import path consists of path to module followed by
  package name
- 🍎 /Relative path/ is bad idea and in the past
- _aliasing_ package

```go
import (
  crand "crypto/rand" // alias is crand
  "math/rand"
)
```

- You can use more prefix for
  - dot (.) to make the exports available without namescope,
    like javascript `import * dayjs`
- underscope (\_) to run the [[#Side Effect]] of the package

### Naming

- Good idea to go with student/teacher/school kind package

### Internal Package

- create a package called `internal`
  - exported identifiers are accessible to direct parent
    and sibling packages

### Side Effect

- a `func init()` with no args, no return is called side effect
- It is called when the pkg is referenced
- Avoid it

### Documentations

- It's in a separate file [[Go Documentations]]

---

[[Learning Go]]
