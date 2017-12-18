@ECHO OFF
ECHO API SERVICE
SETLOCAL
TITLE Run "GO"
COLOR 0f

:: start
REM Setting evernoument
SET GOPATH=%CD%
SET GOBIN=%CD%\bin
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
CLS

rem UNIX
rem Created server for Unix System
rem SET GOOS=linux
rem SET GOARCH=amd64
rem SET CGO_ENABLED=0

REM Получение библиотек для базы Gorethinkdb
rem go get -u github.com/dancannon/gorethink
rem go get -u github.com/googollee/go-socket.io
rem go get  github.com/op/go-logging

rem Start database
rem start c:/work/db/rethinkdb.exe

REM Build process
go build -o service.exe
service.exe
pause
