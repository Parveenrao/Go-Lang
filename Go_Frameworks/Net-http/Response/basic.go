/* 
   If r is what client send then w(http.ResponseWriter) is what you send back 

   => Set headers   w.Header().Set("Content-Type", "application/json")

   => Status code w.WriteHeader(http.StatusCreated) // 201

  
   */
package main