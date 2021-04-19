# User-Guide

Karla Sofía González Rodríguez 0214774

Saúl Eduardo Zepeda de la Torre 0214016



## Sign in
First of all, the server needs to be started, to do so, open a new terminal and check that you are in the correct directory (dc-labs/challenges/second-partial) to run the program. 

`$ cd dc-labs/challenges/second-partial`

Once you are in the correct directory, start the server by running the next command. This will run the server in localhost in the port 8080

`$ go run main.go`

To start using the API, you need to open another terminal. It also needs to be in the same location of  dc-labs/challenges/second-partial, so run again in the new terminal the command:

`$ cd dc-labs/challenges/second-partial`

Once you are in the second terminal and in the correct directory, you will need to register with the username and password that you want

`curl -u <USERNAME>:<PASSWORD> localhost:8080/signin`

>If this is an original username and you do not have any problems, you will recive the next message:{ "message": "Hi <-USERNAME-> your user has been created."}

> If the username was already created it will give you an error, please type another username. { "message": "Error, your username already exists." }

## Login
Once you created your account, it is time to login. Please have your username and your password, since it is going to be needed. Run the next command:

`curl -u <USERNAME>:<PASSWORD> localhost:8080/login`

>If the user and the password are correct, the following message will appear and it will generate a secret token that you will need to remember in order to do certain things in the program:  { "message": "Hi <-USERNAME->, welcome to the DPIP System", "token" <-TOKEN-> }

>If the user is not correct or it is not registered, it will display: { "message": "The username is not registered" } 

>Or if the password is incorrect, this message will appear: { "message": "The password is incorrect" }

## Status
If you logged in but you want to doble check that you logged in correctly, you can see that with the next line (remember to have your token) :

`curl -H "Authorization: Bearer <TOKEN>" localhost:8080/status`

>If the token exists and there are no mistakes, the following message will be shown: { "message": "Hi <-USERNAME->, the DPIP System is Up and Running" "time": <-DATE_TIME-> }

>If the token is incorrect and there is no user with that token, then it will display this message { "message": "Error, you have to login." }

## Upload
Once you logged in and there are no problems with your status, then you can upload an image using the following command (remember to have your token):

`curl -F 'data=@test.jpg' -H "Authorization: Bearer  <TOKEN>" localhost:8080/upload`

>If the image was uploaded correctly a similar message will appear: { "message": "An image has been successfully uploaded", "filename": <-FILE_NAME->, "size": <-SIZE_IMG->}

>If not, then the next message will appear: { "message": "Error uploading image", "filename": <-FILE_NAME->}

>Or if the token is incorrect and there is no user with that token, then it will display this message { "message": "Error, you have to login." }

## Logout
 If you want to logout you can do that, however your token that was generated in the login is going to be needed. Run the command:

`curl -H "Authorization: Bearer  <TOKEN>" localhost:8080/logout`

>With this action the information of the user will be deleted, and it will show the next line: { "message": "Bye <-USERNAME->, your token has been revoked" }
