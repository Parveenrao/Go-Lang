package gointerviewseries

/* 
   => What is zero value 
      
      In Go 
	     -> Every variable gets a default value automatically if u don't intialise it
		    This default is called Zero value 




    => Example 

	   var a int
       var b string
       var c bool
       var d float64
       var e *int
       var f []int
       var m map[string]int
       var s struct{}

	   fmt.Println(a) // 0
       fmt.Println(b) // ""
       fmt.Println(c) // false
       fmt.Println(d) // 0
       fmt.Println(e) // nil
       fmt.Println(f) // nil
       fmt.Println(m) // nil
       fmt.Println(s) // {}


	=> Struct Zero Value (VERY IMPORTANT)
       type User struct {
           Name string
          Age  int
          }

         var u User
         fmt.Println(u)   
    
		  
    => Zero value means 
	
	    ✔ No undefined memory
        ✔ Safe defaults
        ✔ No constructor required
        ✔ Cleaner code
        ✔ Predictable behavior

    => Zero vs Nil
	
	   ->   var p *int
              fmt.Println(p) // nil
       
	   -> s := []int{}
              fmt.Println(s == nil) // false

	 		

*/

