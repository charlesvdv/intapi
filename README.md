# servapigen

> :construction::wrench: Please note this project is still in early stage.
> Feel free to give your input in the discussions tab or open an issue! 

`servapigen` is a `go` library to declare web service interfaces as code. The project 
has the same goal as [openapi](https://www.openapis.org/) but instead of using `json` or 
`yaml`, the full power of `go` is used. Read more about the *why* in the rationale section.

The project also define generators to export your declarations into client/server code, 
documentation or even in a openapi format.

## rationale

First, an obligatory *xkcd*:

![xkcd standards](https://imgs.xkcd.com/comics/standards.png)

API-first development is great to ensure consistency and understandability in your services. 
Most of the API-first development tooling are using data-serialization format like `json`, `yaml`, ...
to declare the API. This project takes a different take: the API is declared as code.

This provides various benefits:

- Easy composability without restrictions or custom extensions (like [`$ref` in JSON schema](https://json-schema.org/draft/2020-12/json-schema-core.html#referenced))
- Static types and documentation
- Full power of auto-completion and IDE-related features

On top of it, since the API definition is done as code, the generation is super customizable.

## examples

See the `examples` folder.

## generator

TODO