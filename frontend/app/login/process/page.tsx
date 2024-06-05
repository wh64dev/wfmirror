"use client";

import {redirect, useSearchParams} from "next/navigation";

export default function Process() {
    const params = useSearchParams();
    const success = params.get("success");
    if (success === null || success === "1") {
        return redirect("/");
    }

    alert("username or password not matches!");
    redirect("/login");
}
