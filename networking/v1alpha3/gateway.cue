package v1alpha3 

import "list"

Servers: [...Server] & list.MinItems(1)

Server: {
    Hosts: [...host] & list.MinItems(1)

    host: "^(*|[-a-z-A-Z0-9]*[a-zA-Z0-9])(.[-a-z-A-Z0-9]*[a-zA-Z0-9])*$"

    Port: {
        number: uint32

        protocol: [ "HTTP" | "HTTPS" | "GRPC" | "HTTP2" | "MONGO" | "TCP" | "TLS" ]

        name?: string
    }
} 
