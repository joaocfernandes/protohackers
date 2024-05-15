#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/in.h>

#define PORT 8080
#define MAX_CONNECTIONS 10
#define BUFFER_SIZE 1380

void handle_client(int client_socket) {
  char buffer[BUFFER_SIZE+1];
  ssize_t bytes_received = 0;

  buffer[BUFFER_SIZE] = '\0';
  printf("Handling client.\n");


  while ((bytes_received = recv(client_socket, buffer, BUFFER_SIZE, 0)) > 0) {
    send(client_socket, buffer, bytes_received, 0);
  }

  close(client_socket);
}

int main() {
  int server_socket, client_socket;
  struct sockaddr_in server_address, client_address;
  socklen_t client_address_length = sizeof(client_address);

  // Create socket
  server_socket = socket(AF_INET, SOCK_STREAM, 0);
  if (server_socket == -1) {
    perror("Failed to create socket");
    exit(EXIT_FAILURE);
  }

  // Set server address
  server_address.sin_family = AF_INET;
  server_address.sin_addr.s_addr = INADDR_ANY;
  server_address.sin_port = htons(PORT);

  // Bind socket to address
  if (bind(server_socket, (struct sockaddr *)&server_address, sizeof(server_address)) == -1) {
    perror("Failed to bind socket");
    exit(EXIT_FAILURE);
  }

  // Listen for connections
  if (listen(server_socket, MAX_CONNECTIONS) == -1) {
    perror("Failed to listen for connections");
    exit(EXIT_FAILURE);
  }

  printf("Server listening on port %d\n", PORT);

  while (1) {
    // Accept incoming connection
    client_socket = accept(server_socket, (struct sockaddr *)&client_address, &client_address_length);
    if (client_socket == -1) {
      perror("Failed to accept connection");
      exit(EXIT_FAILURE);
    }

    // Handle client connection concurrently
    if (fork() == 0) {
      close(server_socket);
      handle_client(client_socket);
      exit(EXIT_SUCCESS);
    }

    close(client_socket);
  }

  close(server_socket);

  return 0;
}
