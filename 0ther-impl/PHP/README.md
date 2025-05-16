### xorfiles written in PHP

Performs XOR (Exclusive or) operation on two files

```sh
$ xorfiles.php {first-file | stdin} second-file [output-file]
```

Examples:

```sh
$ xorfiles.php first-file second-file output-file
$ cat first-file | xorfiles.php second-file output-file
$ xorfiles.php second-file output-file < first-file
$ cat first-file | xorfiles.php second-file > output-file
$ xorfiles.php first-file second-file > output-file
```

Notice! Last example will work as expected if second-file exists

If not, will processed like ex.2 as in second-file = output-file and user input = STDIN (interactive, use ENTER and CTRL+D to send EOF)
