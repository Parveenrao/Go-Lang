package main

/* Map => A map is a built-in reference type that store key value pair
          key --> hash---> bucket---> value
		  access is o(1) time average 

		  
=> Internal Structure of Map 
   think of GO map as a  -----> Runtime-managed hash table made of buckets pointer and metadata


map variable
   ↓
map header (small struct)
   ↓
pointer to bucket array
   ↓
buckets → entries (keys + values)


1. Map Header  
   



*/