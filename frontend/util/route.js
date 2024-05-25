export async function getData(dir) {
    if (dir === "/") {
        dir = "";
    }

    const backend = `http://localhost:${process.env.SERVER_PORT}${dir}`;
    const res = await fetch(backend);
    if (!res.ok) {
        const obj = res.json();
        throw new Error(obj.errno);
    }

    return res.json();
}
