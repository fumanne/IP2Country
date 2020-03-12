# Introduce

IP2Country is a tool convert ipaddress to country code. 
The idea is from other private repository and thanks the author
Now I rewrite it with golang to public. 
And I will add more functions and fix code in further when I am free


## Install

```
$ go get github.com/fumanne/IP2Country
$ cd /GOPATH/bin
$ IP2Country update
$ IP2Country search 178.134.11.231
$ GE 

```


 
## Build From Source

```
$ sh build/buildall.sh
Building darwin
Building linux
Building windows

# example:
$ IP2Country.linux update 
$ IP2Country.linux search 2c0f:ff10::12
```



## Usage

```
$ ip2country -h
Convert ip address to country code.
Also it can update ipdata file from internet.
For example:
  IP2Country update // update data file
  IP2Country search 8.8.8.8  // convert address to country code.
  IP2Country search 2c0f:ff10::12 // convert ipv6 address 

Usage:
  IP2Country [command]

Available Commands:
  help        Help about any command
  search      covert ip address to country code
  update      update data file from internet
  version     show version of IP2Country

Flags:
  -h, --help   help for IP2Country
  
  
$ ip2country update --help
  update data file from internet and cached in $HOME/.IP2Country/
  
  Usage:
    IP2Country update [flags]
  
  Flags:
    -h, --help    help for update  
  
  
 
```   
## Feature
    1. supported IPV6     
    2. build scripts    
    3. Private IPv4 filter
    

## TODO
    1. Add context (Timeout) when no country code found (Done)
    2. Optimize code to good performance 

    


    

 
