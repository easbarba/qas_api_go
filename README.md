# Qas | API

## Config Endpoint

| HTTP method | Path                            | Status Code      | Description                        |
| ----------- | ------------------------------- | ---------------- | ---------------------------------- |
| GET         | /v1/configs                     | 200 (OK)         | Fetches all Configs resources.     |
| GET         | /v1/configs/view?lang           | 200 (OK)         | Fetch a single Config resource.    |
| POST        | /v1/configs/new                 | 200 (CREATED)    | Create a new Config resource.      |
| PUT         | /v1/configs/update?lang         | 200 (OK)         | Updates a Config resource.         |
| PATCH       | /v1/configs/append?lang?project | 200 (OK)         | Append project to Config resource. |
| DELETE      | /v1/configs/delete?lang         | 204 (No content) | Deletes a Config resource.         |

## Guix

To load all dependencies, just run `guix shell`

## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
