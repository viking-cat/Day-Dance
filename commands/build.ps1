#!/bin/bash

function bashMain {
    echo "command script says hello!"
}

"bashMain" # This runs the bash script, in PS it's just a string
"exit"     # This quits the bash version, in PS it's just a string

# PowerShell code here
# --------------------

function powershellMain {
    Write-Host "Building Service" -ForegroundColor red
    Write-Host "PowerShell Version: " $PSVersionTable.PSVersion
    $gov = go version
    Write-Host "Golang Version: " $gov
}

powershellMain

# Documentation
# https://stackoverflow.com/questions/39421131/is-it-possible-to-write-one-script-that-runs-in-bash-shell-and-powershell