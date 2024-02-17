# payyans

This is the Go port of [payyans](https://github.com/smc/payyans) and [freaknz](https://gitlab.com/kannanvm/freaknz-qt) together as a shared library. The goal is the logic abstraction through a shared library.

## Usage

By default the `payyanscli` binary comes with some font map files. See list of fonts:

```bash
payyanscli -fonts
```

Convert:

```bash
# Output to terminal
./payyanscli -font ML-TTKarthika.map file-to-convert.txt

# Output to a file
./payyanscli -font ML-TTKarthika.map file-to-convert.txt output-file.txt
```

See `payyanscli -h` for more.

## Development

```bash
go run cli.go -font ML-TTKarthika.map file-to-convert.txt
make test
```

## Goals

- [x] ASCII to Unicode conversion
- [ ] Write more unit tests for ASCII to Unicode conversion
- [ ] Unicode to ASCII conversion
- [ ] Include all the Malayalam ASCII fonts
- [ ] Shared library
- [ ] Make a WASM web app
