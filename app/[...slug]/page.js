import { Render } from "@/components/Render";
import { RenderContent } from "@/components/RenderContent";

export default async function Path({ params }) {
	async function read() {
		"use server";

		let path = "";
		for (const p of params.slug) {
			path += `/${p}`;
		}

		const res = await fetch(`${process.env.SERVER_URL}/path${path}`, {
			mode: "cors",
			method: "GET",
			cache: "no-cache"
		});

		try {
			return await res.json();
		} catch (err) {
			// IGNORE
		}

		return res;
	}

	const data = await read();
	if (data instanceof Response) {
		return <RenderContent obj={data} />
	}

	return <Render url={process.env.SERVER_URL} data={data} back={true} />
}
