# Where do Syscalls Originate?

A tool to analyse which libraries are making syscalls

---
## Overview

In a dynamically linked executable, different libraries will make system calls. It is useful to understand which libraries are responsible for making different system calls when deciding how to best compartmentalise an application

---
## Approach

1. Spawn a process in a container with **ASLR disabled**
2. Look at the the virtual address map of the process using `/proc/<getpid()>/maps` to see which addresses correspond to which `.so` files
3a. Spawn a child process responsible for running the executable being analysed
3b. Use `strace` to monitor system calls while running the application
4. Using the memory map data to identify which calls are made from which libraries
5. Output statistics, including
    - Total system calls made
    - Number of different libraries that made system calls
    - For each library...
        - Number of syscalls made
        - Percentage of syscalls made
        - Count for each syscall made by the library

---
## Limitations

- **Dynamic Analysis**: will not capture every system call made by the application, missing especially those on error paths

