/* 
  => Errorgroup is professional way to run multiple go routine  + collect errors + cancell everything on first failure

  => Errorgroup gives you 

    1. start many question 
	2. wait for all 
	3. return first error
	4. cancel the first rest automatically
*/

package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)


func main(){

	var g errgroup.Group

	g.Go(func()  error {

		time.Sleep(1 * time.Second)

		fmt.Println("Task 1 done")

		return nil
		
	})


	g.Go(func() error {

		time.Sleep( 2 * time.Second) 

		fmt.Println("Task 2 done")

		return nil
		
	})


	if err := g.Wait(); err != nil {
		fmt.Println("Error : " , err)
	}

	fmt.Println("All Task finished")
}
