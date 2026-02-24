// Append into file 

package main

import "os"

func main() {
	file, err := os.OpenFile("data.txt",
		os.O_APPEND|os.O_WRONLY,
		0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("New Line\n")
}


/* 
🧠 File Flags (VERY IMPORTANT)

Flag	        Meaning

os.O_RDONLY	    read only
os.O_WRONLY	    write only
os.O_RDWR	    read + write
os.O_APPEND	    append
os.O_CREATE	    create if not exists
os.O_TRUNC	    truncate
*/

/*  

writer := bufio.NewWriter(file)
writer.WriteString("Buffered write\n")
writer.Flush()

 */