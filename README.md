# Urna: In-Memory Cache Library for Golang

Urna is a lightweight and efficient in-memory cache library for Golang. It is designed to provide a simple and flexible caching solution for your Golang applications, helping you improve performance by storing frequently used data in memory.

Please note that Urna is still **under active development** and it's not **production-ready**.

## Features

- **Easy to Use:** Urna comes with a straightforward API that makes it easy to integrate caching into your Golang applications.
- **High Performance:** With an efficient in-memory caching mechanism, Urna ensures fast access to cached data, reducing the need for time-consuming data retrieval operations.
- **Configurability:** Customize the cache behavior to suit your application's specific requirements, including cache expiration, eviction policies, and more.
- **Concurrency-Safe:** Urna is designed to be safe for concurrent use, making it suitable for multi-threaded applications.

## Installation

To use Urna in your Golang project, simply run:

```bash
go get github.com/1garo/urna
```

## Usage 
```go
package main

import (
	"fmt"
	"time"

	"github.com/1garo/urna"
)

func main() {
	// Create a new cache with a default expiration time of 5 minutes
	cache := urna.NewCache(5 * time.Minute)

	// Add data to the cache
	cache.Set("key", "value")

	// Retrieve data from the cache
	result, found := cache.Get("key")
	if found {
		fmt.Println("Value:", result)
	} else {
		fmt.Println("Key not found in the cache.")
	}
}

```

## License
Urna is licensed under the [MIT License.](LICENSE)
