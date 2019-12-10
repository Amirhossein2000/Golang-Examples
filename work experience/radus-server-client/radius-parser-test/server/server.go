package main

import (
	"log"

	"parspooyesh.com/scm/ibsng/go-lib-radius/radius"
)

func main() {
	handler := func(w radius.ResponseWriter, r *radius.Request) {
		// username := rfc2865.UserName_GetString(r.Packet)
		// password := rfc2865.UserPassword_GetString(r.Packet)
		// class := rfc2865.Class_GetString(r.Packet)
		// log.Println(username, password, class)
		w.Write(r.Response(radius.CodeAccessAccept))
	}

	server := radius.PacketServer{
		Handler:      radius.HandlerFunc(handler),
		SecretSource: radius.StaticSecretSource([]byte(`secret`)),
	}

	log.Printf("Starting server on :1812")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
