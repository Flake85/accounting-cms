## Name:
CLI flags

---
## Description:
use CLI flags to set env variables

---
#### Subtasks:
- [x] Add Makefile to root of project
    - [x] set variables for flags
    - [x] set a target for golang
    - [x] write the go run command with the flags using the set variables
- [x] Create/Retrieve flags
    - [x] create default flags (Makefile golang command will over-ride)
    - [x] the database connection string (main.go) should be a function returning the database connection string using the flags either sent by makefile command, manual input command or default.  
- [x] Create frontend env file to pass variables (rest api url)