import { NextRequest, NextResponse } from "next/server";

/**
 * @param {NextRequest} req
 * @param {NextResponse} res
 */
export default async function handler(req, res) {
    if (req.method !== "POST") {
        res.setHeader("Allow", ["POST"]);
        
        res.status = 405;
        return res.send(`Method ${req.method} Not Allowed`);
    }

    console.log(req);
    // const formData = req.body();
    // console.log(formData);

    // const resp = await fetch(`http://localhost:${process.env.SERVER_PORT}/auth/login`, {
    //     method: "post",
    //     headers: {
    //         "Content-Type": "application/x-www-form-urlencoded",
    //     },
    //     body: new URLSearchParams(formData)
    // });

    // const obj = await resp.json();
    // if (obj.status !== 200) {
    //     res.status = 401;
    //     return res.send("credential information not matches");
    // }

    // const token = obj.token;
    // res.status = 200;
    // res.setHeader("Authorization", token);

    // return res.send("login success!").redirect("/");
}
