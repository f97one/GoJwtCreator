# GoJwtCreator
cli app to create JWT 

## Requirements

- Go 1.13 or later (latest stable release recommended)
- dgrijalva/jwt-go

## usage

```bash
JwtCreator {-expires [expires]} -issuer [issuer] {-nolf} -private [private]
  -expires int
        Time to expiration (in minutes, 30 minutes if not specified) (default 30)
  -help
        Display usage and exit
  -issuer string
        issuer string (required)
  -nolf
        not append LF to end of line
  -private string
        Private Key file to sign JWT (required)
```
