export async function getData(dir, token) {
    if (dir === "/") {
        dir = "";
    }

    const backend = `http://localhost:${process.env.SERVER_PORT}/path${dir}`;
    const res = await fetch(backend, {
        headers: {
            "Authorization": token
        }
    });
    if (!res.ok) {
        if (res.status == 401) {
            return {
                status: 401
            };
        }

        return {
            status: 404
        };
    }

    return res.json();
}
