![Screenshot 2021-03-25 at 13 56 24](https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png)


# Simple CMS made with love in VueJS and Go lang


## Installation
### Production docker container

On docker hub is located ajandera/arualCMS container which contains api and administration app in vuejs. Simply add it to your docker configuration to run you application based on arualCMS.

### Use Docker for development
Copy content of production folder to your destination and run **docker-compose up**, it will create container and app will be running on next urls:

Application has predefined credentials to be ready to use. Go to do admin url and insert

- username: arualcms@arualcms.com
- password: Password123

!! After first login change the password !!

### Use without docker
In that case you have run apache on your server or hosting and have Go lang 1.17+ installed. Then copy the content of production folder to your destination and go to urls as usual.

### Development
For development clone repository to your local machine and simple run docker-compose up from root folder. It creates all needed containers and run VueJS app in development mode. Application will be found on next ports:

- api http://localhost:8000
- admin http://localhost:8080
- frontend http://localhost:9000

### API documentation
https://arualcms.docs.apiary.io/#

### Storage description
As a storage is used Postgres database.
For getting information from database are used API endpoints. 
In arualCMS are defined this entities:
- posts
- texts
- users
- files
- settings
- languages
With each entity you can read, create, edit and delete.

### VueJS environment
Admin part is using standard VueJS cli to build the application or run in development mode.
For development start dev server from admin folder: npm run serve
For production build after changes use: npm run build

### Security
To secure communication between API and VueJS part is using standard JWT token.
