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
    $userPassword
)

function Install-PD {
    Write-Host "downloading Podman Desktop"
    cd $targetFolder
    curl.exe -L $pdUrl -o pd.exe
    cd ..
}

# Run e2e
$env:PATH="$env:PATH;$env:HOME\$targetFolder;"

# Force install just in case
wsl -l -v
$installed=$?

if (!$installed) {
    Write-Host "installing wsl2"
    wsl --install  
}

if (!$pdPath)
{
    Install-PD
    pd-e2e.exe --user-password $userPassword --junit-filename $junitResultsFilename --pd-path "$env:HOME\$targetFolder\pd.exe"
} else {
    pd-e2e.exe --user-password $userPassword --junit-filename $junitResultsFilename --pd-path "$pdPath\pd.exe"
}

#Workaround 
mv junit_report.xml "$targetFolder\$junitResultsFilename"