@echo off
set GOPATH=%cd%
go build -o %GOPATH%\bin\cowsay.exe %GOPATH%\src\main.go