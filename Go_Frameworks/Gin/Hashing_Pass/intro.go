package main 

/*   
   => Hashing Password means you never store real password in database 
      --> if db leaks , your passwords also leaks 
	  --> so we store a long hash passs (random - long unique every time)
	  
   => Hash Password 
      Cannot be reversed
	  Unique every time
	  safe to store 
	  
   => IN Go we use bcrypt 
     
      bcrypt does two things 

	  1. Turn password into hash 
	  2. Compare password <-> hash
     */