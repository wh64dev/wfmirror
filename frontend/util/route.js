export async function getData(dir) {
    const backend = `http://localhost:${process.env.SERVER_PORT}/${dir}`;
    const res = await fetch(backend);
    if (!res.ok) {
        throw new Error("error!");
    }

    return res.json();
}
