primes  [![GoDoc](https://godoc.org/github.com/otiai10/primes?status.svg)](https://godoc.org/github.com/otiai10/primes) [![Circle CI](https://circleci.com/gh/otiai10/primes.svg?style=svg)](https://circleci.com/gh/otiai10/primes)
==========

- Find primary numbers
```sh
% primes p 20
[2 3 5 7 11 13 17 19]
```
- Factorize numbers
```sh
% primes f 329
[7 47]
```
- Reduce fractions
```sh
% primes r 144/360
2/5
```

- all can be executed in go code
```
package main

import (
	"fmt"
	"github.com/otiai10/primes" // :$ go get github.com/otiai10/primes/primes
)

func main() {
	fmt.Println(
	primes.Factorize(144).All(),
	primes.Factorize(144).Powers(),)
	}

// [2 2 2 2 3 3]
// map[2:4 3:2]
```

# install

```sh
go get github.com/otiai10/primes/primes
```
