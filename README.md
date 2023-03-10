# Qas | API

Qas backend API built against Golang `net/http`.

## Endpoints

| Method | Pattern                  | Code | Action                           |
|--------|--------------------------|------|----------------------------------|
| GET    | /config/list             | 200  | Fetches all configurations.      |
| GET    | /config/one?lang         | 200  | Fetch a single configuration.    |
| POST   | /config/new              | 201  | Create a new configuration.      |
| PUT    | /config/replace?lang     | 200  | Overwrite a configuration.       |
| PATCH  | /config/update?lang?info | 200  | Append project to configuration. |
| DELETE | /config/destroy?lang     | 204  | Deletes a configuration.         |

## Port

Default port is at `:5000/VERSION`

## Configurations

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

## GNU Guix

To load all system dependencies, just run `guix shell`

## TODO

- implement simpler Alice-like chaining
- config: branch defaults to master

## LICENSE

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
