
# Codoworks Go Boilerplate

[![Go Build](https://github.com/codoworks/go-boilerplate/actions/workflows/go.yml/badge.svg)](https://github.com/codoworks/go-boilerplate/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/codoworks/go-boilerplate)](https://goreportcard.com/report/github.com/codoworks/go-boilerplate)
[![CodeQL](https://github.com/codoworks/go-boilerplate/actions/workflows/codeql.yml/badge.svg)](https://github.com/codoworks/go-boilerplate/actions/workflows/codeql.yml)
[![Codacy Security Scan](https://github.com/codoworks/go-boilerplate/actions/workflows/codacy.yml/badge.svg)](https://github.com/codoworks/go-boilerplate/actions/workflows/codacy.yml)
[![Actively Maintained](https://img.shields.io/badge/Maintenance%20Level-Actively%20Maintained-green.svg)](https://gist.github.com/cheerfulstoic/d107229326a01ff0f333a1d3476e068d)

This is a backend service skeleton or boilerplate to speed up development process. Over time, this package has become opinionated and behaves more like a framework with a set of predefined features.
This package was developed using `GoLang v1.21`. 
Thank you for using Codoworks Go Boilerplate. 

<img align="right" src="docs/Go-Boilerplate_Logo_Dark_BG_Small_V1.0.png">

```
   ___          _                         _        
  / __\___   __| | _____      _____  _ __| | _____ 
 / /  / _ \ / _  |/ _ \ \ /\ / / _ \| '__| |/ / __|
/ /__| (_) | (_| | (_) \ V  V / (_) | |  |   <\__ \
\____/\___/ \__,_|\___/ \_/\_/ \___/|_|  |_|\_\___/ cli.

```

<details>
<summary><b>Mindset</b></summary>
	
This service is designed for developers to build backend API as quickly as possible - almost as simply as copy and pasting components. The goal is to be able to clone this repository, rename it and get started with your first RESTful CRUD API within minutes.

There are many HTTP libraries on GitHub that enable quick and easy API development, but as your project scales it can quickly become messy without proper structure and workflows, and even more when preparing for production or managing security.

Codoworks Go Boilerplate has all your production needs taken care of in advance, so you can focus your efforts on creating business and application logics.

It's worth noting that this package builds upon [Echo](https://echo.labstack.com) but you can easily adapt it to a different framework.

#### Multi-routers

It's designed with 3 separate routers (public, protected and hidden). Each router has its individual configuration that you can customise to your needs. This enables the following structure:

```
https://your-domain.com/api/<your public api>            # Publicly consumable API
https://api.your-domain.com/<your protected api>         # Your application's API 
https://api.your-domain.com/.admin/<your hidden api>     # Administrative API, not supposed to be available via public internet
```

With this structure, the default router is assumed to be the protected one as most use cases tend to contain a user interface element with a login functionality. 

#### Modular

Speaking of login, this service is designed to be used with self-hosted [ory kratos](https://www.ory.sh/open-source/) for authentication. Since this boilerplate is designed to be modular, you may choose another service provider instead of Kratos. All you have to do is change the authentication.go middleware to your desired service. 

Similarly for authorisation, [ory keto](https://www.ory.sh/open-source/) is the default client for this service as it is well-designed to manage large volumes of transactions.

#### Maintenance

Often, you may need to run something in the background like a clean up job or perhaps an email watcher. This service is designed with that in mind too. It even provides a watcher that you can run with an http server or independently. Using the same structure you can create your own too.

#### Administration

A well-designed service should also enable the user to perform certain administrative tasks out of the box, like providing a specific user with a given email address or system admin permission. It's likely that a user interface for such feature is not a priority on your product roadmap, hence an API is never developed and as a result, you cannot make use of Postman. To prevent this type of scenario, this boilerplate is shipped with tasks that you can easily extend and execute via command line.

#### Database

Needless to say, almost every service requires a database, this one is no different too. Codowworks Go Boilerplate is designed with MySQL, Postgres and SQLite integration. By default, it uses SQLite to allow a quick start, switching platforms is just a matter of changing an environment variable.

#### Others

Within this boilerplate are also database migrations, logging, routing, hot-reloading, CORS, timeout and even graceful shutdown, which are some of the features you need to get to production as soon as possible. 

We hope you enjoy using Codoworks Go Boilerplate. If you do, please support us by giving this repository a star.
</details>

<details>

<summary><b>Full Feature List</b></summary>

- CLI commands (via Cobra)
- HTTP server (via Echo)
  - Public router
  - Protected router
  - Hidden router
- Daemon processes or workers
- Tasks for custom one-off operations
- Middlewares
  - HTTP header checks and setters
  - Auto error handling and response
  - Authentication via Ory Kratos
  - Authorisation via Ory Keto
  - CORS handling
  - Auto trim trailing slashes
  - Request timeout
  - Gzip responses
  - XSS check
- Databases
  - DB connection (PostgreSQL, MySQL)
  - DB models (ORM using Gorm)
  - DB migrations (using Go Migrate)
  - DB seeds (using Go Migrate)
- JSON forms and model mapping
- Data validation
- Clients
  - Forward HTTP client to forward authorization headers
  - Ory Kratos Client - authentication
  - Ory Keto Client - authorization
- Custom logger
- Graceful shutdown
- Feature toggle: [ory_kratos, ory_keto, db, redis]
</details>

# Getting Started

1. Clone this repository `git clone git@github.com:codoworks/go-boilerplate.git`
2. Run `cd go-boilerplate`
3. Run `go get`
4. Run `go run . db migrate`
5. Run `go run . db seed`
6. Run `go run . start` to start the server, you should see the following:
```
⇨ http server started on [::]:8081
⇨ http server started on [::]:8080
⇨ http server started on [::]:8079
```
7. List available routes using `go run . info protected-api-routes` and use your favourite API client to test. or use the following to get started and make sure you're up and running.
```bash
curl -H "Accept: application/json" http://127.0.0.1:8081/health/alive
curl -H "Accept: application/json" http://127.0.0.1:8081/health/ready
```

> Recommended: run `go run .` and explore all available options, it should be straightforward.

For more details on running and using the service, scroll down to "[Operations](#operations)" section. 

To learn about developing and extending this service, scroll down to "[make it your own](#make-it-your-own)" section. 

#### Simplified Architecture Diagram

<img align="middle" src="docs/Go-Boilerplate_arch.png">

<details>
<summary><b>Docker</b></summary>

The service is shipped with a few Docker compose files to get you started, all of which are automated with a Makefile to make things consistent.

#### Quick Start

From the boilerplate root folder, run the quick-start target from the Makefile.

```bash
make quick-start
```

#### Quick Start with MySQL

To run an example using MySQL database, from the boilerplate root folder, run the quick-start-mysql target from the Makefile.

```bash
make quick-start-mysql
```

#### Quick Start with Postgres

To run an example using Postgres database, from the boilerplate root folder, run the quick-start-mysql target from the Makefile.

```bash
make quick-start-postgres
```
</details>


<details>
<summary><b>Usage</b></summary>

### Env Vars

Environment variables are evaluated in the following order to allow flexibility when running in production: 
1. `.env` file
2. environment variables 
3. cmd flags (if available)

During development, it is recommended to use a `.env` file. You can find a reference under /.env.sample` to get started. 

To ease your development process, we've included a command to print the environment to better understand your app behaviour. Simply run `go run . info env`. Together with `go run . info features`, you should be able to get to the bottom of an issue. 


<details>
<summary><b>List of all available env vars</b></summary>

| Var Name | Required | Description |
| -------- | -------- | ----------- |
| `HOST` | optional | service host address. default: 0.0.0.0 |
| `PROTECTED_API_PORT` | optional | Service port. Default: 8080 |
| `PUBLIC_API_PORT` | optional | Service port. Default: 8081 |
| `HIDDEN_API_PORT` | optional | Service port. Default: 8079 |
| `DB_HOST` | optional | Database host |
| `DB_PORT` | optional | Database port |
| `DB_USER` | optional | Database username |
| `DB_PASSWORD` | optional | Database password |
| `DB_NAME` | optional | Database name |
| `DB_TIMEZONE` | optional | Database timezone. Required with Postgres platform |
| `DB_PLATFORM` | optional | Enum: ["postgres", "mysql", "sqlite"]. default: "sqlite" |
| `KRATOS_PUBLIC_SERVICE` | optional | Ory Kratos public API URL |
| `KRATOS_ADMIN_SERVICE` | optional | Ory Kratos admin API URL |
| `KETO_READ_SERVICE` | optional | Ory Keto read API URL |
| `KETO_WRITE_SERVICE` | optional | Ory Keto write API URL |
| `REDIS_HOST` | optional | Redis host URL. Required if Redis is enabled |
| `REDIS_PORT` | optional | Redis port. Required if Redis is enabled |
| `REDIS_PASSWORD` | optional | Redis password. Required if Redis is enabled |
| `LOG_LEVEL` | optional | Enum: ["info", "warn", "debug", "error"]. default: "info" |
| `CORS_ALLOW_ORIGINS` | optional | Allowed origins. Default: "*" |
| `REQUEST_TIMEOUT_DURATION` | optional | Number in seconds. Default: "60" |
| `DISABLE_FEATURES` | optional | List of features to disable in runtime, make sure its comma separated without spaces |
</details>

### Execution Modes

The service can run in one of two modes: production mode or development mode. 

Development mode is activated using the `-d` or `--dev` flag. Running in this mode will lock the service host to `127.0.0.1` to avoid firewall issues when developing using MacOS. You can override this setting using `-H 0.0.0.0` if needed. 

Development mode will also activate useful middlewares that help print incoming request body, input data validation errors for debugging, and set the logger level to debug for ease of development. Everything else is identical to running in production mode.

You can change the behaviour of the service using flags, see the list of flags below for more.

<details>
<summary><b>List of all flags</b></summary>

| Flag Name | Shorthand | type | Description |
| --------- | --------- | ---- | ----------- |
| `--dev` | `-d` | bool | Run in development mode |
| `--env` | `-e` | bool | Print environment variables |
| `--host` | `-H` | string | (optional) Service host. Overrides env vars |
| `--port` | `-P` | string | (optional) Service port. Overrides env vars |
| `--watcher` | (N/A) | bool | (optional) Start watcher in the backgoround |
| `--log` | `-l` | string | (optional) Log level |
</details>

### Live Reload / Hot-swap <a name="live-reload"></a>

It is convenient to automatically restart the service every time you save your changes. For that, you can use [air](https://github.com/cosmtrek/air), which is a separate Go package you can install using the following command:

```bash
go install github.com/cosmtrek/air@latest
```

Once `air` is installed, you simply need to run `air` to start the service. Configurations for this can be found under `./.air.toml`.

Live reloading will also work in Docker. The `Dockerfile.dev` is configured to install and run the service via `air`.

### Operations <a name="operations"></a>

This service is shipped with a cmd client, which means you can use `./go-boilerplate` to view all available commands and help menu.

> - You need to build the service first before you can use `go-boilerplate` 
> - both `./go-boilerplate` and `go run .` can be followed by any flags, commands and sub-commands 

### Required Headers

The service requires `Accept: application/json` header by default for all requests. 

It also requires `Content-Type: application/json` with `POST`, `PUT` and `DELETE` requests.

### Native Development 

If you're writing a small project with a few endpoints then running Go in your terminal shouldn't be much of a problem. You can use [live-reload](#live-reload) while you're editing your code in your favourite editor. 

To run the service without building, run `go run .` which will achieve the same result as running `./go-boilerplate` after building the binary.

> The name `go-boilerplate` will change if you change the package name [as mentioned here](#change-pkg-name).

### In-Docker Development

However, when you are running a large project with multiple micro-services (multiple instances of this boilerplate), it can be handy to live edit your code while in Docker. For this, we have designed the `Dockerfile.dev` to get you started.

Simply run `make quick-start` to get up and running. To stop it, use `Ctrl+C`.

### Build

To build, run `go build .` which will generate a binary with the default name of the package. In this case, it will be `./go-boilerplate` unless you change it (which is recommended).

If you have executed the above, you may notice that the version `./go-boilerplate version` is set to `2.x.x-default` during run time. That's because it is the second iteration of this boilerplate. It is recommended that you burn the version into the binary in build time to create versioned builds. To do that, use the following command to build:

```bash
go build -ldflags="-w -s -extldflags '-static' -X main.VERSION=<YOUR.VERION.HERE>"

# Example
go build -ldflags="-w -s -extldflags '-static' -X main.VERSION=1.0.0"
./go-boilerplate version
# v1.0.0
```

Once built, a single binary file is generated. It is an executable file that you can rename and place in any folder as long as your profile PATH can find it. A good place to place it on your local machine would be in `/usr/bin` which is where most binaries are. 

### Deployment

If you wish to deploy this service locally, all you need to do is build as per the section above then ship the outputted binary into a location where your terminal's PATH can find it. You should be able to use it just by calling its name in your terminal. 

The "Usage" section should get you familiarised with all the parameters that are configurable. Deploying it should not be a problem in any dockerised environment.

From a containerisation perspective, I'd encourage you to place this binary in an empty container i.e. `FROM scratch` in your Dockerfile. This helps keep the container size to a minimum. When tested on an M1 Mac Machine, we got an 18MB container. 
</details>

<a name="make-it-your-own"></a>
<details>
<summary><b>Extending the service (make it your own)</b></summary>

This section is all about extending the service to create your own application and APIs. 

<a name="change-pkg-name"></a>
> The first thing you should do is to change the package name, find `github.com/codoworks/go-boilerplate` in all the files and replace it with your own package name. You can choose to use the general `github.com/(org-name)/(project-name)` naming pattern for consistency.


<details>
<summary><b>Migrations</b></summary>

Migrations help create your database and track how it evolves overtime. Here, we use [GoMigrate](https://github.com/go-gormigrate/gormigrate) to achieve this. Some added complexity is added to enable easy extendability and generate better logs throughout your development process. 

Migrations go under `pkg/db/migrations/<myNewMigration>.go`. Its implemention uses `Go`'s `init()` function, which means they're added to the list in alphabetical order. They migrate in that order (top to bottom) and rollback in the reverse order (bottom up). For this, it is best to maintain the naming convention of `YYYYMMDD[00-99]_migration_description`.


Here's a sample migration to get you started:
```Go
func init() {
	m := &gormigrate.Migration{}

	m.ID = "2022081801_create_cats_table"

	m.Migrate = func(db *gorm.DB) error {
		type Cat struct {
			models.ModelBase
			Name string `gorm:"size:255"`
			Type string `gorm:"size:255"`
		}

		return AutoMigrateAndLog(db, &Cat{}, m.ID)
	}

	m.Rollback = func(db *gorm.DB) error {
		if err := db.Migrator().DropTable("cats"); err != nil {
			logFail(m.ID, err, true)
		}
		logSuccess(m.ID, true)
		return nil
	}

	AddMigration(m)
}
```

The variable `m` holds the migration details and is added to the list of migrations at the end. `m.ID` is the identifier used by `gomigrate` to keep track of the migrations that already ran. So, make sure to change that for every migration.


Every migration has 2 methods to be implemented, the `Migrate()` and `Rollback()` method as described above. Make sure you use the `logSuccess`, `logFail` and `AutoMigrateAndLog()` functions to print the migrations that ran. This will come in very handy for remote deployments. 

It's recommended to declare your models within each migration (separately from the models package) to keep track of the database schema change through time. You can add or delete columns, rename columns, and execute raw SQL in migrations.

> A general good practice would be to flatten your migrations once your application achieves version 1, leaving only neat table creation in each migration.
</details>

<details>
<summary><b>Seeds</b></summary>

Seeds are very similar to migrations, but seeds do not implement the `Rollback` function.

Just like migrations, seeds are applied once and tracked using their unique identifier `ID` by [GoMigrate](https://github.com/go-gormigrate/gormigrate).

Seeds are part of the whole package which allows you to access models, clients and other components directly to configure the application, and perhaps provide dummy data to help with development. 

Here's a seed skeleton to get you started. Copy the following structure into a new file under seeds and change the `s.ID` property.

```Go
func init() {

	var s = &gormigrate.Migration{}
	s.ID = "2022081801_new_seed"

	s.Migrate = func(db *gorm.DB) error {

		logSuccess(s.ID)
		return nil
	}

	AddSeed(s)
}
```

And here's a sample seed to give an idea of how you can utilise seeds.

```Go
func init() {

	var s = &gormigrate.Migration{}
	s.ID = "2022081801_seed_cats_data"

	s.Migrate = func(db *gorm.DB) error {

		var err error
    var cats []*models.Cat

		cats = append(cats, &models.Cat{
			Name: "Kitty",
			Type: "Persian",
		})

		cats = append(cats, &models.Cat{
			Name: "Tom",
			Type: "Siamese",
		})

    for _, cat := range cats {
      err = cat.Save()
      if err != nil {
        logFail(s.ID, err)
        return err
      }
    }

		logSuccess(s.ID)
		return nil
	}

	AddSeed(s)
}
```
</details>

<details>
<summary><b>Models</b></summary>

Models can sometimes be a complex aspect of any application. In this section, you'll find a rundown on how you can compose your models or database entities. 

#### Model Structure 

The first thing is to create a struct that matches your database schema. Almost all models should embed the `ModelBase` struct that provides the `ID`, `CreatedAt` and `UpdatedAt` properties. Exceptions may include a many-to-many table where you only need to store 2 identifiers. To learn more about model declarations, you can refer to [Gorm](https://gorm.io)'s official comprehensive documentation. 

Here's a Cat model that should correspond to a Cats table that contains 5 properties i.e. `ID`, `CreatedAt`, `UpdatedAt`, `Name` and `Type` in a database.

```Go
type Cat struct {
	ModelBase
	Name string `gorm:"size:255"`
	Type string `gorm:"size:255"`
}
```
Notice how every property contains a `gorm` decoration to specify things like field size, uniqueness or foreign keys etc. For more details, please refer to [Gorm](https://gorm.io)'s documentation.

Your model may sometimes contain properties that do not correspond to a database column. To do that, you simply need to use the `gorm:"-"` decoration. 

> Note: Given that this package is designed to work with multiple database servers like MySQL or Postgres, some data types may be available in some servers and not others. It's worth testing your application with different servers from time to time to accomodate easy switching of database server, unless your use case relies on a specific data type - in which case you're making a calculated decision to lock your application to that server.

#### Common Basic Functionality

Now that you have a structure that corresponds to a table in your database, some common functionality is in order. Generally, one would at least expect the basic CRUD functionality. Here's a basic CRUD implementation that is required for any model:

- `FindAll()`, for retrieving all records in the table
```Go
func (model *Cat) FindAll() (models []*Cat, err error) {
	result := db.Model(model).Find(&models)
	return models, result.Error
}
```

- `FindMany()`, for retrieving many items given an array of IDs
```Go
func (model *Cat) FindMany(ids []string) (models []*Cat, err error) {
	result := db.Model(model).Find(&models, ids)
	return models, result.Error
}
```

- `Find()`, for retrieving a single item with a given ID
```Go
func (model *Cat) Find(id string) (m *Cat, err error) {
	result := db.Model(model).Where("ID=?", id).First(&m)
	return m, result.Error
}
```

- `Save()`, for creating a new record in the database and assigning a new ID to it
```Go
func (model *Cat) Save() error {
	return db.Model(model).Create(&model).Error
}
```

- `'Update()`, for updating a record in the database given an existing ID
```Go
func (model *Cat) Update() error {
	return db.Model(model).Updates(&model).Error
}
```

- `Delete()`, for deleting a record in the database given an existing ID
```Go
func (model *Cat) Delete(id string) error {
	return db.Model(model).Where("ID=?", id).Delete(&model).Error
}
```

All of the above functions will return an error if they cannot perform what they're supposed to. That's useful to inform users if the data they're looking for exists or is stored. For detailed utilisation of these functions, check out the handlers folder. 

These functions are not abstracted to allow granular control over each model, as each individual model can quickly morph into something very large with child elements, preload functions and pagination. 


#### Model Accessibility

Given the basic functionality defined in the previous section, we've created the ability to do something like the following:
```Go
...
catModel := &Cat{}

myCat, err := catModel.Find(catID)
if err != nil {
	fmt.Println("couldn't find cat with ID", catID)
}
...
```

The problem with the code above is that you will need to instantiate a new struct `catModel` from `&Cat{}` in order to have a pointer receiver that can call the `Find()` function. You can avoid that by using the following common getter structure for all models, right at the top of the model before its declaration to maintain consistency.
```Go
var cat *Cat = &Cat{}

func CatModel() *Cat {
	return cat
}
```
The above will now create a singleton pattern that you can access from any component within the package like `models.CatModel().Find()`.

> Note: the `CatModel()` method should only be used to fetch data from the database. Saving, updating and deleting data should be applied to an actual instance that has been returned through a `Find()`, `FindAll()` or `FindMany()` function.

#### Working with JSON Forms

Once you have retrieved the records needed from the database, you may want to send those records as a response. To do that, you can use forms. Every model is expected to have at least one method named `MapToForm()` that returns a JSON representation of that model. 

Forms are basic structures that may or may not exactly match all the properties that a model has. The reason it has been done this way is to enable multiple forms where one can contain all model properties e.g. intended for an admin user to view, while another may contain a sanitised version of that model e.g. intended only for a read-only user. 

For more details on creating a form, scroll down to the forms section below. Here you'll find a sample implementation of `MapToForm()` function.
```Go
func (model *Cat) MapToForm() *CatForm {
	form := &CatForm{
		Name: model.Name,
		Type: model.Type,
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}
```

#### Complete Code 

Here's a complete code as a model sample that you can copy as a base model.

```Go
var cat *Cat = &Cat{}

func CatModel() *Cat {
	return cat
}

type Cat struct {
	ModelBase
	Name string `gorm:"size:255"`
	Type string `gorm:"size:255"`
}

func (model *Cat) MapToForm() *CatForm {
	form := &CatForm{
		Name: model.Name,
		Type: model.Type,
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}

func (model *Cat) FindAll() (models []*Cat, err error) {
	result := db.Model(model).Find(&models)
	return models, result.Error
}

func (model *Cat) FindMany(ids []string) (models []*Cat, err error) {
	result := db.Model(model).Find(&models, ids)
	return models, result.Error
}

func (model *Cat) Find(id string) (m *Cat, err error) {
	result := db.Model(model).Where("ID=?", id).First(&m)
	return m, result.Error
}

func (model *Cat) Save() error {
	return db.Model(model).Create(&model).Error
}

func (model *Cat) Update() error {
	return db.Model(model).Updates(&model).Error
}

func (model *Cat) Delete(id string) error {
	return db.Model(model).Where("ID=?", id).Delete(&model).Error
}
```
Copy the code above and replace the name `Cat` to get started. 
</details>

<details>
<summary><b>Forms</b></summary>

Forms are data contracts that are used to send responses to clients and receive/ bind user input. 

Each model can have many forms to enable sending specific values with different endpoints. An example scenario would be having an admin with full access to all data in a record whereas a customer has access only to a subset of that data.

Data validation is applied to fields in forms. Here's a sample form to get you started.
```Go
type CatForm struct {
	FormBase
	Name string `json:"name" validate:"required,min=2,max=50"`
	Type string `json:"type" validate:"required,min=2,max=80"`
}

func (form *CatForm) MapToModel() *Cat {
	return &Cat{
		Name: form.Name,
		Type: form.Type,
	}
}
```

The `FormBase` struct provides the `ID`, `CreatedAt` and `UpdatedAt` fields. 

Each field should specify the name mapping in JSON format along with validation rules. For more on validations check out the [Playground Validator documentation](https://github.com/go-playground/validator). To skip validations all together, use `validate:"-"`.

Finally, each form should have a `MapToModel()` function that returns a model, so it can be stored after it has been validated. Note that forms do not set a model's `ID` property as that is the job of the model. Instead, it must be set manually prior to a database operation. Think of this like an actual form you fill up that has a section "for office use only". 
</details>

<details>
<summary><b>Handlers</b></summary>

> Note: This go boilerplate uses [Echo](https://echo.labstack.com). If you're ever in doubt, you can refer to Echo's documentation for more details on what's possible with routers.

A handler is any function with the `func (c echo.Context) error` signature. All handlers should be stored under `pkg/api/handlers` and categorized in directories following their entity name in plural form. For readability and maintainability, we encourage maintaining a single handler in a single file as we all know that Go files can quickly grow.

Handlers should also be nested, which means a Cats handlers directory can contain a sub directory, such as `cats/tags`, that helps avoid long file names and improve readability. 

How handlers look will largely depend on your project's business logic and requirements. For reference, here's a quick sample to give an idea on how you should construct your handler.

```Go
func Get(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return helpers.Error(c, constants.ERROR_ID_NOT_FOUND, nil)
	}

	m, err := models.CatModel().Find(id)

	if err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(m.MapToForm()))

}
```
And perhaps another example to demonstrate how to receive user input and store a model.
```Go
func Post(c echo.Context) error {

	f := &models.CatForm{}

	if err := c.Bind(f); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}

	if err := helpers.Validate(f); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ValidationErrors(err))
	}

	m := f.MapToModel()

	if err := m.Save(); err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(m.MapToForm()))

}
```

</details>

<details>
<summary><b>Routers</b></summary>

> Note: This go boilerplate uses [Echo](https://echo.labstack.com). If you're ever in doubt, you can refer to Echo's documentation for more details on what's possible with routers.

This boilerplate is shipped with 3 routers, public, privdate and hidden routers - all of which ollow the same structure and procedure with slight differences in what is registered within each. 

Why have 3 routers? Well, some projects may have public and protected routes, and such use case is straightforward. The latter implements an authentication middleware while the first does not. Attempting to achieve such behaviour within a single router can be tricky, so isolated routers running on different ports are used instead. The third "hidden" router is provided to enable a pattern commonly used to allow one microservice to communicate with another without exposing those routes to the public internet. With that said, wiring those 3 routers can easily be achieved through a different service like Kubernetes or NGINX. 

All routers should go through the following process: 
1. Initialisation 
2. Checking for DevMode and enabling related middlewares 
3. Register common middleware 
4. Register health routes
5. Register security middleware 
6. Register user-defined routes
7. Register error handler 

Have a look at `pkg/api/routers/protectedApi.go` to familiarise yourself with router initialisation process. If you've already created your handlers from the previous section, all you need is to add your new route to this file as such:
```Go
func registerProtectedAPIRoutes() {
	cats := protectedApiRouter.Echo.Group("/cats") // Your new REST resource
	cats.GET("", catsHandlers.Index)               // GET "/cats/" route and handler 
	cats.GET("/:id", catsHandlers.Get)             // GET "/cats/:id" route and handler
	cats.POST("", catsHandlers.Post)               // POST "/cats/" route and handler
	cats.PUT("/:id", catsHandlers.Put)             // PUT "/cats/:id" route and handler
	cats.DELETE("/:id", catsHandlers.Delete)       // DELETE "/cats/:id" route and handler

	// add more routes here ...
}
```

</details>

<details>
<summary><b>Tasks</b></summary>

Tasks is a way to extend the command line CLI without having to go through the trouble of understanding the initialisation process.

To create a new task, simply add the following sample into a new file `./pkg/tasks/<myTask>.go`:

```Go

func init() {
	var t = &Task{
		Name:        "myTask",
		Description: "This is my task",
		RequiredArgs: []string{"key1", "key2"}, // add args here
		Run:         execMyTask,
	}
	Tasks.AddTask(t)
}

func execMyTask(env *TaskEnv, args map[string]string) error {
	// task implementation goes here...
	fmt.Println("My task is executed!")
	return nil
}
```

Tasks are automatically injected with an `env` object that contains the environment. They are also injected with an `args` map containing any values added to the exec command, as long as they're separated by '=' e.g. `key1=value1 key2=value2 key3=value3 `.

You can also set the required arguments in `myTask.RequiredArgs = []string{"key1", "key2"}` to prevent execution until all arguments are provided.

To execute the task above, simply run `go run . task exec myTask` and you should get the "My first task is executed!" message.

</details>

<details>
<summary><b>Error Handling</b></summary>

This boilerplate automates error handling and error responses. 

First, let's talk about logging errors. When using proper logging mechanics and log levels, you can then leave all your logs in the code and have them printed depending on their severity. The package is shipped with the function `helpers.Error()`, a wrapper that's intended to log an error and return it. These logs will only be visible if the `LOG_LEVEL` env var permits. Avoid using `fmt.Println()` at all times, instead, use `logger.Debug()` or if you're within a handler, you can use the `c.Logger().Debug()` helper.

Given that each handler can return an error, the router is configured to handle the error using `pkg/api/handlers/errors/automatedHttpError.go` which will unwrap the error and match it with a list of registered errors under `pkg/api/handlers/errors/errors.go`. Finally, it will construct an error response and respond to that request. 

Validation errors are no different, except they're unwrapped further and sent to the user as individual form inputs so they can be displayed. 

You're encouraged to register and maintain as many errors as you can in the same way. It's useful to have a specific error code mapped to each error, that way we can determine exactly what went wrong in each user flow. 

</details>

<details>
<summary><b>Adding Env Vars & Features</b></summary>

All environment variables reside in `pkg/config/features`. They're categorised within their respective features such as `database.go` or `service.go`. Each env var must have a `mapstructure:` decoration that spells it in caps when parsing an ENV. You can add your own, it's as simple as adding a new line in any of these files, or create your own. 

Below is a sample of `pkg/config/features/service.go`:
```Go
type ServiceConfig struct {
	Host                   string `mapstructure:"HOST"`
	ProtectedApiPort       string `mapstructure:"PROTECTED_API_PORT"`
	PublicApiPort          string `mapstructure:"PUBLIC_API_PORT"`
	HiddenApiPort          string `mapstructure:"HIDDEN_API_PORT"`
	LogLevel               string `mapstructure:"LOG_LEVEL"`
	RequestTimeoutDuration string `mapstructure:"REQUEST_TIMEOUT_DURATION"`
	WatcherSleepInterval   string `mapstructure:"WATCHER_SLEEP_INTERVAL"`
}

var service = &Feature{
	Name:       constants.FEATURE_SERVICE,
	Config:     &ServiceConfig{},
	enabled:    true,
	configured: false,
	ready:      false,
	requirements: []string{
		"Host",
		"ProtectedApiPort",
		"PublicApiPort",
		"HiddenApiPort",
		"LogLevel",
		"RequestTimeoutDuration",
		"WatcherSleepInterval",
	},
}

func init() {
	Features.Add(service)
}
```

From the example above, you can find a type `ServiceConfig` that states what env vars are to be expected. These are automatically read from the environment. Env vars must belong to a feature which can be toggled on or off. A feature can also define which env vars are required for it to start. 

If you wish to disable a feature, you can mention it in the list of `DISABLE_FEATURES` var in run-time. 

Reading the env vars is the job of `pkg/config/envVars.go`. Each config struct must be registered in `envVars.go`. The config struct is then automatically injected to its respective feature after initialisation.

It is possible to set a default value for each variable, this can be done in `pkg/config/envVars.go` under `setDefaults()`.

By the time the CMD calls the Proc, all env vars should have already been read and injected into their features, making them available for the rest of the package. 

</details>


<details>
<summary><b>Folder Structure</b></summary>

### Root Folder Structure

The package is split into 3 directories

| Directory | Description |
| --------- | ----------- |
| `/ci` | Contains all files related to building or deploying the service such as Docker, Docker compose, configuration and K8S files|
| `/cmd` | Contains all available commands |
| `/pkg` | Contains all source code files. This is where you'll be spending most of your time |

### pkg Folder Structure

| Directory | Description |
| --------- | ----------- |
| `/pkg/api` | Everything related to `Echo`, routers and handlers go in here |
| `/pkg/clients` | These are clients used throughout the service. They can be third-party services or simple config providers for workflows |
| `/pkg/config` | Service configuration and environment variable management |
| `/pkg/db` | Everything related to database entities and models, migrations, and seed data |
| `/pkg/proc` | Entry points for all processes |
| `/pkg/tasks` | User-defined tasks available via the command line CLI |
| `/pkg/utils` | General utilities used throughout the package that do not belong to any specific package |


<details>
<summary><b>Package tree view</b></summary>

```
+- pkg
|  +- api
|  |  +- handlers
|  |  |  +- healthz
|  |  |  +- errors
|  |  |  +- cats            <--- example handlers
|  |  |  +-                 <--- additional handlers
|  |  +- helpers
|  |  |  +- helpers.go
|  |  |  +-                 <--- additional helpers
|  |  +- middlewares
|  |  |  +- authentication.go
|  |  |  +- authorization.go
|  |  |  +-                 <--- additional middleware
|  |  +- routers
|  |  |  +- router.go
|  |  |  +- hiddenApi.go
|  |  |  +- protectedApi.go
|  |  |  +- publicApi.go
|  |  |  +-                 <--- additional routers
|  |  +- api.go
|  +- clients
|  |  +- db
|  |  +- fhttp
|  |  +- keto
|  |  +- kratos
|  |  +- logger
|  |  +- redis
|  +- config
|  |  +- autoEnv.go
|  |  +- config.go
|  |  +- feature.go
|  |  +- features.go
|  |  +- flags.go
|  |  +- service.go
|  +- db
|  |  +- migrations
|  |  |  +- migrations.go   <--- list of migrations to run, be sure to add yours here
|  |  |  +-                 <--- additional migrations
|  |  +- models
|  |  |  +- models.go
|  |  |  +- forms.go
|  |  |  +- cat.go          <--- example model
|  |  |  +- catForm.go      <--- example form
|  |  |  +-                 <--- additional routers
|  |  +- seeds
|  |  |  +- seeds.go        <--- list of seeds to run, be sure to add yours here
|  |  |  +-                 <--- additional seeds
|  +- proc                  <--- entry point to all processes
|  |  +- proc.go
|  |  +- hiddenApi.go
|  |  +- protectedApi.go
|  |  +- publicApi.go
|  |  +- watcher.go
|  +- tasks
|  |  +- myFirstTask.go
|  |  +-                    <--- additional tasks
|  +- utils
|  |  +- constants
|  |  |  +- constants.go    <--- all literal values
|  |  +- init.go
|  |  +- utils.go           <--- reusable functions that don't belong anywhere else
+- go.mod
+- main.go

```
</details>

</details>

</details>

<details>
<summary><b>Dependencies</b></summary>

This package is purely written in Go, which helps with dependency management. All dependencies can be easily installed using the `go get` command. 

There are only 2 optional dependencies that can be installed separately. The first is [Air](https://github.com/cosmtrek/air) used for [live-reload](#live-reload), and the other is [Docker](https://www.docker.com/products/docker-desktop/).

List of run-time dependencies:

- [GoLang](https://go.dev) v1.20
- [Cobra](https://github.com/spf13/cobra) v1.8.0
- [Viper](https://github.com/spf13/viper) v1.18.2
- [Echo](https://echo.labstack.com) v4.11.3
- [Gorm](https://gorm.io) v1.25.6
- [MySQL](https://github.com/go-gorm/mysql) v1.5.2
- [PostgreSQL](https://github.com/go-gorm/postgres) v1.25.6
- [SQLite](https://github.com/go-gorm/sqlite) v1.5.4
- [GoMigrate](https://github.com/go-gormigrate/gormigrate) v2.1.1
- [Playground Validator](https://github.com/go-playground/validator) v10.17.0
- [Ory Kratos](https://github.com/ory/kratos-client-go) v1.0.0
- [Ory Keto](https://github.com/ory/keto-client-go) v0.5.2
- [Redis](https://github.com/redis/go-redis)

List of development dependencies: 

- [Air](https://github.com/cosmtrek/air)       # Optional for live-reload
- [Docker](https://www.docker.com/products/docker-desktop/)

</details>


<details open>
<summary><b>Known Issues</b></summary>

- Gorm v1.25.6 and v1.25.7 are known to cause issues with PostgreSQL database. If you experience an error `this driver does not support LastInsertID()`, try downgrading Gorm to v1.25.5

</details>


<details open>
<summary><b>Roadmap</b></summary>

- [ ] Enhanced routers
- [ ] Feature toggle for [`hiddenApi`, `protectedApi`, `publicApi`]
- [ ] Keto client
- [ ] Redis client
- [ ] Forward HTTP client
- [ ] Enhanced error handling
- [ ] Quick start examples
	- [ ] Example with Kratos for authentication
	- [ ] Example with Keto for authorisation
- [ ] Code cleanup and in-line documentation
- [ ] Swagger integration
- [ ] Postman collection
- [ ] More documentation
- [ ] Tests
- [ ] Github Actions 
- [ ] Landing page / Website 

</details>


### Contribution

Feel free to start a new discussion, submit a new PR, make a feature request or etc.. If you would like to join the team, reach out to us on Discord. We are always looking for contributors!

### Contacts

- [Join our Discord channel](https://discord.gg/Q27kgPVub7)
- [Dexter Codo](mailto:dexter@dexterexplains.com)
