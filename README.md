![Screenshot 2021-03-25 at 13 56 24](https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png)


# Simple CMS made with love in VueJS and PHP


## Instalation
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
endpoint: 
method:
parameters:

endpoint: 
method:
parameters:


### Storage description
As a storage are use json files stored in database folder. For running arualCMS you not need any database server.
For getting information from database are used api endpoints. In arualCMS are defined this entities:
- posts
- texts
- users
- files
- settings
With each entity you can read, create, edit and delete.

### VueJS environment
Admin part is using standart VueJS cli to build the application or run in development mode.
For development start dev server from admin folder: npm run serve
For production build after chnages use: npm run build
### Security
To secure communication between api and vuejs parst is using standard jwt token.
