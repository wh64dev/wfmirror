import { Render } from "@/components/Render";
import {cookies, headers} from "next/headers";

export default async function Path() {
    const header = headers()
    const cookie = cookies();

    const path = header.get("x-current-path")!;
    const token = cookie.get("token")?.value;

    return (
        <main>
            <Render url={path} token={token} />
        </main>
    );

}
