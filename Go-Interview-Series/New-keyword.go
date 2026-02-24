package gointerviewseries

/* 
  New keyword in Go  

  1. Allocates memory for a type 
  2. Initilizes it with zero value 
  3. Returns a pointer to that memory 

     p := new(int)
	 fmt.Println(*p)   //0

	 what happened 

	 Memory allocated for int

     Value set to zero (0)

     Returned pointer to that memory


  => Escape Analysis 
    
     type User struct {
        Name string
        Age  int
      }

      u := new(User)

     fmt.Println(u.Name) // ""
     fmt.Println(u.Age)  // 0




*/