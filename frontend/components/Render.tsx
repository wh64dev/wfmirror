import Link from "next/link";
import style from "./render.module.scss";
import {notFound, permanentRedirect} from "next/navigation";

type RenderProps = {
    url: string,
    name: string,
    size: string,
    modified: string,
    dir: boolean
};

async function FileEntry({ root, props }: { root: string, props: RenderProps }) {
    let symbol = " bi bi-folder-fill";
    let name = `${props.name}/`;
    let url = props.url;

    if (!props.dir) {
        name = `${props.name}`;

        url = `/f/${props.url}`;
        if (root !== "") {
            url = `/f${props.url}`;
        }

        symbol = "";
    }

    return (
        <>
            <i className={`${style.dir}${symbol}`}></i>
            <p className={style.name}>
                <Link href={url} className={style.name_item}>{name}</Link>
            </p>
            <p className={style.size}>{props.size}</p>
            <p className={style.modified}>{props.modified}</p>
        </>
    );
}

export async function Render({ url, token }: { url: string, token: string | undefined }) {
    async function get() {
        "use server";

        const res = await fetch(`http://localhost:${process.env.SERVER_PORT}/path${url}`, {
            method: "GET",
            headers: {
                Authorization: token ? `Bearer ${token}` : ""
            }
        });

        if (res.status !== 200) {
            if (res.status === 401) {
                permanentRedirect("/login");
            }

            notFound();
        }

        return await res.json();
    }

    const entries = await get();
    let back = <></>;
    if (url !== "") {
        back = (
            <div className={style.entry}>
                <FileEntry
                    root={url}
                    props={{
                        url: `${url}/../`,
                        name: "..",
                        size: "-",
                        modified: "-",
                        dir: true
                    }}
                />
            </div>
        );
    }

    let dir = entries.dir;
    if (dir === "") {
        dir = "/";
    }

    return (
        <>
        <h2>{dir}</h2>
            <div className={style.entries}>
                <div className={style.title}>
                    <p className={style.dir}></p>
                    <p className={style.name}>Name</p>
                    <p className={style.size}>Size</p>
                    <p className={style.modified}>Modified</p>
                </div>
                {back}
                {entries.data.map((entry: any, index: number) => {
                    return (
                        <div className={style.entry} key={index}>
                            <FileEntry
                                root={url}
                                props={{
                                    url: entry.url,
                                    name: entry.name,
                                    size: entry.size,
                                    modified: entry.modified,
                                    dir: entry.type === "dir"
                                }}
                            />
                        </div>
                    );
                })}
            </div>
        </>
    );
}
