# Go OAuth2 Proxy Cookie Decryptor

This Go project decrypts OAuth2 proxy cookies using a shared secret defined in the `.env` file. The program is designed to read and decrypt the combined cookie values of `_oauth2_proxy_0` and `_oauth2_proxy_1`.

## Prerequisites

- **Go 1.16+** (or a later version)
- **.env file** for secret management

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://your-repo-url.git
   cd your-repo-directory
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

   - **Important**: The `COOKIE_SECRET` must be a valid AES key of 16, 24, or 32 bytes.

4. **Build the project:**

   ```bash
   go build -o decrypt-cookie
   ```

5. **Run the program:**

   The program expects the combined cookie values of `_oauth2_proxy_0` and `_oauth2_proxy_1` as a single argument (enclosed in quotes). For example:

   ```bash
   ./decrypt-cookie "cookie_value_0cookie_value_1"
   ```

   - Replace `cookie_value_0` and `cookie_value_1` with the actual combined cookie values.

## How It Works

- The program decrypts the combined cookie values using the AES algorithm with the secret key provided in the `.env` file.
- It expects the **combined values** of the `_oauth2_proxy_0` and `_oauth2_proxy_1` cookies to be passed as a single string (enclosed in quotes).
- The decrypted data is processed and printed to the terminal.

## Example

Assuming the values for `_oauth2_proxy_0` and `_oauth2_proxy_1` are `cookie_value_0` and `cookie_value_1`, run the following command:

```bash
./decrypt-cookie "cookie_value_0cookie_value_1"
```

The program will output the decrypted values and any relevant session data.

## Notes

- Ensure the `.env` file is **not committed to version control**. You can add it to `.gitignore` to prevent accidental commits.
- The secret key (`COOKIE_SECRET`) must be of valid AES length: **16, 24, or 32 bytes**.
- You can adjust the decryption logic as necessary based on your specific cookie structure and the encryption method used.
