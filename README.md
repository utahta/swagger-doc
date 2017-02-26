# swagger-doc

swagger-doc is tool of easy visualize your swagger.json using swagger-ui and redoc.

## Install

```bash
go get -u github.com/utahta/swagger-doc/cmd/swagger-doc
```

## Usage


### Local file

```bash
swagger-doc -s /path/to/swagger.json
```
then open browser.


### URL

```bash
swagger-doc -s http://localhost/swagger.json
```
then open browser.


## Contributing

1. Fork it ( https://github.com/utahta/swagger-doc/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
