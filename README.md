# md5hashtool

This is a md5hashtool which makes http requests to the external server and prints the address of the request along with the MD5 hash of the response.

## Prerequisites

Need to have Golang installed on your system. If not, you can install it using below doc:

```bash
https://golang.org/doc/install
```

## Getting Started

Clone the git repository in your system and then cd into project root directory

```bash
$ git clone git@github.com:Anand55/md5hashtool.git
$ cd md5hashtool
```

Build the tool by executing the below commands

```bash
$ go build -o myhttp
```

## Example Outputs

This tool takes address of the http requests as command line arguments and returns hash value of its body. See below examples

```bash
$ ./myhttp google.com
http://google.com 17787a1f24e3437b902ebd25c0bee804
```

```bash
$ ./myhttp google.com http://gmail.com
http://github.com 875f2dea1bc62cc43058c0504a70131b
http://twitter.com 891609f5fc22d8eb6a496c313856cd4a
```

This tool takes "parallel" option (integer value) to execute the http requests parallely. The default value of "parallel" is 10.

```bash
$ ./myhttp -parallel 3 google.com adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
http://google.com 3825e1eb64f023cc000952c3aaaa6871
http://adjust.com 8317d92e2ed17d6b11ee1a8d3951a0d0
http://facebook.com 079d91ddb8a03e2ecff6230bf637eb7c
http://yandex.com 3c3d6d74d7b0b4e63a4603eed4c40d37
http://twitter.com 7680b57a6857260859d99aa9be974cc6
http://yahoo.com df543f11ee67557ae19a9b7839084b97
```

It will throw errors when address is incorrect

```bash
$ ./myhttp googt
Error making http request:  Get "http://googt": dial tcp: lookup googt: Temporary failure in name resolution
```
