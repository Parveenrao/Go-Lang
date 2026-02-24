 // Go does not have while and do while loop  , Everything is done using for

 // Classis for loop 

package main
import "fmt"

func  main() { 
	for i:=1; i<=5; i++{
		fmt.Println(i)
	}	
	// while style loop we can do in go 

	j:= 1

	for j<=5{
		fmt.Println(j)
		j++
	}

    // infine lopp 

	for {
    fmt.Println("Running...")
    }   
	// break 
    for i:=1; i<=10; i++ { 
		  if i==5 {
		   break
		      }
        fmt.Println(i)
	    }


	for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }
    fmt.Println(i)
}

    for i := 1; i <= 3; i++ {
    for j := 1; j <= 2; j++ {
        fmt.Println(i, j)
    }
}


}

