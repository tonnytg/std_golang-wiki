package main

/*
# include <stdio.h>
# include <stdlib.h>
# include <stdint.h>
typedef struct Message {
    uint8_t     m_type;
    uint32_t    buff_size;
} message_t;
char * gl_msg = "message_in_c_global";
void receiver(message_t* msg) {
    printf("%d \t %d \n", msg->m_type, msg->buff_size);
}
message_t* getMessage() {
    message_t *msg;
    msg = (message_t*)malloc(sizeof(message_t));
    msg->m_type = 1;
    msg->buff_size = 1024;
    return msg;
}
void hello() {
    printf("hello from C\n");
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type (
	Message struct {
		MType    uint8
		BuffSize uint32
	}
)

func main() {

	var (
		msg string

		msgC  *C.message_t
		msgGo *Message
	)

	C.hello()

	// Convert C char pointer to Go String
	msg = C.GoString(C.gl_msg)
	fmt.Println("Converted C string to Go string", msg)

	// Fetch C struct of same data type
	msgC = C.getMessage()

	// Convert the C pointer using unsafe and convert it back
	// to the Go Message struct pointer
	msgGo = (*Message)(unsafe.Pointer(msgC))

	fmt.Println("Converted C struct to Go struct of similar data types",
		msgGo.MType, msgGo.BuffSize)

	// Change the struct variable and cast it back to C struct
	msgGo.BuffSize = 4096

	// Convert the msgGo pointer to unsafe pointer
	// and pass it back to the C receiver method
	C.receiver((*C.message_t)(unsafe.Pointer(msgGo)))

}
