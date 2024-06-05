import {cookies} from "next/headers";
import {permanentRedirect, RedirectType} from "next/navigation";
import style from "./login.module.scss";

async function login(formData: FormData) {
    "use server";

    const username = formData.get("username")?.toString()!;
    const password = formData.get("password")?.toString()!;

    const form = new URLSearchParams();
    form.append("username", username);
    form.append("password", password);

    const res = await fetch(`http://localhost:${process.env.SERVER_PORT}/auth/login`, {
        method: "POST",
        body: form
    });

    if (res.status !== 200) {
        if (res.status === 401) {
            permanentRedirect("/login/process?success=0", RedirectType.push);
        }

        return;
    }

    const obj = await res.json();
    cookies().set({
        name: "token",
        value: obj.token,
        httpOnly: true,
        sameSite: "strict"
    });

    permanentRedirect("/login/process?success=1", RedirectType.push);
}

export default function Login() {
    return (
        <main className={style.main}>
            <form className={style.form} method={"POST"} action={login}>
                <h1 className={style.title}>{process.env.FRONT_TITLE} Login</h1>

                <input className={style.input} name={"username"} type={"text"} placeholder={"Username"} required/>
                <input className={style.input} name={"password"} type={"password"} placeholder={"Password"} required/>

                <button className={style.submit} type={"submit"}>Login</button>
            </form>
        </main>
    );
}
