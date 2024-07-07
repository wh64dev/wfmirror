import { Render } from "@/components/Render";
import { permanentRedirect } from "next/navigation";

export default async function Path({ params }) {
	async function read() {
		"use server";

		let path = "";
		for (const p of params.slug) {
			path += `/${p}`;
		}

		const url = `${process.env.SERVER_URL}/path${path}`;
		const res = await fetch(url, {
			mode: "cors",
			method: "GET",
			cache: "no-cache"
		});

		if (!res.headers.get("Content-Type").startsWith("application/json")) {
			return permanentRedirect(url);
		}

		return await res.json();
	}

	const data = await read();
	return <Render url={process.env.SERVER_URL} data={data} back={true} />
}
