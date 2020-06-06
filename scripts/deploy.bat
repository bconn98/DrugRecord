rem @echo off
cd A:\Documents\JetBrains\GolandProjects\DrugRecord\
go build
mkdir A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\assets
mkdir A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\Prerequisites
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\web\templates
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\log
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
copy .\scripts\start.bat DrugRecord\start.bat