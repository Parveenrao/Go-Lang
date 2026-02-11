package main

/*  

=> Philosopy of Error Handling in Go 

   Go does not use try - catch block 

   Instead:

          Function return error explicitly
		  You must check error manualy 
     
    result, err := someFunc()
    if err != nil {
    // handle error
    }




*/