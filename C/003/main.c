# include <stdio.h>
# include <stdlib.h>
# include <stdint.h>

typedef struct Message {
    uint8_t     m_type;
    uint32_t    buff_size;
} message_t;

char * gl_msg = "message_in_c_global";

void receiver(message_t* msg) {
    printf("Result: %d \t %d \n", msg->m_type, msg->buff_size);
}

message_t* getMessage() {
    message_t *msg;
    msg = (message_t*)malloc(sizeof(message_t));
    msg->m_type = 1;
    msg->buff_size = 1024;
    return msg;
}

void hello() {
    printf("Message: hello from C\n");
}
