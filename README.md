# Service

Install app as system service, support Linux and Darwin

Precompiled files in /build

service-darwin-amd64<br />
service-linux-amd64<br />
service-linux-arm64<br />

If you need other system files, please compile it yourself.

## How to use
1. Compile program according to the operating system. 
2. Copy the compiled file into your project package and rename to "service".
3. Run cmd "sudo ./service [install/remove/start/stop/restart/status] [app name]"to manage process

### for example
```
├─yourProjectFolder
│  ├─configs    //cofig folder
│  ├─logs       //log folder
│  ├─assets     //assets folderr
│  ├─myapp      //executable file
│  └─service    //service file compiled and copy from this package
```

run cmd
```
//enter yourProjectFolder
cd ./yourProjectFolder

sudo ./service install myapp [arg1] [arg2] ...
sudo ./service start myapp
sudo ./service status myapp
sudo ./service restart myapp
sudo ./service stop myapp
sudo ./service remove myapp
```


