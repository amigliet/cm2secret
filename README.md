# cm2secret

Simple tool to quickly convert a ConfigMap to a Secret.


## Installation
### Build
Firstly, clone the repository and build the binary from source:

```
$ git clone https://github.com/amiglietta/cm2secret.git
$ cd cm2secret
$ make build
```

### Install
This installation places the executable in your local **$HOME/.local/bin**
(no superuser's priveleges are required):

```
$ make install
```

### Uninstall
To uninstall **cm2secret** binary:

```
$ make uninstall
```

## Usage
To view the command line syntax use the option **-h** or **--help**:

```
$ cm2secret -h
Usage of cm2secret:
  -f string
    	ConfigMap filename to convert.
  -o string
    	Output format. Choose 'json' or 'yaml'. (default "json")
```
