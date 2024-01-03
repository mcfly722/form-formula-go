## Install
To compile "C" code
1. install [tdm-gcc](https://jmeubank.github.io/tdm-gcc/)
2. if you using [VSCode](https://code.visualstudio.com/), according [issue #38990](https://github.com/golang/go/issues/38990), set
```
{
  "gopls": {
    "build.env": {
      "GOARCH": "386"
    }
  }
}
```
