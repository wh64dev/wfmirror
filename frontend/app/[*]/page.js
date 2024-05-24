import { Render } from "@/components/Render";
import { getData } from "@/util/route";
import { headers } from "next/headers";

export default async function Page() {
    const header = headers();
    console.log(header.get("x-pathname"));
    const data = await getData(header.get("x-pathname"));

    return (
        <main>
            <Render data={data.data} url={data.dir} />
        </main>
    );
}
