@echo off
cd C:/Users/brcon/GolandProjects/DrugRecord/DrugRecord/
go build
start chrome.exe "http://localhost:80"
call  C:/Users/brcon/GolandProjects/DrugRecord/DrugRecord/DrugRecord.exe