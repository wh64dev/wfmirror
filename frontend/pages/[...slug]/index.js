import Head from "next/head";
import { getData } from "@/util/route";
import localFont from "next/font/local";
import { Render } from "@/components/Render";
import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";
import styles from "@/styles/Dir.module.scss";

const pretendard = localFont({
    src: "../fonts/Pretendard.woff2",
    variable: "--font-default",
    weight: "300",
});

export default function Slug({ name, data }) {
    return (
        <>
            <Head>
                <title>Create Next App</title>
                <meta name="description" content="Generated by create next app" />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <div className={`${styles.page} ${pretendard.variable}`}>
                <Header name={name} />
                <main className={styles.main}>
                    <Render data={data.data} url={data.dir} />
                </main>
                <Footer />
            </div>
        </>
    );
}

export async function getServerSideProps(context) {
    const name = process.env.FRONT_TITLE;
    const data = await getData(context.resolvedUrl);

    if (data.status === 401) {
        return {
            redirect: {
                destination: "login",
                permanent: true
            }
        };
    }

    if (data.status === 404) {
        return {
            notFound: true
        };
    }

    return {
        props: { name, data }
    };
}
