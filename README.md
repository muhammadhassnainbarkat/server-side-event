# Full Stack Application: Vue.js Frontend, Python Flask Backend, and Go SSE Server

This project is a full-stack application that involves a Vue.js frontend, a Python Flask backend, and a Go server for handling Server-Sent Events (SSE). The application allows clients to subscribe to SSE and receive real-time updates.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Vue.js Frontend](#vuejs-frontend)
  - [Installation](#installation)
  - [Running the Development Server](#running-the-development-server)
- [Python Flask Backend](#python-flask-backend)
  - [Installation](#installation-1)
  - [Running the Flask Server](#running-the-flask-server)
- [Go Backend](#go-backend)
  - [Installation](#installation-2)
  - [Running the Go Server](#running-the-go-server)

## Prerequisites

Ensure you have the following installed:

- [Node.js](https://nodejs.org/) (for Vue.js)
- [Python 3.x](https://www.python.org/) (for Flask)
- [Go 1.x](https://golang.org/) (for Go SSE server)
- [Git](https://git-scm.com/) (for version control)

## Project Structure

```plaintext
project-root/
├── sse-frontend-vue/               # Vue.js frontend
├── sse-backend-python/             # Python Flask backend
└── sse-backend-go/                 # Go SSE server

```

## VueJs Frontend
### Installation
```bash
cd frontend
npm install
```


### Running the Development Server
```bash
npm run dev
```


## Python Flask Backend
### Installation
```bash

cd sss-python-backend
# activate enviroment
pip install -r requirments.txt
```
Note: If requirements.txt does not exist, create it and add:

```plaintext
Copy code
Flask
flask-cors
```


### Running the Flask Server
```python
python app.py
```


## Go Backend
### Installation
```bash
cd sss-backend-go
go mod init
go build
```


### Running the Go Server
```go
go run main.go
```
