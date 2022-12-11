import { createSignal } from "solid-js";

export default function Shortener() {
    const [url, setUrl] = createSignal("");
    const [short, setShort] = createSignal("");

    const isValidUrl = (urlString: string) => {
        try { 
            return Boolean(new URL(urlString)); 
        }
        catch(e){ 
            return false; 
        }
    }

    const onClick = async () => {
        if (!isValidUrl(url())) {
            alert("Invalid URL");
            return;
        }

        const response = await fetch("http://localhost:8002/create", {
            method: "POST",
            body: JSON.stringify(
                {
                    full: url(),
                }
            )
        });
        const json = await response.json();
        if (json.message) alert(json.message);
        else setShort(json.short)
    }

    return (
        <section>
            <div>
                <input placeholder="URL" type="text" value={url()} onInput={(e) => setUrl(e.currentTarget.value)} />
                <button onClick={onClick}>Go</button>
            </div>
            <a id="short" href={short()}>{short()}</a>
        </section>
      );
}