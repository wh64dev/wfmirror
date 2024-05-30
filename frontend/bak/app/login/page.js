import { Footer } from "@/components/Footer";
import style from "./login.module.scss";

export default async function Login() {
    const action = async (formData) => {
        "use server";

        const res = await fetch(`http://localhost:${process.env.SERVER_PORT}/auth/login`, {
            method: "POST",
            "body": formData
        });

        if (res.ok) {
            const data = await res.json();
            console.log(data);
        }
    }

    return (
        <main>
            <div className={style.container}>
                <form
                    className={style.form}
                    action={action}
                >
                    <input
                        name="username"
                        className={style.input}
                        placeholder="Username"
                        type="text"
                        required
                    />
                    <input
                        name="password"
                        className={style.input}
                        placeholder="Password"
                        type="password"
                        required
                    />

                    <input
                        className={style.submit_btn}
                        value={"Login"}
                        type="submit"
                    />
                </form>
            </div>
            <Footer />
        </main>
    );
}
