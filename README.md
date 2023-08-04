# ip_checker.io

This tool is used to retrieve information about an IP address, this tool uses the ipinfo.io platform

Due to the limitations of the ipinfo.io token, you need to replace the api token in the source code first before using it

How to use it
```
λ  ~ cd ip_checker.io
λ  ip_checker.io git:(main) go build ip_checker.go
λ  ip_checker.io git:(main) ✗ ./ip_checker --help
Develop by: LazyySec
Usage of ./ip_checker:
  -host string
    	Check single IP address. Provide the IP address to check.
  -l	Check using list. Provide a filename containing a list of IP addresses.
λ  ip_checker.io git:(main) ✗ ./ip_checker -host 38.242.147.249
Develop by: LazyySec
IP Address: 38.242.147.249
Info:
ip: 38.242.147.249
hostname: vmi1146818.contaboserver.net
city: Düsseldorf
loc: 51.1809,6.8598
postal: 40599
company: map[domain:contabo.com name:Contabo GmbH type:business]
region: North Rhine-Westphalia
country: DE
timezone: Europe/Berlin
asn: map[asn:AS51167 domain:contabo.com name:Contabo GmbH route:38.242.144.0/21 type:hosting]
----------------------------------------
```
Happy using. 
