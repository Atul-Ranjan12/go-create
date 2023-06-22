# go-create

## Introduction 
This repository sets up the basic go server with all the requirements for quick web development using Golang. The project is also
organized into folders with boiler plate code for ease. Please go through the Project File Structure for information on the different 
files and its intended use.

## Getting Started:
### Demo:

https://github.com/Atul-Ranjan12/go-create/assets/33323817/2c501c04-c1f8-4ec2-9dcf-74c23a06d489

### Get Started Here:
**Step 1** Clone the Repository 
```
git clone github.com/Atul-Ranjan12/go-create <<YOUR_FOLDER_NAME>>
```
**Step 2** Change Directory 
```
cd <<YOUR_FOLDER_NAME>>
```
**Step 3** Allow Bash Files to Run on the Terminal
```
chmod +x ./run.sh
```
```
chmod +x ./setup.sh
```
**Step 4** Run the setup file to install required dependencies
Run the setup file to install required dependencies
```
./setup.sh
```
#### Step 5
You are good to go!
Go to cmd -> web -> main.go, make changes to your dsn string for the appropriate connection. To run the project:
```
./run.sh
```

#### Step 6
Change the go mod file to your preference
```
go mod edit -module <<YOUR_NEW_NAME>>
```

## Project File Organization

1. _main.go_ This file contains the driver code to run the application.
2. _middleware.go_: This file sets up the required middleware 
3. _routes.go_: This file manages the required routes for the website
4. _config.go_: Describes the App Configurations 
5. _driver.go_: Handles Database Connections
6. _forms_: Contains form functions
7. _handlers.go_: Contains all the handler fucntions for all the website routes
8. _helpers_: Contains helper fucntions for the website 
9. _models.go_: Contains Model structures for the website 
10. _render.go_: Handles rendering of templates 
11. _static_: Contains static contents of the website
12. _templates_: Contains all the go templates. render.go renders all the templates in this folder

## Dependencies
1. github.com/alexedwards/scs/v2 : Session Manager 
2. github.com/go-chi/chi/v5L : Routes
3. github.com/justinas/nosurf : CSRF Tokens
4. github.com/jackc/pgx/v4 : Database Connection
