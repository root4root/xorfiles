### xorfiles written in Go, no extra goroutines

Performs XOR (Exclusive or) operation on two files
Instead of a first file, could be used STDIN
If output-file not specified - print to STDOUT

```sh
$ xor-one {first-file | stdin} second-file [output-file]
```

Examples:

```sh
$ xor-one first-file second-file output-file
$ cat first-file | xor-one second-file output-file
$ xor-one second-file output-file < first-file
$ cat first-file | xor-one second-file > output-file
$ xor-one first-file second-file > output-file
```

Notice! Last example will work as expected if second-file exists

If not, will processed like ex.2 as in second-file = output-file and user input = STDIN (interactive, use ENTER and CTRL+D to send EOF)
