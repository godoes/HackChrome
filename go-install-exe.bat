@echo off
color 07
title ������װ 64 λ Windows ϵͳ��ִ�г����
:: file-encoding=GBK
rem by liutianqi

cd /d %~dp0 & cd & echo.

set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64

echo ���ڹ�����װ HackChrome.exe...
go build -o %GOPATH%/bin/HackChrome.exe
echo.

color 0f
set now=%date:~0,4%-%date:~5,2%-%date:~8,2% %time:~0,2%:%time:~3,2%:%time:~6,2%.%time:~9,3%
echo [%now%] ������װ��ɣ� & echo.
if not "%1%" == "NoPause" (
  pause
)
