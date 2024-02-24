import { Title } from "@solidjs/meta";
import Shortener from "~/components/Shortener";

export default function Home() {
  return (
    <main>
      <Title>Shortify</Title>
      <h1>Shortify</h1>
      <Shortener />
    </main>
  );
}
