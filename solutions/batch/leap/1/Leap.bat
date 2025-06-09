@echo off
setlocal enabledelayedexpansion

set "year=%~1"
set "result="

REM Your code goes here

if [ $year % 4 -eq 0 && $year % 100 -ne 0 || $year % 400 -eq 0 ] then
result="leap"
fi

echo %result%
