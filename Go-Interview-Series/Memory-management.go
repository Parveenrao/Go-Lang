package gointerviewseries

/*
   => Go manages memory automatically using heap , stack  and garbage collector 

   => Go use three main pieces
      1. Stack
	  2. Heap 
	  3. Garbage collector

	
    => 1. Stack memory 
	   -> Local variables that don't escape the function go on the stack

	   func add() {
	       x:= 10
		   y:= 20
		   z:= x + y
		   fmt.println(z)
		   }
        here x , y, x lives on the stack 
		when add function is returns -- stack frame is destoryed
		memory is freed automatically
	
	Stack memory is automatic and not handled by GC

	=> 2. Heap memory 
	      -> if a variable escape the function , go puts it on the heap 

		  func getNumber() *int {
              x := 10
              return &x
            }
        x is returned as a pointer 
		so it must survive after function exist 
		Go moves x to the heap 

		This decision is made by using escape analysis

	=> 3. Escape Analysis (compiler magic)
            -> Go compiler checks:
            -> Will this variable be used outside this function
            -> If YES → heap
	        -> If NO → stack
            -> You don’t control this manually.

    => 4. Garbage Collector 

	      -> Heap objects are  cleaned by Go GC
		  
	  GC Does:
	  -> Finds reachable objects( from stack + global)
	  -> Delete unreachable ones

	  
	 -> Example
	     package main

        import "fmt"

        type User struct {
	     Name string}

		 func main() {
	        u := &User{Name: "Ram"}
	       fmt.Println(u.Name)

	       u = nil
            }

		here 
		1. &user object in memory(heap)
		   u points to it (Memory in use)

		2. fmt.Println(u.Name)
		   Everything is normal 
		
		3. u = nil
		   reference removed
		   
		   but user object still in memory and no one is pointing it to anymore

		   this is called unreachable
		
		
		4.  Garbage Collector runs

                Later (automatically), GC wakes up and asks:

               “Which objects are still reachable from variables?”

                It checks:

                stack variables

                global variables   
                   
				So GC says:

               This is garbage.

               Then:

              👉 memory is freed.

              You did NOTHING.

              That’s Go GC.

			 Create object
                 ↓
             Use object
                 ↓
             Remove reference
                 ↓
             GC detects unreachable memory
                 ↓
             GC frees memory



*/