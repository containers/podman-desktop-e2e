param(
    [Parameter(HelpMessage='url to download the exe for podman desktop, in case we want to test an specific build')]
    $pdUrl="",
    [Parameter(HelpMessage='path for the exe for podman desktop to be tested')]
    $pdPath="",
    [Parameter(Mandatory,HelpMessage='folder on target host where assets are copied')]
    $targetFolder,
    [Parameter(Mandatory,HelpMessage='junit results filename')]
    $junitResultsFilename,
    [Parameter(HelpMessage='user password to run privileged commands (run installers)')]
    $userPassword,
    [Parameter(HelpMessage='Check if wsl is installed if not it will install, default true.')]
    $wslInstallFix="true"
)

function Install-PD {
    Write-Host "downloading Podman Desktop"
    cd $targetFolder
    curl.exe -L $pdUrl -o pd.exe
    cd ..
}

# Run e2e
$env:PATH="$env:PATH;$env:HOME\$targetFolder;"

# Force installation of WSL if needed
if ( $wslInstallFix -match 'true' )
{
    wsl -l -v
    $installed=$?

    if (!$installed) {
        Write-Host "Installing WSL, setting default version 2"
        wsl --set-default-version 2
        wsl --install --no-distribution
    }
}

# Ensure no previous pd is running
$pdpid=Get-Process "Podman Desktop" | Select Id -ExpandProperty Id | Select-Object -first 1
Stop-Process -ID $pdpid -Force

if (!$pdPath)
{
    Install-PD
    pd-e2e.exe --user-password $userPassword --junit-filename $junitResultsFilename --pd-path "$env:HOME\$targetFolder\pd.exe"
} else {
    pd-e2e.exe --user-password $userPassword --junit-filename $junitResultsFilename --pd-path "$pdPath"
}

#Workaround 
mv junit_report.xml "$targetFolder\$junitResultsFilename"