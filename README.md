# teamscli

requirements: working microsoft teams app (windows/linux)

It works by reading the cookies database and reusing your auth cookie to do stuff (currently only setting presences)

This must run on the same host as you're running your GUI teams.

## presence

```bash
./teamscli help presence
possible options are Busy, Available, Offline, DoNotDisturb, Away, BeRightBack

Usage:
  teamscli presence [flags]

Flags:
  -h, --help   help for presence
```

```bash
./teamscli presence Busy
presence set to Busy
```
