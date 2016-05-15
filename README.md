# GoFetch

HTML download and parsing library with command line tool to print (in JSON) product detail's
 from the Sainsbury's website.

## Requirements

Golang 1.6.x

## Dependencies

Dependencies are installed into the vendor directory using godeps. If you are using go 1.6 you shouldn't need
to install any dependencies.

## Usage

```
go install
gofetch --url <url_here>
```

## Testing

```
go test ./...
```

# TODO

- The generic doParse function streams the results on a channel but the Product and ProductListParser
doesn't make use of this to return results as they are streamed
- Only the first page of product results will be fetched
- The ProductParser and ProductListParser should conform to an interface
