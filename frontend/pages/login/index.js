import style from "@/styles/Login.module.scss";
import { createRef } from "react";

export default function Login({ title }) {
    const username = createRef();
    const password = createRef();
    const err = createRef();

    const login = async (ev) => {
        ev.preventDefault();
        const formData = new FormData();
        formData.append("username", username.current.value);
        formData.append("password", password.current.value);

        const res = await fetch("/api/auth", {
            method: "post",
            body: formData
        });

        if (res.status !== 200) {
            err.current.value = "username or password not matches";
            return;
        }

        console.log(res);
    };

    return (
        <main className={style.main}>
            <form className={style.form} onSubmit={login}>
                <h1 className={style.title}>{title} Login</h1>
                <p ref={err}></p>
                <input className={style.input} type="text" placeholder="Username" ref={username} required />
                <input className={style.input} type="password" placeholder="Password" ref={password} required />
                
                <input className={style.submit} type="submit" value={"Login"} />
            </form>
        </main>
    );
}

export async function getServerSideProps() {
    const title = process.env.FRONT_TITLE;

    return {
        props: { title }
    };
}
