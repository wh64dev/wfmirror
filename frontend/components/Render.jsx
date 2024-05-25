import Link from "next/link";

export function render(name, size, raw, modified, type) {
    const backend = `http://localhost:${process.env.SERVER_PORT}`;
    let url = `${backend}/${raw}`;
    if (type === "dir") {
        url = `${raw}`;
    }

    return <Link href={url} key={"directory"}>
        <p>{name}</p>
        <p>{size}</p>
        <p>{modified}</p>
    </Link>
}

export async function Render({ data, url }) {
    return (
        <>
            <h2>{url}</h2>
            <div key={"directory"}>
                {data.map(obj => {
                    return render(obj.name, obj.size, obj.url, obj.modified, obj.type);
                })}
            </div>
        </>
    );
}
