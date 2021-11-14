# weeklynote

Generate a template for a weekly task list.

## Install

```
$ git get git@github.com:yasuhikomachino/weeklynote.git
```

or download binary from [here](https://github.com/yasuhikomachino/weeklynote/releases).

## Usage

```
$ weeklynote
```

output:

```
# 2021-11-15(MON) - 2021-11-21(SUN) 

## 2021-11-15(MON)
- [ ] task

## 2021-11-16(TUE)
- [ ] task

## 2021-11-17(WED)
- [ ] task

## 2021-11-18(THU)
- [ ] task

## 2021-11-19(FRI)
- [ ] task

## 2021-11-20(SAT)
- [ ] task

## 2021-11-21(SUN)
- [ ] task
```

## Options

```
--start value, -s value         Specify the start date(YY-MM-DD). Default is the first day of the week of the current day. (default: 2021-11-15)
--language value, --lang value  Specify the display language. "en" or "ja". (default: en)
--location value, --loc value   Specify the output location. "stdout" or "clipboard" (default: stdout)
--help, -h                      show help (default: false)
--version, -V                   print only the version (default: false)
```
