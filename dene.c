#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

#define PORT 11111

int main() {
   
    int sockfd;
    struct sockaddr_in server_addr;

    // Socket oluşturma
    if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
        perror("Socket hatası");
        exit(1);
    }

    // Sunucu adresinin ayarlanması
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(PORT);
    
    server_addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    memset(&(server_addr.sin_zero), '\0', 8);

    // Sunucuya bağlanma
    if (bind(sockfd, (struct sockaddr*)&server_addr, sizeof(server_addr)) < 0) {
        perror("Connection failed");
        exit(EXIT_FAILURE);
    }
    
    
    if (connect(sockfd, (struct sockaddr *)&server_addr, sizeof(struct sockaddr)) == -1) {
        perror("Bağlantı hatası");
        exit(1);
    }

    // Mesaj alma ve gönderme döngüsü
    for(int dx = 0; dx<3 ; dx++) {
        char message[1024];
        printf("İstemci 1'e mesaj gönderin: ");
        fgets(message, sizeof(message), stdin);
        send(sockfd, message, strlen(message), 0);

        // Sunucudan gelen mesajın alınması
        char received_message[1024];
        int bytes_received = recv(sockfd, received_message, sizeof(received_message), 0);
        received_message[bytes_received] = '\0';
        printf("İstemci 1'den gelen mesaj: %s", received_message);
    }

    // Soketin kapatılması
    close(sockfd);

    return 0;
}