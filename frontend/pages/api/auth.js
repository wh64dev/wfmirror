import { NextRequest, NextResponse } from "next/server";
import { IncomingForm } from "formidable";

/**
 * @param {NextRequest} req
 * @param {NextResponse} res
 */
export default async function handler(req, res) {
    if (req.method !== "POST") {
        res.setHeader("Allow", ["POST"]);
        res.status(405).end(`Method ${req.method} Not Allowed`);
        return;
    }

    const formData = new IncomingForm();
    const resp = await fetch(`http://localhost:${process.env.SERVER_PORT}/auth/login`, {
        method: "post",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: new URLSearchParams(formData)
    });

    const obj = await resp.json();
    if (obj.status !== 200) {
        res.status(401);
        return;
    }

    console.log(obj);
    res.headers.append("Authorization", obj.token);
    res.status(200);
}
