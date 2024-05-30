import { Footer } from "@/components/Footer";
import { Header } from "@/components/Header";
import { Render } from "@/components/Render";
import { getData } from "@/util/route";
import { headers } from "next/headers";

export default async function Page() {
    const header = headers();
    const data = await getData(header.get("x-current-path"));

    return (
        <main>
            <Header />
            <Render data={data.data} url={data.dir} />
            <Footer />
        </main>
    );
}
