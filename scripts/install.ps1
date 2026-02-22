#Requires -RunAsAdministrator
$ErrorActionPreference = "Stop"

# Force UTF-8 output for ASCII + emojis
[Console]::OutputEncoding = [System.Text.UTF8Encoding]::UTF8

# ---------------- CONFIG ----------------
$Repo = "Saisathvik94/codemaxx"
$Binary = "codemaxx.exe"
$InstallDir = "$env:ProgramFiles\codemaxx"
$TempDir = Join-Path $env:TEMP "codemaxx-install"

# ---------------- FUNCTIONS ----------------
function Show-ASCII {
@"
                           /$$                                                      
                          | $$                                                      
  /$$$$$$$  /$$$$$$   /$$$$$$$  /$$$$$$  /$$$$$$/$$$$   /$$$$$$  /$$   /$$ /$$   /$$
 /$$_____/ /$$__  $$ /$$__  $$ /$$__  $$| $$_  $$_  $$ |____  $$|  $$ /$$/|  $$ /$$/
| $$      | $$  \ $$| $$  | $$| $$$$$$$$| $$ \ $$ \ $$  /$$$$$$$ \  $$$$/  \  $$$$/ 
| $$      | $$  | $$| $$  | $$| $$_____/| $$ | $$ | $$ /$$__  $$  >$$  $$   >$$  $$ 
|  $$$$$$$|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$ | $$ | $$|  $$$$$$$ /$$/\  $$ /$$/\  $$
 \_______/ \______/  \_______/ \_______/|__/ |__/ |__/ \_______/|__/  \__/|__/  \__/
"@
}

function Ensure-Admin {
    $currentUser = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = New-Object Security.Principal.WindowsPrincipal($currentUser)
    if (-not $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
        Write-Host "❌ Please run this script as Administrator!" -ForegroundColor Red
        exit 1
    }
}

function Get-LatestVersion {
    Write-Host "🔍 Fetching latest release..." -ForegroundColor Yellow
    $Latest = Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/latest"
    return $Latest.tag_name
}

function Download-CodeMaxx {
    param($Version)
    $ZipName = "codemaxx_${Version}_windows_amd64.zip"
    $Url = "https://github.com/$Repo/releases/download/$Version/$ZipName"

    Write-Host "📦 Downloading $ZipName ..." -ForegroundColor Yellow

    if (-Not (Test-Path $TempDir)) {
        New-Item -ItemType Directory -Path $TempDir | Out-Null
    }

    $ZipPath = Join-Path $TempDir $ZipName
    Invoke-WebRequest -Uri $Url -OutFile $ZipPath
    return $ZipPath
}

function Install-CodeMaxx {
    param($ZipPath)

    Write-Host "📂 Extracting archive ..." -ForegroundColor Yellow
    Expand-Archive -LiteralPath $ZipPath -DestinationPath $TempDir -Force

    if (Test-Path "$InstallDir\$Binary") {
        Write-Host "🧹 Removing existing CodeMaxx ..." -ForegroundColor Yellow
        Remove-Item "$InstallDir\$Binary" -Force
    }

    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
    Move-Item (Join-Path $TempDir $Binary) $InstallDir -Force

    # Add to PATH if not already present
    $currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
    if ($currentPath -notlike "*$InstallDir*") {
        [Environment]::SetEnvironmentVariable(
            "Path",
            "$currentPath;$InstallDir",
            "Machine"
        )
    }

    # Cleanup
    Remove-Item $TempDir -Recurse -Force

    Write-Host "✅ CodeMaxx installed successfully!" -ForegroundColor Green
    Write-Host "🔁 Restart terminal and run: codemaxx --help"
}

# ---------------- SCRIPT EXECUTION ----------------
Ensure-Admin
Show-ASCII
Write-Host "🚀 Installing CodeMaxx CLI Tool..." -ForegroundColor Cyan

$Version = Get-LatestVersion
$ZipPath = Download-CodeMaxx -Version $Version
Install-CodeMaxx -ZipPath $ZipPath