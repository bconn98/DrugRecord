
; installer.nsi
;
; This script is perhaps one of the simplest NSIs you can make. All of the
; optional settings are left to their default settings. The installer simply
; prompts the user asking them where to install, and drops a copy of "MyProg.exe"
; there.

;--------------------------------

!include "nsDialogs.nsh"
!include WinMessages.nsh

; The name of the installer
Name "C-ll Installer"

; The file to write
OutFile "cll_installer.exe"

!include "Sections.nsh"
!include nsDialogs.nsh

; The default installation directory
InstallDir $DESKTOP\C-ll

; The text to prompt the user to enter a directory
DirText "This will install C-ll on your computer. Choose a directory"


Var Dialog
Var TextPgDir

;--------------------------------
Page Directory
Page Instfiles

Section !Required
  SectionIn RO
        ; Set install path
        SetOutPath $INSTDIR

        ; Put files there
        File /r "A:\Documents\JetBrains\GolandProjects\DrugRecord\DrugRecord\*"
SectionEnd

Page custom LogoPage

Function LogoPage
    nsDialogs::Create 1018
    Pop $Dialog

    ${If} $Dialog == error
        Abort
    ${EndIf}

    ${NSD_CreateGroupBox} 5% 86u 90% 34u "Select your logo, or next to continue"
    Pop $0

        ${NSD_CreateDirRequest} 15% 100u 49% 12u "$INSTDIR\web\assets\logo.png"
        Pop $TextPgDir

        ${NSD_CreateBrowseButton} 65% 100u 20% 12u "Browse..."
        Pop $0

        ${NSD_OnClick} $0 OnLogoDirBrowse

    GetDlgItem $0 $HWNDPARENT 1
    SendMessage $0 ${WM_SETTEXT} 0 `STR:$(^NextBtn)`
    EnableWindow $0 1
    nsDialogs::Show
FunctionEnd

Function OnLogoDirBrowse
    ${NSD_GetText} $TextPgDir $0
    nsDialogs::SelectFileDialog "Select Logo" "$0"
    Pop $0
    ${If} $0 != error
        ${NSD_SetText} $TextPgDir "$0"
    ${EndIf}
    CopyFiles $0 $INSTDIR\web\assets\logo.png
FunctionEnd

Page custom RestorePage

Function RestorePage
    nsDialogs::Create 1018
    Pop $Dialog

    ${If} $Dialog == error
        Abort
    ${EndIf}

    ${NSD_CreateGroupBox} 5% 86u 90% 34u "Select your backup file or next to continue"
    Pop $0

        ${NSD_CreateDirRequest} 15% 100u 49% 12u "$INSTDIR\backups\"
        Pop $TextPgDir

        ${NSD_CreateBrowseButton} 65% 100u 20% 12u "Browse..."
        Pop $0

        ${NSD_OnClick} $0 OnRestoreDirBrowse

    GetDlgItem $0 $HWNDPARENT 1
    SendMessage $0 ${WM_SETTEXT} 0 `STR:$(^CloseBtn)`

    EnableWindow $0 1
    nsDialogs::Show
FunctionEnd

Function OnRestoreDirBrowse
    ${NSD_GetText} $TextPgDir $0
    nsDialogs::SelectFileDialog "Select Backup" "$0"
    Pop $0
    ${If} $0 != error
        ${NSD_SetText} $TextPgDir "$0"
    ${EndIf}
    ExecWait '$INSTDIR\scripts\restore.cmd $0'
FunctionEnd

; This is save database
Section Uninstall
   ExecWait '$INSTDIR\scripts\backup.cmd "$INSTDIR\backups"'
SectionEnd

Section
  ; Install Prereq
  IfFileExists $PROGRAMFILES64\PostgreSQL endPostgreSQL beginPostgreSQL
  Goto endPostgreSQL
  beginPostgreSQL:
    ExecWait "$INSTDIR\Prerequisites\postgresql-12.3-1-windows-x64.exe --mode unattended  --servicepassword Zoo123"
  endPostgreSQL:

  ; Create uninstaller
  WriteUninstaller "$INSTDIR\backup.exe"
SectionEnd
