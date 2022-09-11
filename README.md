# go-user-api

Documentation :
- Swagger Doc can be assecced using : http://localhost:3000/swagger/index.html <br><br>

Api Contains :<br>
    - Healcheck : http://localhost:3000/healtcheck (mostly used for Docker)<br>
    - Login : http://localhost:3000/api/login<br>
        Param Body: <br>
        - Username<br>
        - Password<br>
    - Register : http://localhost:3000/api/register<br>
        Param Body : <br>
        - Name
        - Username
        - Password
    - RefreshToken : http://localhost:3000/api/refresh-token <br>
        Header Authorize replace token with refresh token for generate new token<br>
    - User : http://localhost:3000/api/users<br>
    - User : http://localhost:3000/api/user/{id}<br>
        Path Variable {id}<br>

<br>
Default User :
- Admin : <br>
    username : admin<br>
    password : admin<br>
- User : <br>
    username : user<br>
    password : user<br>
