
package main

/* 
  => After login server knows --> user_Id = 1

     but next request Get/profile  --> so server does not who is calling 

  => Http server is stateless 
     1. Every request is independent
	 2. Server forget everything

  => So without jwt 
     
     Login works      
	 but next request = anonymous again


  => Old session  uses jwt 
     1. server side sessionpackage 
	 2. cookies
	 3. stored login state in memory 
	 
   Problems 

   1. Not scalable
   2. hard with microservices 
   3 . memory heavy


   Moder api avoid this 



  => JWT --> JSON WEB TOKEN (proof of login)
  
    Its like a digital ID card

	-> Example 
	  You go to office 

	  Secuirty checks you and give to ID card
	  You show ID card everywhere 
	  Nodbody re-checks your passport

	  JWT is that ID Card
  
	  
	=> Inside JWT
	
	  (user_id = 1
	  expiry = tommorrow)

	  Signed by server signature 
	  client stores it 


	=> Every Request 

	 Authorization : <JWT> 

	 Server :

	 verifies signature 
	 reads  user_id 
	 allows access

	No Db lookup 



🔁 Full flow (important)
Login → server creates JWT
↓
Client stores JWT
↓
Client sends JWT with every request
↓
Middleware verifies JWT
↓
Request allowed


🧠 Why JWT is powerful
1️⃣ Stateless

Server stores NOTHING.

Token carries user identity.

2️⃣ Scalable

100 servers can verify same token.

No shared session memory.

3️⃣ Secure

Signed + expiration.

Tampering breaks signature.

4️⃣ Fast

No DB call per request.



	  
*/