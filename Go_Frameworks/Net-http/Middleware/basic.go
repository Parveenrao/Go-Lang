package main 

/*  => Middleware is a function that run before and after your handler


    => Request Flow
        Client
          ↓
      Middleware
          ↓
    Your Handler (CRUD logic)
          ↓
    Middleware (after)
          ↓
       Response


	=> Middleware is used for 
	
	  ✅ logging
      ✅ authentication
      ✅ panic recovery
      ✅ timing requests
      ✅ adding headers




🔑 Core Idea (memorize this)

In Go:

middleware takes http.Handler
returns http.Handler


Signature:

func(next http.Handler) http.Handler	  */