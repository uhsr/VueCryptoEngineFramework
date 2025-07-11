# VueCryptoEngineFramework

## Description



## Features

- Integrates WebAssembly modules for accelerated cryptographic operations within the Vue.js framework.
- Provides a reactive Vuex module for managing and persisting cryptographic keys securely in the browser's local storage.
- Implements a custom Vue directive for easily encrypting and decrypting data bindings within templates using AES-256-GCM.
- Generates cryptographically secure random numbers using the Web Crypto API and exposes them as reactive data properties.
- Supports ECDSA signature verification with customizable curve parameters directly within Vue components.
- Offers a pluggable architecture for integrating different cryptographic libraries and algorithms based on project requirements.
- Includes a dedicated Vue component for displaying and verifying cryptographic hashes (SHA-256, SHA-3, etc.) of user-uploaded files.
- Validates the integrity of downloaded resources using Subresource Integrity (SRI) hashes to prevent malicious code injection.
## Installation

```bash
go get github.com/uhsr/VueCryptoEngineFramework
```

## Usage

```go
package main

import (
    "vuecryptoengineframework/internal/vuecryptoengineframework"
)

func main() {
    app := vuecryptoengineframework.NewApp(false)
    app.Run()
}
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
