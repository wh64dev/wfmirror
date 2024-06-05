import { Render } from "@/components/Render";
import style from "./page.module.scss";
import {Header} from "@/components/Header";
import {Footer} from "@/components/Footer";

export default function Root() {
    return (
        <main className={style.main}>
            <Header />
            <Render url={""} token={""} />
            <Footer />
        </main>
    );
}