export const config = {
    api: {
        bodyParser: {
            externalResolver: true
        }
    }
};

/**
 * @param { import("next").NextApiRequest } req
 * @param { import("next").NextApiResponse } res
 */
export default async function handler(req, res) {
    if (req.method !== "POST") {
        res.setHeader("Allow", ["POST"]);
        
        res.statusCode = 405;
        return res.end(`Method ${req.method} Not Allowed`);
    }

    const { username, password } = req.body;
    const form = new URLSearchParams();

    form.append("username", username);
    form.append("password", password);

    const resp = await fetch(`http://localhost:${process.env.SERVER_PORT}/auth/login`, {
        method: "post",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: form
    });

    const obj = await resp.json();
    if (obj.status !== 200) {
        res.statusCode = 401;
        return res.end("credential information not matches");
    }

    const token = obj.token;
    res.statusCode = 200;
    res.setHeader("Authorization", token);

    res.redirect(307, "/");
}
