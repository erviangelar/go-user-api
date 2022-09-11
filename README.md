# go-user-api

Step Running : <br>
1. clone this repo with git <br>
2. Open with code editor <br>
3. run in terminal go run main.go or using air (live reload) <br>
4. open http://localhost:3000/healthcheck in your browser <br>
5. if you see "{"data":"Server is up and running"}" so go api running successfully<br>
<br><br>
Documentation :
- Swagger Doc can be accessed using : http://localhost:3000/swagger/index.html <br><br>

Api Contains :<br>
1. Healcheck : http://localhost:3000/healthcheck (mostly used for Docker)<br>
2. Login : http://localhost:3000/api/login<br>
    Param Body: <br>
    - Username<br>
    - Password<br>
3. Register : http://localhost:3000/api/register<br>
    Param Body : <br>
    - Name<br>
    - Username<br>
    - Password<br>
4. RefreshToken : http://localhost:3000/api/refresh-token <br>
    Header Authorize replace token with refresh token for generate new token<br>
5. User : http://localhost:3000/api/users<br>
6. User : http://localhost:3000/api/user/{id}<br>
    Path Variable {id}<br>

<br>
Default User :<br>
- Admin : <br>
    username : admin<br>
    password : admin<br>
- User : <br>
    username : user<br>
    password : user<br>
