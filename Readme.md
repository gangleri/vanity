# Vanity

A simple cli utility to generate HTTML pages that can be used to give packages a vanity 
import path. This is an update on [canonical-gen](https://github.com/gangleri/canonical-gen) that 
uses embed and has much less code

## Installation 

```shell
go get gangleri.com/pkg/vanity
```

## Usage 

```shell
vanity [-p] pkg repo 
```

- `-p` is an optional flag used to indicate if the underlying repo is private 
- `pkg` the vanity name that will be used by your package
- `repo` the repo that hosts the actual code



example:

For a public repo 
```shell
vanity example.com/pkg github.com/username/repo
```

For a private repo 
```shell
vanity -p example.com/pkg github.com/username/repo
```

