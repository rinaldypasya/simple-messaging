# A Simple Messaging Application written in Go

## Requirement and Dependencies

- [websocket](https://github.com/gorilla/websocket) - `go get github.com/gorilla/websocket`
- [gin](https://github.com/gin-gonic/gin) - `go get github.com/gin-gonic/gin`
- [gin-cors](https://github.com/itsjamie/gin-cors) - `go get github.com/itsjamie/gin-cors` 
- [mysql-driver](https://github.com/go-sql-driver/mysql) - `go get github.com/go-sql-driver/mysql`  
- [gorm](https://github.com/jinzhu/gorm) - `go get github.com/jinzhu/gorm`

## Installation Guide
Before installing and running this app, make sure you have install [Go](https://golang.org/doc/install). After Go is installed, Follow these steps:

### Linux/Mac
- Set your [gopath](https://til.codes/how-do-i-set-the-gopath-environment-variable-on-ubuntu/) somewhere in your filesystem, preferably in /home.
- After setting up gopath, create 3 main folder inside gopath. For example, if your gopath is /home/user/go, then create these 3 folder under that go folder:
    - bin/
    - pkg/
    - src/
- Then, open terminal inside your gopath(e.g /home/user/go). Then get this app by using commands:
    >`go get github.com/rinaldypasya/simple-messaging`
- The codebase for the app will be downloaded inside `src/` folder inside your gopath.
- After codebase downloaded, change directory to that app -> e.g 
    >`cd $GOPATH/src/github.com/rinaldypasya/simple-messaging`
- Install/Build the app from source to build the executable and packages object by using this command
    >`make build` (assuming you're already inside the app folder path)

- After finishing installing/building run the app by calling the name of app folder(where main package reside), in this case, if you're inside the app path, just call the name of the app:
    >`./simple-messaging`
    
- If program installed and run successfully, the routing informations will be showed on the terminal showing all API endpoints registered inside router. The port used is :8000, so access it by calling 

    >`localhost:8000/<endpoints>`

### Windows

For installation on Windows, after installing the Go you're required to install **MinGW-64 GCC** because some of depedencies depend and compiled only by gcc 64 bit version, and on Windows you have to use [MinGW64](https://sourceforge.net/projects/mingw-w64/). Remember to download, install, and activate the **64 bit version**.

After installing Go and MinGW64, set [gopath in windows](https://github.com/golang/go/wiki/SettingGOPATH#windows).

After setting up the gopath, make sure the gopath set successfully by calling this command in command prompt:
    >go env

If the gopath set and match with the one you set, the rest of the steps are the same with the Linux/Mac. Follow them one by one and you're all set.

### Rebuilding and Cleaning

Please typing command in project root folder:

- Clean

    >`make clean`

- Rebuild and clean

    >`make rebuild`

### Endpoints
- Send Message

    >Send POST: `/message/send`

- Get All Messages

    >GetMessages GET: `/messages`

- Websocket for real-time API    

    >Websocket GET: `/ws`

---
That's all. Thank you and if you have any problems or questions, reach me an email at **rinaldypasya@gmail.com**.