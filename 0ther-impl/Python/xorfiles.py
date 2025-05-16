#!/usr/bin/env python3

import sys
import os

BUF = 524288

# Start processing CLI params
if len(sys.argv) < 2 or len(sys.argv) > 4:
    print("Usage:\n$ xorfiles.py {first-file | stdin} second-file [output-file]\n")
    sys.exit(22)

if len(sys.argv) == 2:
    first_file = sys.stdin.buffer
    second_file = open(sys.argv[1], 'rb')
    output_file = sys.stdout.buffer
elif len(sys.argv) == 3:
    if os.path.exists(sys.argv[2]):
        first_file = open(sys.argv[1], 'rb')
        second_file = open(sys.argv[2], 'rb')
        output_file = sys.stdout.buffer
    else:
        first_file = sys.stdin.buffer
        second_file = open(sys.argv[1], 'rb')
        output_file = open(sys.argv[2], 'wb')
else:
    first_file = open(sys.argv[1], 'rb')
    second_file = open(sys.argv[2], 'rb')
    output_file = open(sys.argv[3], 'wb')

if first_file is None or second_file is None or output_file is None:
    print("\nERROR! Can't open some specified files. Please check.\n")
    sys.exit(2)

while True:

    chunk1 = first_file.read(BUF)
    chunk2 = second_file.read(BUF)

    if not chunk1 or not chunk2:
        break

    length = min(len(chunk1), len(chunk2))

    result = bytes(a ^ b for a, b in zip(chunk1[:length], chunk2[:length]))

    if output_file.write(result) is None:
        print("\nERROR! Can't write data to the output-file (perhaps no space left). Exiting.", file=sys.stderr)
        sys.exit(5)

first_file.close()
second_file.close()
output_file.close()
