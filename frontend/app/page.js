import { Render } from "@/components/Render";
import { getData } from "@/util/route";

export default async function Home() {
    const data = await getData("//");

    return (
        <main>
            <Render data={data.data} url={data.dir} />
        </main>
    );
}
