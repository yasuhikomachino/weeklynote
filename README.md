weeklynote - Generate a template for a weekly task list

## INSTALLATION

```
$ go install git@github.com:yasuhikomachino/weeklynote.git@latest
```

or download binary from [here](https://github.com/yasuhikomachino/weeklynote/releases).

## USAGE

```
$ weeklynote
---
2021-11-15(MON) - 2021-11-21(SUN)
---

# 2021-11-15(MON)

- [ ] task
- [ ] task
- [ ] task

# 2021-11-16(TUE)

- [ ] task
- [ ] task
- [ ] task

# 2021-11-17(WED)

- [ ] task
- [ ] task
- [ ] task

# 2021-11-18(THU)

- [ ] task
- [ ] task
- [ ] task

# 2021-11-19(FRI)

- [ ] task
- [ ] task
- [ ] task

# 2021-11-20(SAT)

- [ ] task
- [ ] task
- [ ] task

# 2021-11-21(SUN)

- [ ] task
- [ ] task
- [ ] task
```

## OPTIONS

```
--start value     start date(YY-MM-DD). Default is the first day of the week of the current day. (default: 2021-11-15)
--language value  display language. "en" or "ja". (default: en)
--location value  output location. "stdout" or "clipboard" (default: stdout)
--help, -h        show help (default: false)
--version, -v     print only the version (default: false)
```
