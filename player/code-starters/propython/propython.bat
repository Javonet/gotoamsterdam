@echo off
setlocal enabledelayedexpansion

REM Set the base folder name and path
set "base_folder_name=propython"
set "base_folder=%~dp0%base_folder_name%_0"

REM Get the latest folder number
set "latest_number=0"
for /d %%F in ("%~dp0%base_folder_name%_*") do (
    set "folder_name=%%~nxF"
    for /f "tokens=2 delims=_" %%A in ("!folder_name!") do (
        set "number=%%A"
        if !number! gtr !latest_number! set "latest_number=!number!"
    )
)

REM Increment the latest folder number
set /a new_number=latest_number+1
set "new_folder_name=%base_folder_name%_%new_number%"
set "new_folder_path=%~dp0%new_folder_name%"

REM Copy the base folder to the new folder
xcopy "%base_folder%" "%new_folder_path%" /e /i /h /k /y

REM Copy the Robot client file
copy ..\..\websocket-clients\dotnet\bin\Debug\net8.0\RobotConnector.dll "%new_folder_path%"

REM Open the new folder in Visual Studio Code
cd "%new_folder_path%"
code .

endlocal
