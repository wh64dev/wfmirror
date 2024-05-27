import Link from "next/link";
import style from "./render.module.scss";

const entry = ({ name, size, raw, modified, type }) => {
    return {
        raw: raw,
        name: name,
        size: size,
        type: type,
        modified: modified
    };
};

export function render(obj, lock, key) {
    const backend = `http://localhost:${process.env.SERVER_PORT}`;
    console.log(obj);
    let url = `${backend}/${obj.raw}`;
    if (obj.type === "dir") {
        url = `${obj.raw}`;
    }

    let symbol = "";
    if (lock) {
        symbol = ` bi bi-lock-fill`;
    }

    return <Link className={style.entry} href={url} key={key}>
        <i className={`${style.lock}${symbol}`}></i>
        <p className={style.name}>{obj.name}</p>
        <p className={style.size}>{obj.size}</p>
        <p className={style.modified}>{obj.modified}</p>
    </Link>
}

export async function Render({ data, url }) {
    return (
        <>
            <h2>{url}</h2>
            <div className={style.entries} key={"directory"}>
                <div className={style.title}>
                    <p className={style.lock}></p>
                    <p className={style.name}>Name</p>
                    <p className={style.size}>Size</p>
                    <p className={style.modified}>Modified</p>
                </div>
                {data.map((obj, index) => {
                    return render(entry({
                        raw: obj.url,
                        name: obj.name,
                        size: obj.size,
                        type: obj.type,
                        modified: obj.modified
                    }), false, index);
                })}
            </div>
        </>
    );
}
