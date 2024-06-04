import Link from "next/link";
import style from "@/styles/Render.module.scss";

const entry = ({ name, size, url, modified, type }) => {
    return {
        url: url,
        name: name,
        size: size,
        type: type,
        modified: modified
    };
};

export function render(obj, key, port) {
    const backend = `http://localhost:${port}`;
    
    let symbol = "";
    let name = obj.name
    let url = `${backend}/f${obj.url}`;
    if (obj.type === "dir") {
        url = `${obj.url}`;
        name = `${obj.name}/`
        symbol = ` bi bi-folder-fill`;
    }

    return (
        <div className={style.entry} key={key}>
            <i className={`${style.dir}${symbol}`}></i>
            <p className={style.name}>
                <Link href={url} className={style.name_item}>{name}</Link>
            </p>
            <p className={style.size}>{obj.size}</p>
            <p className={style.modified}>{obj.modified}</p>
        </div>
    );
}

export function Render({ data, url, port }) {
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
                        url: obj.url,
                        name: obj.name,
                        size: obj.size,
                        type: obj.type,
                        modified: obj.modified
                    }), index, port);
                })}
            </div>
        </>
    );
}
