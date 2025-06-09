@echo off
setlocal enabledelayedexpansion

set "year=%~1"
set "result="

set /a "mod4=year %% 4"
set /a "mod100=year %% 100"
set /a "mod400=year %% 400"

if %mod4%==0 (
    if %mod100%==0 (
        if %mod400%==0 (
            set "result=1"
        ) else (
            set "result=0"
        )
    ) else (
        set "result=1"
    )
) else (
    set "result=0"
)

echo %result%
