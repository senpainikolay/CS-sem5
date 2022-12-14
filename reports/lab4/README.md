# Laboratory Work 4. Digital Signature.

### Course: Cryptography & Security
### Author: Nicolae Gherman

----

## Theory 
Digital signatures can provide evidence of origin, identity and status of electronic documents, transactions or digital messages.Similarly, a digital signature is a technique that binds a person/entity to the digital data. This binding can be independently verified by receiver as well as any third party. <br> It assures: Message authentication, Data Integrity and Non-repudiation.



## Objectives:

* Get familiar with digital signatures.
* Implement the Digital Signatures with assymetric ciphers and hashing techniques.

## Implementation description
I have decided to  simulate a simple Registration and Log In cases of an user to a simple in memory, map database. All that happens through POST requests on a http server. <br> 
So, the user registers and a token/string is generated and saved in the backend and sent as a response to the user. Once the user has to log in, they have to provide the token generated by the server in order  to be logged in (conceptually). Digital signatures happens once the provided token being hashed is compared to the   decrypted hash of the same saved token from database.

## The implementation is in db_simulation folder of the project.  

I will just get straight to the point on lab requirements: 
 
1. Hashing   

The build-in argon2 hash technique with given parameters.

``` 
// db_simulation/argon2/hash.go

func GetTheHashOnText(password string) []byte {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}.
	salt := []byte{113, 75, 156, 151, 17, 222, 194, 185, 42, 59, 119, 32, 22, 216, 81, 1}
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
return hash 
```

2. Database 

It has 2 maps. <br> One for storing the username and their bytes(hashed passwords). Another one for storng the username and their dencrypted hashed token.
```     
// db_simulation/database/db.go 

type Database struct {
	mapDB     map[string][]byte
	UserToken map[string][]byte
}
``` 

3. The users  

The Interfaces defined below explains the logic that happens on server side. 
The methods takes the username and password as input, addition token for LogIn method 

``` 
// db_simulation/client/interfaces.go 

type ClientSimulation interface {
	RegisterCredentials(string, string) string
	LogInCredentials(string, string, string) string
}

func RegisterInterface(ci ClientSimulation, usr string, pw string) string {
	return ci.RegisterCredentials(usr, pw)
}
func LogInInterface(ci ClientSimulation, usr string, pw string, tk string) string {
	return ci.LogInCredentials(usr, pw, tk)
}

```  
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




