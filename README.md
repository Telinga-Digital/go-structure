# Go Structure

This repository is created a reference of golang project structure.

## Prerequisite

* Golang
* Nodejs

## How to Run

* git clone git@github.com:endrureza/go-structure.git
* cd go-structure
* yarn install
* yarn dev

## Application Structure

### `/app`

All API logic, model relation and middleware should go to this folder

### `/cli`

Any command line script should go to this folder

### `/config`

All configuration regarding application should go to this folder

### `/routes`

All endpoint or routing should go to this folder

### `/services`

All third party services configuration or setting should go to this folder

### `/tests`

All test mechanism should go to this folder

### `/utils`

All imported library that is not exposed to outside of package should go to this folder

### `main.go`

This file will be the main to execute / run application

### `apidoc.json`

All route documentation should go to this file

## Authors

* Endru Reza ([Linkedin](https://linkedin.com/in/endrureza))