# Laboratory Work 5. Authentication and Authorization.

### Course: Cryptography & Security
### Author: Nicolae Gherman

----

## Theory 
Authentication and authorization are two vital information security processes that administrators use to protect systems and information. Authentication verifies the identity of a user or service, and authorization determines their access rights.



## Objectives:

* Get familiar with authorization and authentication.
* Implement Authentication and integrate Authorization on existing project components.

## Implementation description
----

## The implementation is in db_simulation folder of the project.  

* Registration(username,password) Method saves the hashed password on the database, generated a random string on function randStr(lenght), hashes it, encryptes by RSA  and returns the original string token as a response to the request of the user.

```   
// db_simulation/client/client.go 

hashedPassword := hashArgon2.GetTheHashOnText(password)
c.db.RegisterUser(username, hashedPassword) 

token := randStr(10)
hashedToken := hashArgon2.GetTheHashOnText(token)

en, err := rsa.EncryptOAEP(...)
if err != nil {
	log.Fatal(" RSA Encrytion failed ")
} 

c.db.UserToken[username] = en 

return token
``` 

* LogIn(username,password,token) Method  compares the hashed provided password with one saved in the database. <br>
Then it compares the hashed provided token with the decrypted hashed from the database. The Digital Signature itself.

``` 
// db_simulation/client/client.go 

hashedPassword := hashArgon2.GetTheHashOnText(password)
db_pass := c.db.GetUserPassword(username)
res := bytes.Compare(hashedPassword, db_pass)
if res != 0 {
	return "wrong password or user does not exist"
} 

hashedToken := hashArgon2.GetTheHashOnText(token)
de, _ := rsa.DecryptOAEP()
res = bytes.Compare(hashedToken, de)
if res != 0 {
	return "Unsucessful Authentication"
}

return "Success Log In"
```

4. The server 

You have to provide the data as routes variables.
``` 
// db_simulation/server/httpServer.go 

func RunDBSimulationServer() {
	r := mux.NewRouter()
	r.HandleFunc("/register/{usr}/{val}", RegisterUser).Methods("POST")
	r.HandleFunc("/login/{usr}/{val}/{token}", LogInUser).Methods("POST")
	http.ListenAndServe(":8080", r)
}
```

---  


## The proof will be Real Time presentation  :-). Cheers!




