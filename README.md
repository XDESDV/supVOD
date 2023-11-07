# supVOD
Online video streaming at SupDeVinci

## Installation

```console
$ cp .env.example .env
```

## Routes

### URI

Base URI: http://localhost:${PORT}

### Users

| Method | URI | Name |
| --- | --- | --- |
| POST | /v1/users | Create user| 
| PUT | /v1/users | Update User|
| GET | /v1/users/:user_id | Get User By Id|

### Kinds

| Method | URI | Name |
| --- | --- | --- |
| POST | /v1/kinds | Create kind | 
| PUT | /v1/kinds | Update Kind|
| GET | /v1/kinds | List Kind|

### Movies

| Method | URI | Name | Query Optional Params |
| --- | --- | --- | --- |
| POST | /v1/movies | Create movie | |
| PUT | /v1/movies | Update movie| |
| GET | /v1/movies | List movies| page,limit,search (search string),kinds (array of kinds) |
| GET | /v1/movies/:movie_id | Get movie| |

### Historics

| Method | URI | Name | Query Optional Params |
| --- | --- | --- | --- |
| POST | /v1/historics | Create historics | |
| GET | /v1/historics | List historics | user_id |
| GET | /v1/movies/:historics_id | Get specific histric | |


