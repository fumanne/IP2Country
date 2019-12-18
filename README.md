# Introduce

IP2Country is a tool convert ipaddress to country code. 
The idea is from other private repository and thanks the author
Now I rewrite it with golang to public. 
And I will add more functions and fix code in further when I am free


## Usage

```
$ ipcountry -h
Convert ip address to country code.
Also it can update ipdata file from internet.
For example:
  IP2Country update // update data file
  IP2Country search 8.8.8.8  // convert address to country code.

Usage:
  IP2Country [command]

Available Commands:
  help        Help about any command
  search      covert ip address to country code
  update      update data file from internet
  version     show version of IP2Country

Flags:
  -h, --help   help for IP2Country
 
```   
## Todo
    1. support IPV6    x
    2. optimization code (sqlite3 to store ip data) √
       2.1 update is too slow, need to optimize   x
    3. add build scripts and deployment (dockerfile) √
    4. add ipaddress Judgement (regexp) √
    


    

 