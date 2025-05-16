### xorfiles written in C

Performs XOR operation between two files. Instead of a first file, could be used STDIN. If output-file not specified - print to STDOUT

```sh
$ xorf {first-file | stdin} second-file [output-file]
```

Examples:

```sh
$ xorf first-file second-file output-file
$ cat first-file | xorf second-file output-file
$ xorf second-file output-file < first-file
$ cat first-file | xorf second-file > output-file
$ xorf first-file second-file > output-file
```

Notice! Last example will work as expected if second-file exists

If not, will processed like ex.2 as in second-file = output-file and user input = STDIN (interactive, use ENTER and CTRL+D to send EOF)

