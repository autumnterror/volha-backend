@echo off
setlocal enabledelayedexpansion

set "CONTAINER_NAME=productsdb"
set "DB_NAME=productsdb"
set "DB_USER=postgres"

set "LATEST_DUMP="

for /f "delims=" %%F in ('dir /b /a:-d /o:-n ".\dump.*.sql" 2^>nul') do (
    set "LATEST_DUMP=%%F"
    goto found
)

:found

if "%LATEST_DUMP%"=="" (
    echo not found dumps
    exit /b 1
)

echo use dump: %LATEST_DUMP%

REM ---- Восстановление дампа ----
docker exec -i %CONTAINER_NAME% psql -U %DB_USER% -d %DB_NAME% < ".\%LATEST_DUMP%"

echo restore success.
