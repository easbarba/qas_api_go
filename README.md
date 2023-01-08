# Qas | API

## Config Endpoint

| Method | Path                     | Status Code      | Description                        |
| ------ | ------------------------ | ---------------- | ---------------------------------- |
| GET    | /v1/cfg/all              | 200 (OK)         | Fetches all Configs resources.     |
| GET    | /v1/cfg/one?lang         | 200 (OK)         | Fetch a single Config resource.    |
| POST   | /v1/cfg/new              | 200 (CREATED)    | Create a new Config resource.      |
| PUT    | /v1/cfg/update?lang      | 200 (OK)         | Updates a Config resource.         |
| PATCH  | /v1/cfg/append?lang?info | 200 (OK)         | Append project to Config resource. |
| DELETE | /v1/cfg/delete?lang      | 204 (No content) | Deletes a Config resource.         |

## Guix

To load all dependencies, just run `guix shell`

## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
