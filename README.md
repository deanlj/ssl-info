# poll

ssl-info is used to gather interesting information about ssl endpoints and certs.  It is licensed by the [MIT](LICENSE) license.

## build

install go-lang, dep, make; run `make`

## Binaries

- [Mac](bin/mac/poll)

- [Linux](bin/linux/poll)

- [Windows](bin/windows/poll.exe)

[SHA256](sha256sums.txt) for binaries.

### remote

Ever wish there was a command line tool that would print the most interesting information about
a SSL certificate on a remote web server?  

Wish this tool would just "pretty print" JSON on your terminal, and would work the same regardless of popular
operating system you were running?

Now you've got it:

```
bin/mac/ssl-info remote -a www.starbucks.com
{
  "DNS Names": [
    "www.starbucks.com",
    "globalassetshost.starbucks.com",
    "Starbucks.com"
  ],
  "Issuer Organization": [
    "Symantec Corporation"
  ],
  "NotAfter": "2018-09-12T23:59:59Z",
  "NotBefore": "2017-09-11T00:00:00Z",
  "Signature Algorithm": "SHA256-RSA",
  "Supported SSL Versions": [
    "TLS12",
    "TLS10",
    "TLS11"
  ]
}
```
