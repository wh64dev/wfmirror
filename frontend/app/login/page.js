"use client";

import React from "react";
import { login } from "@/app/lib/login";
import style from "./login.module.scss";
import { Footer } from "@/components/Footer";

export default function Login() {
    const [message, action] = React.useActionState(login, null);
    
    return (
        <main>
            <div className={style.container}>
                <form
                    className={style.form}
                    action={action}
                > 
                    <p aria-live="polite" className="sr-only">
                        {message}
                        {console.log(message)}
                    </p>

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
