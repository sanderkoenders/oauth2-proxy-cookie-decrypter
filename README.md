# Go OAuth2 Proxy Cookie Decryptor

This Go project decrypts OAuth2 proxy cookies using a shared secret defined in the `.env` file. The program is designed to read and decrypt the combined cookie values of `_oauth2_proxy_0` and `_oauth2_proxy_1`.

## Prerequisites

- **Go 1.16+** (or a later version)
- **.env file** for secret management

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/sanderkoenders/oauth2-proxy-cookie-decrypter.git
   cd oauth2-proxy-cookie-decrypter
   ```

2. **Install dependencies:**

   If you're using Go Modules, dependencies should automatically be fetched. Otherwise, run the following command:

   ```bash
   go mod tidy
   ```

3. **Create the `.env` file:**

   In the root of your project, create a `.env` file and define the `COOKIE_SECRET` variable. It should contain the secret used for cookie encryption.

   Example `.env` file:

   ```
   COOKIE_SECRET=my-secret-key
   ```

4. **Build the project:**

   ```bash
   go build -o decrypt-cookie
   ```

5. **Run the program:**

   The program expects the combined cookie values of `_oauth2_proxy_0` and `_oauth2_proxy_1` as a single argument (enclosed in quotes). For example:

   ```bash
   ./decrypt-cookie "<_oauth2_proxy_0><_oauth2_proxy_1>"
   ```

   - Replace `_oauth2_proxy_0` and `_oauth2_proxy_1` with the actual combined cookie values.

## Example

Assuming the values for `_oauth2_proxy_0` and `_oauth2_proxy_1` are `_oauth2_proxy_0` and `_oauth2_proxy_1`, run the following command:

```bash
./decrypt-cookie "<_oauth2_proxy_0><_oauth2_proxy_1>"
```

The program will output the decrypted IDToken and AccessToken.

## Notes

- Ensure the `.env` file is **not committed to version control**. You can add it to `.gitignore` to prevent accidental commits.
