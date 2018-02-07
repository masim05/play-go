# Block the runtime
In the first terminal:
```bash
go run block/block.go
```
In the second terminal:
```bash
for i in `seq 16`; do echo exit | telnet localhost 3000; done
```
Several first connections will get response like `Connection #N` but the rest commands will
not due to blocked runtime.