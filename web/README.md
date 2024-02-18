# Web

Setup:

```bash
ln -s ../unicode-conversion-maps/maps/ font-maps
ln -s ../normalizer/libindic/normalizer/ normalizer-rules

cd ..
make web
```

## Development

```bash
python3 -m http.server 3000
```

Then go to http://localhost:3000/template.html for development.

If a new map file is added in `maps/`, then do a new build of `index.html` to make the new `<select>`.

```bash
go run build_index_html.go
```
