# Terraform Provider for Uptime.com
## Requirements
* Terraform v0.12.0 or higher
* Go v1.12 or higher

## Installation
### Downloading the provider
First, install the provider to your local machine:
```
go get -u -v github.com/uptime-com/terraform-provider-uptime
```

### Installing
In order for Terraform to use terraform-provider-uptime, it needs to be linked to the plugin directory. Example commands for an OS X Darwin machine:

```bash
mkdir -p ~/.terraform.d/plugins/integral.com.au/local/uptime/1.0.1/darwin_amd64/

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
