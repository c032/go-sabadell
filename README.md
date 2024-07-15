# go-sabadell

## Executables

### `sabadell-to-json`

Usage:

```
sabadell-to-json exported.txt
```

That command prints one JSON object per line, for example:

```
{"OperatingDate":"15/07/2024","Concept":"Example.","ValueDate":"Example.","Amount":"0.00","Balance":"0.00","Ref1":"","Ref2":""}
```

Values are **not** parsed, this only converts a `.txt` file to a
simplified JSON that can be used for further processing.

## License

MPL 2.0.
