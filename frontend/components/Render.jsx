export function render(name, size, raw, modified) {
    const backend = `http://localhost:${process.env.SERVER_PORT}`;
    let url = `${backend}/${raw}`;
    if (raw.match(/[a-zA-Z0-9]+/g)) {
        url = `/${raw}`;
    }

    return <a href={url} key={"directory"}>
        <p>{name}</p>
        <p>{size}</p>
        <p>{modified}</p>
    </a>
}

export async function Render({ data, url }) {
    return (
        <>
            <h2>{url}</h2>
            <div key={"directory"}>
                {data.map(obj => {
                    return render(obj.name, obj.size, obj.url, obj.modified);
                })}
            </div>
        </>
    );
}
