package main

import "net"

func main() {
	ln, err := net.Listen("tcp", ":8000"); 
	if err != nil {
		panic(err) 
	} 

	for {
		conn, err := ln.Accept();  
		if err != nil {
			panic(err) 
		} 

		for {
			buf := make([]byte, 10)
			conn.Read(buf); 
		} 
	} 

} 