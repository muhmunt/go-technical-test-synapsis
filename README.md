# go-technical-test-synapsis
backend developer technical test at synapsis

# Technical Test Description:
Test Overview:

Overview
The purpose of this project is to develop an online store application focusing solely on the backend (RESTful API). The API is built using Go-lang to meet the minimum viable product (MVP) criteria. This README provides instructions on how to set up, run, and use the application.

# Usage
Inside docker this project will running on default port localhost:8084 with the default database on same container in port 3306

# Endpoint
API Documentation are attached in 
**postman folder** and you could import to other device to see the detailed body and response message.

# Installation

Follow these steps to install and set up your project locally. Make sure you meet the prerequisites before proceeding.

### Prerequisites

Before you begin, ensure you have the following prerequisites installed on your system:

- [Go](https://golang.org/dl/): This project is built with Go, so you need to have Go installed on your system.
- [docker]: docker is required for running inside a container with mysql.

### Step-by-Step Installation

1. **Setup the Repository:**

   Copy an .env-example file edit to .env and configure it as explained in the Configuration section

   ```bash
   MYSQL_CONNECTION="synapsis:secret@tcp(db:3306)/synapsis_test"
   PORT=8484

2. **Running project on docker docker:**

   For running with docker i've a command with Makefile command as a :

   ```bash
   make docker-build
   make docker-up

- **Running project without docker**

1. **Copy an .env-example file edit to .env and configure it as explained in the Configuration section**

- import the database inside a db folder named as synapsis_test.sql

2. Run **go mod tidy** to download go modules

3. **Run the project**
   ```bash
   go run main.go





