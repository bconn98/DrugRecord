
; installer.nsi
;
; This script is perhaps one of the simplest NSIs you can make. All of the
; optional settings are left to their default settings. The installer simply
; prompts the user asking them where to install, and drops a copy of "MyProg.exe"
; there.

;--------------------------------

; The name of the installer
Name "C-ll Installer"

; The file to write
OutFile "c-ll.exe"

; The default installation directory
InstallDir $DESKTOP\C-ll

; The text to prompt the user to enter a directory
DirText "This will install C-ll on your computer. Choose a directory"

;--------------------------------

Section -Prerequisites
; Set output path to the installation directory.
  SetOutPath $INSTDIR\Prerequisites

  IfFileExists $PROGRAMFILES64\PostgreSQL endPostgreSQL beginPostgreSQL
  Goto endPostgreSQL
  beginPostgreSQL:
  MessageBox MB_YESNO "Your system is missing PostgreSQL would you like to install it? Please set the password to 'Zoo123" /SD IDYES IDNO endPostgreSQL
    File "..\Prerequisites\postgresql-12.3-1-windows-x64.exe"
    ExecWait "$INSTDIR\Prerequisites\postgresql-12.3-1-windows-x64.exe"
    Goto endPostgreSQL
  endPostgreSQL:

SectionEnd

; The stuff to install
Section "" ;No components page, name is not important

; Set output path to the installation directory.
SetOutPath $INSTDIR

; Put file there
File /r "A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\*"

SectionEnd ; end the section

