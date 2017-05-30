package main

import (
	"syscall"
	"fmt"
)

func main() {
	fmt.Printf("My PID: %b\n", syscall.Getpid());
	fmt.Printf("My PID: %c\n", syscall.Getpid());
	fmt.Printf("My PID: %d\n", syscall.Getpid());
	fmt.Printf("My PID: %o\n", syscall.Getpid());
	fmt.Printf("My PID: %q\n", syscall.Getpid());
	fmt.Printf("My PID: %x\n", syscall.Getpid());
	fmt.Printf("My PID: %X\n", syscall.Getpid());
	fmt.Printf("My PID: %U\n", syscall.Getpid());
	a := 12345.6789
	fmt.Printf("a: %f\n", a);
	fmt.Printf("a: %2.2f\n", a);
	fmt.Printf("a: %1.f\n", a);
	fmt.Printf("a: %.2f\n", a);
}
