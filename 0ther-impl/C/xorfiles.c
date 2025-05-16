#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
#include <fcntl.h>
#include <unistd.h>

#define MYBUFSIZ 524288
#define STDIN 0  //POSIX standard file descriptor
#define STDOUT 1 //POSIX standard file descriptor

#define CREATE_MASK 0640

int main(int argc, char **argv)
{
    int first_file = STDIN;
    int second_file;
    int result_file = STDOUT;

    /// Start processing CLI params ---->
    if (argc < 2 || argc > 4) {
        fprintf(stderr, "\nUsage:\n$ xorfiles {first-file | stdin} second-file [output-file]\n\n");
        exit(22);
    }

    switch (argc) {
        case 2:
            second_file = open(argv[1], O_RDONLY);
            break;
        case 3:
            if (access(argv[2], F_OK) == 0) {
                first_file = open(argv[1], O_RDONLY);
                second_file = open(argv[2], O_RDONLY);
            } else {
                second_file = open(argv[1], O_RDONLY);
                result_file = open(argv[2], O_RDWR | O_CREAT, CREATE_MASK);
            }
            break;
        default:
            first_file = open(argv[1], O_RDONLY);
            second_file = open(argv[2], O_RDONLY);
            result_file = open(argv[3], O_RDWR | O_CREAT, CREATE_MASK);
    }

    if (first_file < 0 || second_file < 0 || result_file < 0) {
        fprintf(stderr, "\nERROR! Can't open some specified files. Please check\n\n");
        exit(2);
    } /// <---- End processing CLI params

    uint8_t *first_buf, *second_buf, *result_buf;
    ssize_t first_amount, second_amount, min_amount;

    first_buf = malloc(MYBUFSIZ);
    second_buf = malloc(MYBUFSIZ);
    result_buf = malloc(MYBUFSIZ);


    for (;;) {

        first_amount = read(first_file, first_buf, MYBUFSIZ);
        second_amount = read(second_file, second_buf, first_amount);

        min_amount = first_amount < second_amount ? first_amount : second_amount;

        if (min_amount <= 0) { break; }

        if ((min_amount & 0b111) > 0) {
            for (int i = 0; i < min_amount; ++i) {
                result_buf[i] = first_buf[i] ^ second_buf[i];
            }
        } else {
            for (int i = 0; i < min_amount; i += 8) {
                *((uint64_t*)(result_buf + i)) = *((uint64_t*)(first_buf + i)) ^ *((uint64_t*)(second_buf + i));
            }
        }

        if (write(result_file, result_buf, min_amount) < 0) {
            fprintf(stderr, "\nERROR! Can't write data to the output-file (prehaps none space left). Exiting.\n\n");
            exit(5);
        }
    }

    free(first_buf);
    free(second_buf);
    free(result_buf);

    close(first_file);
    close(second_file);
    close(result_file);

    return 0;
}
