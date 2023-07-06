import { connect, StringCodec } from "nats.ws";

const servers = ["ws://localhost:5222"];
// const servers = ["ws://nats.gondor.svc.kube:5222"];

async function main() {
  console.log("Getting started!!!!");
  const nc = await connect({
    servers: servers,
    debug: true,
  });
  console.log(`Connected to ${nc.getServer()}`);
  const codec = StringCodec();
  const sub = nc.subscribe("time");
  for await (const msg of sub) {
    const data = codec.decode(msg.data);
    document.getElementById("msg").innerHTML = `Received msg: ${data}`;
    console.log(`Received msg: ${data}`);
  }
  await nc.drain();
}

main();
