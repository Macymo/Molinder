# Molinder

## Services

### Recommendation Service

#### Requirements

1. Gender Preferences
2. Location
3. Age Preference

#### Routes: /v1/routes

##### /recommendations

* GET:
  - query:
    - gender:
      - Male
      - Female
    - location
      - State
    - age_preference
      - Ranges

##### 

### User Service

#### Requirements

1. Authentication
  - Create new users
  - auth users
2. User profile
  - Profile
    - Username
    - First name
    - Last name
    - Age
    - Male/Female
    - Profile pic

#### Routes /v1/user

##### /profile

* GET /<uid>: Get a certain user
    ```json
    {
        "uid": "string",
        "first_name": "string",
        "last_name": "string",
        "username": "string",
        "age": int,
        "profile_pic": "s3 link",
        "gender": emum
    }
    ```

* PUT: Create or update user
    request data:
    ```json
    {
        "uid": UID,
        "first_name": "string",
        "last_name": "string",
        "username": "string",
        "password": "string",
        "age": int,
        "profile_pic": "s3 link",
        "gender": emum
    }
    ```

##### /auth

* POST: authenticate to the server
    - request:
    ```json
    {
        "username": "string",
        "password": "string"
    }
    ```
    