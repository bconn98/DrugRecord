
; installer.nsi
;
; This script is perhaps one of the simplest NSIs you can make. All of the
; optional settings are left to their default settings. The installer simply
; prompts the user asking them where to install, and drops a copy of "MyProg.exe"
; there.

; Save DB option which finds the database of drugrecord, compresses it to
; a desired location.

;--------------------------------

; The name of the installer
Name "C-ll Installer"

; The file to write
OutFile "c-ll.exe"

!include "Sections.nsh"
!include nsDialogs.nsh

; The default installation directory
InstallDir $DESKTOP\C-ll

; The text to prompt the user to enter a directory
DirText "This will install C-ll on your computer. Choose a directory"

;--------------------------------
Page Directory
Page Components
Page Instfiles
Page Custom Logo
Page Custom CloseWindow

Section "Program Required Files (Required)"
  SectionIn RO
    # common files here
SectionEnd

Section "PostgreSQL" SEC_POSTGRESQL
  SectionIn 1
      IfFileExists $PROGRAMFILES64\PostgreSQL endPostgreSQL beginPostgreSQL
      Goto endPostgreSQL
      beginPostgreSQL:
        ExecWait "$INSTDIR\\Prerequisites\postgresql-12.3-1-windows-x64.exe --mode unattended  --servicepassword Zoo123"
      endPostgreSQL:
SectionEnd

Function .onSelChange
  !insertmacro StartRadioButtons $1
   !insertmacro RadioButton ${SEC_POSTGRESQL}
  !insertmacro EndRadioButtons
FunctionEnd

Section
; Set output path to the installation directory.
SetOutPath $INSTDIR

; Put file there
File /r "A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\*"

SectionEnd ; end the section

Function Logo
  nsDialogs::Create 1018
  ${NSD_CreateLabel} 0 0 100% 20u "Press next to select a new logo. If not desired, close file explorer window."
  nsDialogs::Show
  nsDialogs::SelectFileDialog open "$DOCUMENTS\" ".png files|*.png"
  Delete $INSTDIR\web\logo.png
  Pop $0
  CopyFiles $0 $INSTDIR\web\assets\logo.png
FunctionEnd

Function CloseWindow
  nsDialogs::Create 1018
  ${NSD_CreateLabel} 0 0 100% 20u "Installation complete"
  nsDialogs::Show
FunctionEnd

