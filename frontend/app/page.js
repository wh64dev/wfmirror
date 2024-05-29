import { Footer } from "@/components/Footer";
import { Header } from "@/components/Header";
import { Render } from "@/components/Render";
import { getData } from "@/util/route";

export default async function Route() {
    const data = await getData("/");

    return (
        <main>
            <Header />
            <Render data={data.data} url={data.dir} />
            <Footer />
        </main>
    );
}
