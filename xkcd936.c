#define EX_SOFTWARE 70
#define EX_OSERR 71
#define WORD_LIST_LENGTH 234937

#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include <sodium.h>

extern char *words[WORD_LIST_LENGTH];

void get_random_numbers(uint32_t *random_numbers) {
    for (int i = 0; i < 4; i++) {
        random_numbers[i] = randombytes_uniform(WORD_LIST_LENGTH);
    }
}

int main() {
    int secure = sodium_init();
    if (secure == 0 || secure == 1) {
        uint32_t random_numbers[4];
        get_random_numbers(random_numbers);
        for (int i = 0; i < 4; i++) {
            char *word = words[random_numbers[i]];
            printf("%c", toupper(word[0]));
            for (size_t n = 1; n < strlen(word); n++) {
                printf("%c", word[n]);
            }
        }
        return 0;
    } else if (secure == -1) {
        fprintf(stderr, "sodium_init() FAILED!");
        return EX_OSERR;
    } else {
        fprintf(stderr, "BUG: sodium_init() will only return {-1, 0, 1}.");
        return EX_SOFTWARE;
    }
}
