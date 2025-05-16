### xorfiles

Performs XOR (Exclusive or) on two files
Implemented in different languages with max. performance as a priority

```sh
$ xorfiles {first-file | stdin} second-file [output-file]
```

Examples:

```sh
$ xorfiles first-file second-file output-file
$ cat first-file | xorfiles second-file output-file
$ xorfiles second-file output-file < first-file
$ cat first-file | xorfiles second-file > output-file
$ xorfiles first-file second-file > output-file
```

Notice! Last ex. will work as expected if second-file exists. If not, will processed like ex.2
As in second-file = output-file and user input = STDIN (interactive, use ENTER and CTRL+D to send EOF)

### Performance tests
Ryzen 5 5500U, Ubuntu 22.04, **RAM Disk** to minimize I/O latency

Throughput measured with pv tool, launched 5 times and make the **average**

```sh
$ time -v impl-exec first-file second-file | pv > /dev/null
```

Both files 2GB each. Please, review *0ther-test/Makefile* for details

---
| Implementation *    | Throughput, MiB/s | CPU util., % | Real, s | User, s | System, s |
|:-------------------:|:-----------------:|:------------:|:-------:|:-------:|:----------|
| Go (2 goroutines)   | 1447.25           | 127          | 1.41    | 0.6     | 1.19      |
| C (1 thread)        | 1358.5            | 76           | 1.5     | 0.2     | 0.95      |
| Go (1 goroutine)    | 1314.13           | 79           | 1.55    | 0.32    | 0.92      |
| PHP 8               | 574.16            | 88           | 3.6     | 2.21    | 0.99      |
| Python 3            | 16.7              | 99           | 122.5   | 119     | 2.53      |

**\* To review implementations, please check *0ther-impl* folder**

