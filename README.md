# xk6-redis

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/dgzlopes/xk6-redis@latest
  ```

## Example

```javascript
import redis from 'k6/x/redis';

const client = redis.newClient();

export function setup() {
  redis.set(client,"snake","camel",0)
  redis.set(client,"foo",100,10)
}

export default function () {
  console.log(redis.get(client,"snake"))
  console.log(redis.get(client,"foo"))
  if (redis.do(client,"PING","bzzz") == "bzzz"){
    console.log("PONG!")
  }
}

export function teardown () {
  redis.del(client,"foo")
}

```

Result output:

```
$ ./k6 run script.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: ../example.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] bar                                           source=console
INFO[0000] PONG!                                         source=console

running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU

    █ setup

    █ teardown

    data_received........: 0 B 0 B/s
    data_sent............: 0 B 0 B/s
    iteration_duration...: avg=544.06µs min=428.6µs med=597.41µs max=606.18µs p(90)=604.43µs p(95)=605.31µs
    iterations...........: 1   46.10603/s
```
