import style from "../page.module.scss";
import {Render} from "@/components/Render";
import {cookies, headers} from "next/headers";
import {Header} from "@/components/Header";
import {Footer} from "@/components/Footer";

export default async function Path() {
    const header = headers()
    const cookie = cookies();

    const path = header.get("x-current-path")!;
    const token = cookie.get("token")?.value;

    return (
        <main className={style.main}>
            <Header />
            <Render url={path} token={token} />
            <Footer />
        </main>
    );

}
