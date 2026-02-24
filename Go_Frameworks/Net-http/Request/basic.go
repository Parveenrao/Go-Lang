package main

/*
  => In Every Go Handler:
    func(w http.ResponseWrite , r *http.Request)  --> r(everything the clinet sends)

	Methods , URL , query params , headers and body  --- all inside r

  => What inside http.Request
  
     1. r.Method(Get , Put ,  Post , Delete)
     2. URL Path (r.URL.Path)
	 
	 
	 3. Query Parameters /hello?name = Parveen&age = 22
	    
	    name := r.URL.QUERY().GET("name")
		age := r.URL.QUERY().GET("age")
	 
	 4. Headers 
	    
	   token := r.Header.Get("Authorization")
*/