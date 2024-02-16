# payyans

This is the Go port of [payyans](https://github.com/smc/payyans) and [freaknz](https://gitlab.com/kannanvm/freaknz-qt) together as a shared library. The goal is the logic abstraction through a shared library.

## Usage

```bash
./payyanscli -map maps/maps/ML-TTKarthika.map file-to-convert.txt
./payyanscli -map maps/maps/ML-TTKarthika.map file-to-convert.txt output-file.txt
```

## Development

```bash
go run cli/*.go -map maps/maps/ML-TTKarthika.map payyans/testdata/ml-ttkarthika.txt
go test payyans/*.go
```

## Goals

- [x] ASCII to Unicode conversion
- [ ] Write more unit tests for ASCII to Unicode conversion
- [ ] Unicode to ASCII conversion
- [ ] Include all the Malayalam ASCII fonts
- [ ] Shared library
- [ ] Make a WASM web app
