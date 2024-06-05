import {cookies} from "next/headers";

export async function GET(req: Request, { params }: { params: { slug: string } }) {
    const cookie = cookies();
    console.log(params.slug);

    const token = cookie.get("token");
    if (token === undefined || token?.value === "") {
        return Response.redirect("/login");
    }

    let path = "";
    for (const p of params.slug) {
        path += `/${p}`;
    }

    const res = await fetch(`http://localhost:${process.env.SERVER_PORT}/f${path}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${cookie.get("token")?.value}`
        }
    });

    return new Response(res.body, {
        status: 200,
        headers: {
            "Content-Disposition": "attachment"
        }
    })
}