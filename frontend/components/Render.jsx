import Link from "next/link";
import style from "@/styles/Render.module.scss";

const entry = ({ name, size, raw, modified, type }) => {
    return {
        raw: raw,
        name: name,
        size: size,
        type: type,
        modified: modified
    };
};

export function render(obj, key) {
    const backend = `http://localhost:${process.env.SERVER_PORT}`;
    
    let symbol = "";
    let name = obj.name
    let url = `${backend}/f/${obj.raw}`;
    if (obj.type === "dir") {
        url = `${obj.raw}`;
        name = `${obj.name}/`
        symbol = ` bi bi-folder-fill`;
    }

    return (
        <div className={style.entry} href={url} key={key}>
            <i className={`${style.dir}${symbol}`}></i>
            <p className={style.name}>
                <Link href={url} className={style.name_item}>{name}</Link>
            </p>
            <p className={style.size}>{obj.size}</p>
            <p className={style.modified}>{obj.modified}</p>
        </div>
    );
}

export function Render({ data, url }) {
    let ref = "/";
    let back = <></>;

    if (url !== "") {
        ref = url;
        back = (
            <div className={style.entry}>
                <i className={style.dir}></i>
                <p className={style.name}>
                    <Link href={`${url}/../`} className={style.name_item}>../</Link>
                </p>
                <p className={style.size}>-</p>
                <p className={style.modified}>-</p>
            </div>
        );
    }

    return (
        <>
            <h2>{ref}</h2>
            <div className={style.entries} key={"directory"}>
                <div className={style.title}>
                    <p className={style.dir}></p>
                    <p className={style.name}>Name</p>
                    <p className={style.size}>Size</p>
                    <p className={style.modified}>Modified</p>
                </div>
                {back}
                {data.map((obj, index) => {
                    return render(entry({
                        raw: obj.url,
                        name: obj.name,
                        size: obj.size,
                        type: obj.type,
                        modified: obj.modified
                    }), index);
                })}
            </div>
        </>
    );
}
