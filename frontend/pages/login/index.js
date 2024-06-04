import style from "@/styles/Login.module.scss";
import { createRef } from "react";

export default function Login({ title }) {
    const username = createRef();
    const password = createRef();
    // const err = createRef();

    // const login = async (ev) => {
    //     ev.preventDefault();

    //     const obj = {
    //         username: username.current.value,
    //         password: password.current.value
    //     };
        
    //     const res = await fetch("/api/auth", {
    //         method: "post",
    //         body: JSON.stringify(obj)
    //     });

    //     username.current.value = "";
    //     password.current.value = "";

    //     if (res.status !== 200) {
    //         err.current.innerText = "username or password not matches";
    //         return;
    //     }

    //     err.current.innerText = "";

    //     console.log(res);
    //     // window.location.href = "/";
    // };

    return (
        <main className={style.main}>
            <form className={style.form} action={"/api/auth"} method="POST">
                <h1 className={style.title}>{title} Login</h1>
                {/* <p className={style.err_msg} ref={err}></p> */}
                <input className={style.input} type="text" placeholder="Username" ref={username} required />
                <input className={style.input} type="password" placeholder="Password" ref={password} required />
                
                <input className={style.submit} type="submit" value={"Login"} />
            </form>
        </main>
    );
}

export async function getServerSideProps(context) {
    const title = process.env.FRONT_TITLE;
    if (context.res.getHeader("Authorization") !== undefined) {
        return {
            redirect: {
                distination: "/",
                permanent: true
            }
        };
    }

    return {
        props: { title }
    };
}
