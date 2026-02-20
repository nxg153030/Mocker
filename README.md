# DIY Container in Go
This is a project to understand how containers (like Docker) work under the hood. Instead of using high-level tools, I'm using the Go programming language to talk directly to the Linux Kernel.

## What this code does
This program creates a "restricted bubble" for a process to run in. It uses three main Linux features to create isolation:

UTS Namespace: Isolates the Hostname. Changing the name inside the container doesn't affect the host machine.

PID Namespace: Isolates Process IDs. Inside the container, the program thinks it is the very first process (PID 1).

Chroot: Isolates the Filesystem. The container is "jailed" in a specific folder (/mini-root) and cannot see or touch files on the rest of the computer.

## How to run it
Prerequisites
Linux: This code uses Linux-specific system calls. On a Mac, you must run this inside a Linux Virtual Machine or a Dev Container.

Privileged Access: Creating namespaces requires root permissions.

## Steps
Prepare the mini-root:
Create a folder named mini-root and place a statically compiled shell (like BusyBox) inside /mini-root/bin/sh.

Run the program:

Bash
sudo go run main.go /bin/sh
Lessons Learned
Containers aren't VMs: They are just normal processes with "blinkers" on, provided by the Linux Kernel.

Dependencies matter: If you chroot into a folder, you must ensure all the libraries a program needs are inside that folder, or use a static binary.

Go vs Python: Go provides a very direct, thin wrapper around system calls, making it great for learning how operating systems actually work.