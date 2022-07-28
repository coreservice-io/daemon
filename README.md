# DAEMON

Install app as system service [support Linux and Darwin]

Precompiled files in /build

daemon-darwin-amd64<br />
daemon-linux-amd64<br />
daemon-linux-arm64<br />
daemon-linux-arm32<br />


## How to use
1. Compile program according to the os-arch system. 
2. Copy the compiled file into your project package and rename to "daemon".
3. Run cmd "sudo ./daemon [install/remove/start/stop/restart/status] [app name]"to manage process

### for example
```
├─{your-project-folder}
│  ├─configs    //cofig folder
│  ├─logs       //log folder
│  ├─assets     //assets folderr
│  ├─myapp      //executable file
│  └─daemon    //daemon file compiled and copy from this package
```

run cmd
```
//enter {your-project-folder}
cd ./{your-project-folder}

sudo ./daemon install myapp [arg1] [arg2] ...
sudo ./daemon start myapp
sudo ./daemon status myapp
sudo ./daemon restart myapp
sudo ./daemon stop myapp
sudo ./daemon remove myapp
```


### in the test folder there are two test apps for arm32 and arm64 architecture

#### you can run on arm64 op-system
```
sudo ./daemon install app_test-linux-arm32
```
#### or run on arm32(armv6,armv7,etc..) op-system
```
sudo ./daemon install app_test-linux-arm64
```





