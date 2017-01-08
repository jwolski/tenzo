# Instructions for the Linux Cook

## cgroups and namespaces

## Commands

### dig
DNS lookup utility

### dstat
versatile tool for generating system resource statistics

```
$ dstat
You did not select any stats, using -cdngy by default.
----total-cpu-usage---- -dsk/total- -net/total- ---paging-- ---system--
usr sys idl wai hiq siq| read  writ| recv  send|  in   out | int   csw
0   0 100   0   0   0| 945B 5937B|   0     0 |   0     0 |   4     4
1   0  99   0   0   0|   0     0 | 416B  444B|   0     0 |  17    22
0   0 100   0   0   0|   0     0 | 104B  236B|   0     0 |  11    14
```

_Does not come installed by default on Ubuntu. Therefore, `apt-get install
dstat` is required._

### files
Determine file type

```
$ file main
main: ELF 64-bit LSB  executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

### htop

### iotop

### ldd
Print shared library dependencies

### lsof

### mtr
My Traceroute. Combines traceroute and ping.

### netstat
Print network connections, routing tables, interface statistics, masquerade
connections, and multicast memberships

### ngrep
network grep

### nm
Lists symbols from object files

### nslookup
query Internet name servers interactively

### /proc/net/tcpstat

### strace
trace system calls and signals

### tc

### tcpdump
dump traffic on a network

### telnet

### traceroute
print the route packets trace to network host

### tree
list contents of directories in a tree-like format.

```
$ tree
.
├── main.go
├── pkgone
│   └── source.go
├── pkgthree
│   └── source.go
└── pkgtwo
└── source.go
3 directories, 4 files
```

### whois
client for the whois directory service

## Copy-on-write

## dmesg
dmesg is a command that prints the message buffer of the kernel.

## Kernel Ring Buffer
The kernel ring buffer is a data structure that records messages related to the
operation of the kernel. A ring buffer is a kind of buffer that is always a
constant size, removing the oldest messages when new messages come in.

## Memory Mapping

## Page Cache
* Page cache is also known as the disk cache
* The operating system keeps a page cache in unused portions of main memory

### Links
* https://en.wikipedia.org/wiki/Paging

## Signals

## Questions

### How to find version of Ubuntu?
```
$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 14.04.5 LTS
Release:        14.04
Codename:       trusty
```
