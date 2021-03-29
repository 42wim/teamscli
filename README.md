# teamscli

requirements: working microsoft teams app (windows/linux)

It works by reading the cookies database and reusing your auth cookie to do stuff (currently only setting presences)

This must run on the same host as you're running your GUI teams.

## presence

```bash
./teamscli help presence
possible options are Busy, Available, Offline, DoNotDisturb (DND), Away, BeRightBack (BRB)

Usage:
  teamscli presence [flags]

Flags:
  -d, --duration string   keep presence for specific time (e.g. 1h)
  -h, --help              help for presence
```

```bash
./teamscli presence Busy
presence set to Busy
```

```bash
./teamscli presence brb -d 1h30m
presence set to brb for 1h30m
```