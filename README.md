# Terraform Provider for Uptime.com
## Requirements
* Terraform v0.12.0 or higher
* Go v1.12 or higher

## Installation

1. Configure git to use SSH instead of HTTPS
2. Allow Go to use private repo packages from Bitbucket
3. Build the project
4. Create location in terraform directory for plugin
5. Symlink executible

```bash
# 1 Use SSH instead of HTTPS
git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

# 2 Use private repo
export GOPRIVATE=bitbucket.org

# 3 Build
go build -o ~/go/bin/terraform-provider-uptime

# 4. Create destination for executable
mkdir -p ~/.terraform.d/plugins/integral.com.au/local/uptime/1.0.1/darwin_amd64/

# 5. symlink executable to destination
ln -s ~/go/bin/terraform-provider-uptime \
    ~/.terraform.d/plugins/integral.com.au/local/uptime/1.0.1/darwin_amd64/terraform-provider-uptime
```

For Linux machines, follow the OS X process, replacing `darwin` with `linux`.

For a Windows machine, in PowerShell:
```powershell
New-Item %APPDATA%\terraform.d\plugins\windows_amd64 -Type 'directory' -Force
cmd /c mklink /d $env:GOPATH\bin\terraform-provider-uptime %APPDATA%\terraform.d\plugins\windows_amd64\terraform-provider-uptime
```

## Credits
terraform-provider-uptime was originally created by [Kyle Gentle](https://github.com/kylegentle), with support from Elias Laham and the Dev Team at Uptime.com.
