import Link from "next/link";
import {notFound, permanentRedirect} from "next/navigation";

type RenderProps = {
    url: string,
    name: string,
    size: string,
    dir: boolean
};

async function getPort() {
    "use server";
    return process.env.SERVER_PORT!;
}

async function FileEntry({ root, key, props }: { root: string, key: number | undefined, props: RenderProps }) {
    let name = `${props.name}/`;
    let url = props.url;

    if (!props.dir) {
        name = `${props.name}`;

        url = `http://localhost:${await getPort()}/f/${props.url}`;
        if (root !== "") {
            url = `http://localhost:${await getPort()}/f${props.url}`;
        }
    }

    return (
        <div key={key}>
            <Link href={url}>{name}</Link>
        </div>
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
        back = <FileEntry
            root={url}
            key={undefined}
            props={{
                url: `${url}/../`,
                name: "..",
                size: "-",
                dir: true
            }}
        />
    }

    return (
        <div>
            {back}
            {entries.data.map((entry: any, index: number) => {
                return (
                    <FileEntry
                        root={url}
                        key={index}
                        props={{
                            url: entry.url,
                            name: entry.name,
                            size: entry.size,
                            dir: entry.type === "dir"
                        }}
                    />
                );
            })}
        </div>
    );
}
