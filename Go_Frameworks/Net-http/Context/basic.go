/* 
  => Context is a request - scoped container that carries 
     1. Cancellation signal 
	 2. Timeout deadline 
	 3. Values like (extra id)


  =>   Context = invisible backpack that travels with your request.	
  
 
  => Why do we need it?

Imagine:

client opens browser

sends request

then closes browser

Without context:

❌ your server keeps working
❌ database still running
❌ goroutines never stop

With context:

✅ everything stops automatically
  
*/

package main