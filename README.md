# tkai_circles
A coursework to complete the architecture and infrastructure final project. This work would be deployed in one of deployment sites such as Microsoft Azure, Google Cloud Platform, and AWS

# How to Use
There are 5 API in total,

## Register User
```http request
POST /api/user/new
```
This API is used for new user registration.
### Request Body

```javascript
{
  "username": "refoo",
  "password" : "aaaaaaaaaa"
}
```
### Response Body
```javascript
{
    "account": {
        "ID": 3,
        "CreatedAt": "2019-12-04T17:06:41.641046+07:00",
        "UpdatedAt": "2019-12-04T17:06:41.641046+07:00",
        "DeletedAt": null,
        "username": "refoo",
        "password": "",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjN9.g_tVFfrVjTDWM_XcBtcwVX9-tDtJumHVqKNIJ9lIF4k"
    },
    "message": "Account has been created",
    "status": true
}
```

## Login User
```http request
POST /api/user/login
```
This api is for user to login and get the token for authorization..
### Request Body
```javascript
{
  "username": "refoo",
  "password" : "aaaaaaaaaa"
}
```
### Response Body
```javascript
{
    "account": {
        "ID": 3,
        "CreatedAt": "2019-12-04T17:06:41.641046+07:00",
        "UpdatedAt": "2019-12-04T17:06:41.641046+07:00",
        "DeletedAt": null,
        "username": "refoo",
        "password": "",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjN9.g_tVFfrVjTDWM_XcBtcwVX9-tDtJumHVqKNIJ9lIF4k"
    },
    "message": "Logged In",
    "status": true
}
```

## Circle Area
```http request
POST /api/circle/area
```
This api is to create new circle object, the result will be saved into the database and can be used to calculate 
tube and ball objects.
### Headers
| Key | Value | Description |
| :--- | :--- | :--- |
| `Authorization` | `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.DzOJ7GHkPwiDE3T78dFMriY96VwzytQSBV7-c64dxx8` | **Required**. Your Token from login or registration |
### Request Body
```javascript
{
    "radius" : 10
}
```
### Response Body
```javascript
{
    "circle": {
        "ID": 6,
        "CreatedAt": "2019-12-04T16:58:21.298678+07:00",
        "UpdatedAt": "2019-12-04T16:58:21.298678+07:00",
        "DeletedAt": null,
        "radius": 10,
        "area": 314.15927,
        "owner": "refo"
    },
    "message": "success",
    "status": true
}
```
## Tube Volume
```http request
POST /api/tube/volume
```
This api is to create new tube object, the result will be saved into the database. This API require the circle object 
id created from the above circle API.
### Headers
| Key | Value | Description |
| :--- | :--- | :--- |
| `Authorization` | `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.DzOJ7GHkPwiDE3T78dFMriY96VwzytQSBV7-c64dxx8` | **Required**. Your Token from login or registration |
### Request Body
```javascript
{   
    "circleId" : 6,
    "height" : 10
}
```
### Response Body
```javascript
{
    "message": "success",
    "status": true,
    "tube": {
        "ID": 1,
        "CreatedAt": "2019-12-04T16:59:11.346943+07:00",
        "UpdatedAt": "2019-12-04T16:59:11.346943+07:00",
        "DeletedAt": null,
        "circleId": 6,
        "height": 10,
        "volume": 3141.5928,
        "owner": "refo"
    }
}
```
## Ball Volume
```http request
POST /api/ball/volume
```
This api is to create new ball object, the result will be saved into the database. This API require the circle object 
id created from the above circle API.
### Headers
| Key | Value | Description |
| :--- | :--- | :--- |
| `Authorization` | `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.DzOJ7GHkPwiDE3T78dFMriY96VwzytQSBV7-c64dxx8` | **Required**. Your Token from login or registration |
### Request Body
```javascript
{
	"circleId" : 6
}
```
### Response Body
```javascript
{
    {
        "ball": {
            "ID": 2,
            "CreatedAt": "2019-12-04T17:01:01.690954+07:00",
            "UpdatedAt": "2019-12-04T17:01:01.690954+07:00",
            "DeletedAt": null,
            "circleId": 6,
            "volume": 4188.79,
            "owner": "refo"
        },
        "message": "success",
        "status": true
    }
}
```

