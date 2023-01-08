# Qas | API

Qas backend API built against Golang `net/http`.

## Config Endpoint

| Method | Path                     | Status Code      | Description                        |
| ------ | ------------------------ | ---------------- | ---------------------------------- |
| GET    | /v1/cfg/all              | 200 (OK)         | Fetches all Configs resources.     |
| GET    | /v1/cfg/one?lang         | 200 (OK)         | Fetch a single Config resource.    |
| POST   | /v1/cfg/new              | 200 (CREATED)    | Create a new Config resource.      |
| PUT    | /v1/cfg/update?lang      | 200 (OK)         | Updates a Config resource.         |
| PATCH  | /v1/cfg/append?lang?info | 200 (OK)         | Append project to Config resource. |
| DELETE | /v1/cfg/delete?lang      | 204 (No content) | Deletes a Config resource.         |

## Port

Default port is at `:5000`

## Configuration

`qas` looks for configuration files at `$XDG_CONFIG/qas`:

$XDG_CONFIG/qas/misc.json

```json
[
  {
    "name": "awesomewm",
    "branch": "master",
    "url": "https://github.com/awesomeWM/awesome"
  },
  {
    "name": "nuxt",
    "branch": "main",
    "url": "https://github.com/nuxt/framework"
  },
  {
    "name": "swift_format",
    "branch": "main",
    "url": "https://github.com/apple/swift-format"
  }
]
```

## Guix

To load all dependencies, just run `guix shell`

## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
