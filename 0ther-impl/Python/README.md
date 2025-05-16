### xorfiles written in Python

Performs XOR (Exclusive or) operation on two files

DISCLAIMER: not an expert in Python (didn't find better implementation)

```sh
$ xorfiles.py {first-file | stdin} second-file [output-file]
```

Examples:

```sh
$ xorfiles.py first-file second-file output-file
$ cat first-file | xorfiles.py second-file output-file
$ xorfiles.py second-file output-file < first-file
$ cat first-file | xorfiles.py second-file > output-file
$ xorfiles.py first-file second-file > output-file
```

