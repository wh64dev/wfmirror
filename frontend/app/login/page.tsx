import {cookies} from "next/headers";
import {permanentRedirect, RedirectType} from "next/navigation";

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

    permanentRedirect("/", RedirectType.push);
}

export default function Login() {
    return (
        <main>
            <form method={"POST"} action={login}>
                <h2>{process.env.FRONT_TITLE}</h2>
                <div>
                    <input name={"username"} type={"text"} placeholder={"Username"} required/>
                    <input name={"password"} type={"password"} placeholder={"Password"} required/>

                    <button type={"submit"}>Login</button>
                </div>
            </form>
        </main>
    );
}
