# Email Sender

# Service features:
 1. Sending newsletters using html layout and list of subscribers.
 2. Sending delayed mailings.
 3. Using variables in the mailing layout. (Example: first name, last name, birthday from the list of subscribers)
 4. Tracking email opens.
Implement deferred sends using Celery (relevant only for implementation in Python).

# What did I add? 
1. Make it more universal:
        Using this api everyone can send email to others by registering and providing their emails 
    password, not main password of email. You can get the password by following the steps.
    1. Enable Two step verification in your email
    2. Just below of the two step verification, There is part where you can add password.
 
2. Add swagger documentation:
        To make testing api easier, I added this feature.
  
3. Add docker:
        To make it easy to run my code I added docker. 
   
4. Casbin / middleware / JWT:
        To know if the user is registered or not and controll based on roles.

5. Migrations:
        To make it easy to create tables in database
        
 
# Databases 
1. Redis   (NOSQL)
  I used redis to store users information. I used redis just to show that I know redis and this project is just test.
2. Postgres  (SQL)
  I used postgres to store messages to be sent later.

# Running code  
    As I used docker everything in the code will run in docker,  so just few commands are enough to run the docker.
    
```
docker compose up
```
 ! If there's problem run the second command 
  
```
docker start api cron_job migrate-app
```
When this problem may arise. When your docker machine tries to run api and cron_job before postres or redis is run.
Even if I use "links" in docker-compose.yml file, sometimes this problem occurs.

After application run successfully, type the following location in your browser.
```
http://localhost:9090/swagger/index.html#/
```

You have to register first and put the given JWT token to Authorization header.

# Privacy
As you run this code in your local machine don't worry about privacy.
