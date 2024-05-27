import { notFound } from "next/navigation";

export async function getData(dir) {
    if (dir === "/") {
        dir = "";
    }

    const backend = `http://localhost:${process.env.SERVER_PORT}${dir}`;
    const res = await fetch(backend);
    if (!res.ok) {
        return notFound();
    }

    return res.json();
}
