# pastescraper

![Example Dashboard](screenshots/elk1.png)

A utility to scrape pastebin's incoming feed for known malware techniques.

You must have your ip whitelisted with pastebin for this to work.
[pastebin scraping docs](https://pastebin.com/doc_scraping_api)

## Usage

```go
// NOTE: This code does not check errors

c := &config{}
// fill in configuration info
err := c.getConf(configPath)

// load the requested parsers
parser, err := parse.New(c.Parsers)

// create the scraper with the loaded parsers
scraper, err := New(c, parser)

// start scraping
err = scraper.start(context.Background(), time.Second * 10)
```

## Overview
This project aims to be a modular solution to monitoring pastebin. It can be used to monitor any sort of content within pastebin, but all the currently implemented parsers are aimed at matching malware based techniques. This can be thought of as the following pipeline
```
1. Parsers
2. Normalizers
3. Post Actions
```

### Parsers
Parsers are the mechanisms used to determine the if the current paste is the desired content. This can be anything from a simple string search, to as advanced as complex analysis on the content itself.

### Normalizers
Normalizers are used as way to modify specifically targeted data. For instance if you are looking for the MZ header encoded in base64 and you get a sample that matches, maybe you want to normalize that by base64 decoding the entire content and storing that to disk.

### Post Actions
Post actions are actions that can be returned by a normalizer. For instance, if a normalizer returns a valid PE and ELF file then it can be uploaded to various services for further analysis. The following are services that will be supported in time:
- [X] Virustotal
- [ ] Malshare
- [ ] tria.ge
- [ ] polyswarm
- [ ] unpac.me

## Getting Started

### Standalone Docker Container
The easiest way to test the service itself is to create the following config file named `config.yml`. The parsers now run in order, so if there is something you are specifically trying to target put those parsers at the beginning. Items at the end of the array should be your catch all such as large hex blobs

```yaml
outputdir: "./pastes"
debug: False
maxqueuesize: 100
elastic:
  https: False
  username: ""
  password: ""
  host: ""
  port: 0
  index: ""
parsers: ["powershellFromBase64", "powershellWebClient", "vbsInvocation", "powershellScript", "powershellKeyword", "pythonSyscall", "bashHeader", "base64MZHeader", "base64ELFHeader", "rawMZHeader", "rawMachOHeader", "reverseBase64MZHeader", "reverseBase64ELFHeader", "largeHexBlob", "base64HighEntropy"]
```

Then creating the docker container with
```bash
docker build -t pastescrape:latest .
```

And starting the service with the newly created config file.

```bash
docker run -v ${PWD}/config.yml:/app/config.yml pastescrape:latest
```

If you want to preserve the downloaded pastes

```bash
docker run -v ${PWD}/pastes:/app/pastes -v ${PWD}/config.yml:/app/config.yml pastescrape:latest
```

The config file does support logging to an elasticsearch instance, so if that is configured those values can be filled in.


