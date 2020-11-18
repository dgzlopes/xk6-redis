import redis from 'k6/x/redis';

const client = redis.newClient();

export function setup() {
  redis.set(client,"foo","bar",10)
}

export default function () {
  console.log(redis.get(client,"foo"))
  if (redis.do(client,"PING","bzzz") == "bzzz"){
    console.log("PONG!")
  }
}

export function teardown () {
  redis.del(client,"foo")
}
