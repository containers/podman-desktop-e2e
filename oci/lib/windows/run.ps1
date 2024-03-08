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
    [Parameter(HelpMessage='path to save all results. This folder will be copied out from target host')]
    $resultsFolder,
    [Parameter(HelpMessage='Check if wsl is installed if not it will install, default true.')]
    $wslInstallFix="true"
)

function Install-PD {
    Write-Host "downloading Podman Desktop"
    cd $targetFolder
    curl.exe -L $pdUrl -o pd.exe
    cd ..
}

# Add target to path to be able to run pd-e2e
$env:PATH="$env:PATH;$env:HOME\$targetFolder;"
# Create results folder
New-Item -ItemType Directory -Path "$env:HOME\$targetFolder\$resultsFolder" -Force

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
$pp=Get-Process "Podman Desktop"
if ($pp) {
    $pdpid=$pp | Select Id -ExpandProperty Id | Select-Object -first 1
    Stop-Process -ID $pdpid -Force
}
if (!$pdPath) {
    Install-PD
    $pdPath="$env:HOME\$targetFolder\pd.exe"
}
Start-Process powershell -verb runas -ArgumentList "Set-NetFirewallProfile -Profile Domain,Public,Private -Enabled False" -wait
pd-e2e.exe --user-password $userPassword --junit-filename $junitResultsFilename --pd-path "$pdPath" --screenshotspath "."
Start-Process powershell -verb runas -ArgumentList "Set-NetFirewallProfile -Profile Domain,Public,Private -Enabled True" -wait

#Workaround 
mv junit_report.xml "$targetFolder\$resultsFolder\$junitResultsFilename"
rm junit_report.xml
mv *.png "$targetFolder\$resultsFolder"
rm *.png

# Kill PD testing instance
$pdpid=Get-Process "Podman Desktop" | Select Id -ExpandProperty Id | Select-Object -first 1
Stop-Process -ID $pdpid -Force