@echo off
SET scripts_dir=%cd%
SET working_dir=%scripts_dir%\..
echo %working_dir%
cd %working_dir%
go build
md %working_dir%\DrugRecord\web\assets
md %working_dir%\DrugRecord\configs
md %working_dir%\DrugRecord\Prerequisites
md %working_dir%\DrugRecord\web\templates
md %working_dir%\DrugRecord\log
md %working_dir%\DrugRecord\scripts
md %working_dir%\DrugRecord\backups
copy prerequisites\* DrugRecord\Prerequisites\
cd web\assets
copy * ..\..\DrugRecord\web\assets
cd ..\templates
copy * ..\..\DrugRecord\web\templates
del %working_dir%\DrugRecord\web\assets\homeStyle.css
del %working_dir%\DrugRecord\web\assets\loginStyle.css
del %working_dir%\DrugRecord\web\assets\registerStyle.css
cd %working_dir%
copy .\DrugRecord.exe DrugRecord\cll.exe
copy .\scripts\start.cmd DrugRecord\scripts\start.cmd
copy .\scripts\setup_db.cmd DrugRecord\scripts\setup_db.cmd
copy .\scripts\restore.cmd DrugRecord\scripts\restore.cmd
copy .\configs\configuration.ini DrugRecord\configs\configuration.ini
del %working_dir%\DrugRecord.exe
makensis.exe scripts\installer.nsi
cd %scripts_dir%