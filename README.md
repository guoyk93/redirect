# redirect

[![BMC Donate](https://img.shields.io/badge/BMC-Donate-orange)](https://www.buymeacoffee.com/vFa5wfRq6)

A simple HTTP service doest nothing but redirect

## Get

`docker pull guoyk/redirect`

## Env

* `PORT` port to listen, default to `80`
* `TARGET` target location, for example `https://example.com` or `https://example.com/subpath`

## Path Combination

If `TARGET` has `/` suffixed, original path will be appended to `TARGET`

## Health Check

```text
GET /healthz

OK
```

## Credits

Guo Y.K., MIT License