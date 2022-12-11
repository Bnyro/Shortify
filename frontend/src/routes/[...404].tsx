import { Title } from "solid-start";
import { HttpStatusCode } from "solid-start/server";
import "./404.css"

export default function NotFound() {
  return (
    <main>
      <Title>Not Found</Title>
      <HttpStatusCode code={404} />
      <h1 id="error">404</h1>
      <h2>Page Not Found</h2>
      <button onClick={() => {window.location.href = "/"}}>Back To Home</button>
    </main>
  );
}
