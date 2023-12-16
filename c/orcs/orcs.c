#include <stdio.h>
#include <stdlib.h>

char* munch(char* sentence){
    char* response = malloc(sizeof(char) * 100);
    if (response == NULL) {
        // Handle memory allocation failure if necessary
        return NULL; // Return NULL to indicate failure
    }

    char vowels[] = "aeiou"; // Can also use a string for simplicity
    // change vowels to X and copy non-vowels as is
    for (int i = 0; sentence[i] != '\0'; i++){
        response[i] = sentence[i]; // Copy character first
        for (int j = 0; j < 5; j++){
            if (sentence[i] == vowels[j]){
                response[i] = 'X'; // Replace with 'X' if it is a vowel
                break; // Exit the inner loop if a vowel is found
            }
        }
    }
    response[strlen(sentence)] = '\0'; // Null-terminate the response string

    return response;
}


int main() {
    char sentence[100];
    // read sentence
    fgets(sentence, 100, stdin);
    char* response = munch(sentence);
    if (response != NULL) {
        printf("INPUT: %s\n", sentence);
        printf("OUTPUT: %s\n", response);
        free(response); // Free the memory allocated for response
    } else {
        // Handle memory allocation failure if necessary
    }
    return 0;
}
