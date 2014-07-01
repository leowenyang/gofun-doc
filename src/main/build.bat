@echo off
REM set environment var
set GOPATH=GOPATH;d:\Person\Golang\gofun-doc
set GOBIN=d:\Person\Golang\gofun-doc\bin

REM compile package
go build 

REM compile main package
go build gofun-doc.go

REM install 
go install main 
