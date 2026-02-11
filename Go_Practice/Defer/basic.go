package main

/* 



// Defer tell go 
   => Run this function later , right before the surrounding function return 
   => So instead of executing immediately , deferred call is pushed onto stack  and executed after your function finshes


   
// Why does Go have Defer 
   => Resource Cleanup 
      1. Close filed 
	  2. Release lock 
	  3 . Close Db connection 
	
	=> Gurantee execution 
	   1. Even if function exit early
	   2. Even if panic happen
	
	=> Cleaner code
	   












   */