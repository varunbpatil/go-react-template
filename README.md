# go-react-template

A template for end-to-end type-safe Go + React SPA application using [buf connect](https://connectrpc.com) and Hexagonal architecture (Ports & Adapters pattern).

## Blog

Read the multi-part blog post covering this template [on my website](https://varunbpatil.github.io/tags/go-react-template/).

## Project Structure

This project follows hexagonal architecture (Ports & Adapters pattern).

For a real production app using this template, see [temporal-lens](https://github.com/varunbpatil/temporal-lens).

```
.
├── Makefile                   # Task runner
├── buf.yaml                   # Buf module config
├── buf.gen.yaml               # Buf code generation config
│
├── config/                    # App configuration
├── mocks/                     # Generated GoMock implementations
├── types/                     # Custom type definitions
│
├── protos/                    # Protobuf definitions
│   ├── src/                   # Proto source files
│   └── gen/                   # Generated code
│
├── domains/                   # App domains
│   ├── users/
│   │   ├── models/            # Domain models
│   │   ├── ports/             # Domain ports
│   │   ├── service/           # Domain service
│   │   └── errors.go          # Domain errors
│   │ 
│   └── ...
│
├── inbound/                   # Inbound adapters
│   ├── grpc/
│   ├── http/
│   ├── mcp/
│   └── ...
│
├── outbound/                  # Outbound adapters
│   ├── postgres/
│   └── ...
│
└── ui/                        # React frontend
```

## Development

This project uses:

* [Mise](https://mise.jdx.dev) to manage the development environment.
* [Makefile](./Makefile) as the task runner.
* [Tilt](https://tilt.dev/) as the process runner.
