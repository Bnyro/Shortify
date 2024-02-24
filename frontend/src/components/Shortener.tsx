"use client";
import { createSignal, Show } from "solid-js";
import copy from "~/util/clipboard";
import "./Shortener.css";

export default function Shortener() {
    const [url, setUrl] = createSignal("");
    const [short, setShort] = createSignal("");
    const [cpImg, setCpImg] = createSignal("/clip.svg");
    const BASE_URL = import.meta.env.VITE_API_URL;

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

        const response = await fetch(`${BASE_URL}/create`, {
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

    const onCopy = () => {
        copy(short());
        setCpImg("/done.svg");
        setTimeout(() => {
            setCpImg("/clip.svg");
        }, 2000);
    }

    return (
        <section>
            <div>
                <input placeholder="URL" type="text" value={url()} onInput={(e) => setUrl(e.currentTarget.value)} />
                <button onClick={onClick}>Go</button>
            </div>
            <div id="result">
                <a id="short" href={short()}>{short()}</a>
                <Show when={short()}>
                    <button id="copy" onClick={onCopy}>
                        <img src={cpImg()} alt="Copy" />
                    </button>
                </Show>
            </div>
        </section>
      );
}
