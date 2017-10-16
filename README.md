# dmocks

[![Build Status](https://travis-ci.org/kogasoftware/dmocks.svg?branch=master)](https://travis-ci.org/kogasoftware/dmocks)

[![GoDoc](https://godoc.org/github.com/kogasoftware/dmocks?status.svg)](https://godoc.org/github.com/kogasoftware/dmocks)

Dynamic mock server

## Usage

```
# Create config file
vim config.yml

dmocks -p 3000 -c config.yml
```

## Options

```
Usage: dmocks [options]

options:
  -p, --port
      set port (Default value is 3000)
  -c, --config
      set config file (Default value is 'config.yml')
  -h, --help
      show help message
```

## config.yml

See [config.sample.yml](https://github.com/kogasoftware/dmocks/blob/master/config.sample.yml)
