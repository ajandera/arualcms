![Screenshot 2021-03-25 at 13 56 24](https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png)


# Simple CMS made with love in VueJS and PHP


## Installation
### Use Docker
Copy content of production folder to your destination and run **docker-compose up**, it will create container and app will be running on next urls:

- forntend http://localhost:8001
- backend http://localhost:8001/admin
- api http://localhost:8001/api

Application has predefined credentials to be ready to use. Go to do admin url and insert

- username: arualcms@arualcms.com
- password: Password123

!! After first login change the pasword !!

### Use without docker
In that case you have run apache on your server or hosting and have PHP 7.3 installed. Than copy the content of production folder to your destination and go to urls as usual.

### Development
For development clone repository to your local machine and simple run docker-compose up from root folder. It creates all needed containers and run vuejs app in development mode. Application will be found on next ports:

- api http://localhost:8000
- admin http://localhost:8080
- frontend http://localhost:9000

### API description
## POSTS
endpoint: http://localhost:8000/post

method: GET

parameters: -

endpoint: http://localhost:8000/post/ID

method: GET

parameters: int ID

endpoint: http://localhost:8000/post

method: PUT

parameters: {
    id: 1
    title: "Test",
    body: "Test content",
    published: "2021-01-01",
    meta: {
        title: "Test",
        keywords: "test, arualCMS",
        description: "description"
    }
}

endpoint: http://localhost:8000/post

method: POST

parameters: {
    title: "Test",
    body: "Test content",
    published: "2021-01-01",
    meta: {
        title: "Test",
        keywords: "test, arualCMS",
        description: "description"
    }
}

endpoint: http://localhost:8000/post/ID

method: DELETE

parameters: int ID

## SETTINGS
endpoint: http://localhost:8000/setting

method: GET

parameters: -

endpoint: http://localhost:8000/setting/KEY

method: GET

parameters: string KEY

endpoint: http://localhost:8000/setting

method: PUT

parameters: [{
        "key": "author",
        "value": "Ales Jandera, shopyCRM"
    },
    {
        "key": "title",
        "value": "arualCMS"
    }
]

## TEXTS
endpoint: http://localhost:8000/text

method: GET

parameters: -


endpoint: http://localhost:8000/text/KEY

method: GET

parameters: string KEY

endpoint: http://localhost:8000/text

method: PUT

parameters: [{
        "key": "what_we_do",
        "value": "Lorem ipsum dolor sit amet."
    },
    {
        "key": "contact",
        "value": "name@example.com"
    }
]

## USERS
endpoint: http://localhost:8000/user

method: GET

parameters: -

endpoint: http://localhost:8000/user/ID

method: GET

parameters: int ID

endpoint: http://localhost:8000/user/username/USERNAME

method: GET

parameters: string USERNAME

endpoint: http://localhost:8000/user/ID

method: PUT

parameters: {
    "id": 1
    "username": "arualcms@arualcms.com",
    "password": "sdsdkjfhdskjfh4325234",
}

endpoint: http://localhost:8000/user

method: POST

parameters: {
    "username": "arualcms@arualcms.com",
    "password": "sdsdkjfhdskjfh4325234",
}

endpoint: http://localhost:8000/user/ID

method: DELETE

parameters: int ID

## FILES
endpoint: http://localhost:8000/files

method: GET

parameters: -

endpoint: http://localhost:8000/files/ID

method: GET

parameters: int ID

endpoint: http://localhost:8000/files/upload

method: POST

parameters: FormData with uploaded file

endpoint: http://localhost:8000/files/gallery/ID

method: PUT

parameters: int ID, {
    "gallery": "gallery 1",
}

endpoint: http://localhost:8000/files/ID

method: DELETE

parameters: int ID

## LANGUAGES
endpoint: http://localhost:8000/languages

method: GET

parameters: -


endpoint: http://localhost:8000/languages/KEY

method: GET

parameters: string KEY

endpoint: http://localhost:8000/languages

method: PUT

parameters: [{
"key": "en",
"value": "English"
},
{
"key": "cz",
"value": "Čeština"
}
]

## AUTHORIZATION
endpoint: http://localhost:8000/files/post

method: POST

parameters: {
    username: "registered email",
    password: "user password"
}

### Storage description
As a storage are use json files stored in database folder. For running arualCMS you not need any database server.
For getting information from database are used api endpoints. In arualCMS are defined this entities:
- posts
- texts
- users
- storage
- settings
With each entity you can read, create, edit and delete.

### VueJS environment
Admin part is using standart VueJS cli to build the application or run in development mode.
For development start dev server from admin folder: npm run serve
For production build after chnages use: npm run build
### Security
To secure communication between api and vuejs parst is using standard jwt token.
