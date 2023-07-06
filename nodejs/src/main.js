import { connect, credsAuthenticator, StringCodec } from "nats.ws";

const apiServer = "http://localhost:8833/session";
const servers = ["ws://localhost:5222"];
// const servers = ["ws://nats.gondor.svc.kube:5222"];
const codec = StringCodec();
const userId = "client";

async function main() {
  console.log("Getting started!!!!");

  const creds = await fetch(`${apiServer}/${userId}`).then((r) => r.text());
  const authenticator = credsAuthenticator(codec.encode(creds));
  const nc = await connect({
    servers,
    authenticator,
    debug: true,
  });

  console.log(`Connected to ${nc.getServer()}`);
  const sub = nc.subscribe(`${userId}.time`);
  for await (const msg of sub) {
    const data = codec.decode(msg.data);
    document.getElementById("msg").innerHTML = `Received msg: ${data}`;
    console.log(`Received msg: ${data}`);
  }
  await nc.drain();
}

main();
