import { Render } from "@/components/Render";
import { getData } from "@/util/route";
import { headers } from "next/headers";

export default async function Page() {
    const header = headers();
    const path = header.get("x-current-path");
    const data = await getData(path);

    console.log(path);

    return (
        <main>
            <Render data={data.data} url={data.dir} />
        </main>
    );
}
