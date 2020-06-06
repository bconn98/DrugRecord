rem @echo off
cd A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\
rem go build
mkdir A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\DrugRecord\web\assets
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\DrugRecord\web\templates
md A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\DrugRecord\log
cd web\assets
copy * ..\..\DrugRecord\web\assets
cd ..\templates
copy * ..\..\DrugRecord\web\templates
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\DrugRecord\web\assets\homeStyle.css
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\DrugRecord\web\assets\loginStyle.css
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\DrugRecord\web\assets\registerStyle.css
del A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\DrugRecord\web\templates\writeExcel.html
cd A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\
xcopy .\DrugRecord.exe DrugRecord\
xcopy .\web\start.bat DrugRecord\