@echo off
cd A:\Documents\JetBrains\GolandProjects\DrugRecord\
go build
mkdir A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\assets
mkdir A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\Prerequisites
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\templates
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\log
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\scripts
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\backups
copy prerequisites\* DrugRecord\Prerequisites\
cd web\assets
copy * ..\..\DrugRecord\web\assets
cd ..\templates
copy * ..\..\DrugRecord\web\templates
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\assets\homeStyle.css
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\assets\loginStyle.css
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\assets\registerStyle.css
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\templates\writeExcel.html
cd A:\Documents\JetBrains\GolandProjects\DrugRecord\
copy .\DrugRecord.exe DrugRecord\cll.exe
copy .\scripts\start.cmd DrugRecord\scripts\start.cmd
copy .\scripts\setup_db.cmd DrugRecord\scripts\setup_db.cmd
copy .\scripts\restore.cmd DrugRecord\scripts\restore.cmd
copy .\scripts\backup.cmd DrugRecord\scripts\backup.cmd