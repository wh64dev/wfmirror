import { Render } from "@/components/Render";

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

		return res.json();
	}

	const data = await read();

	return <Render url={process.env.SERVER_URL} data={data} back={true} />
}
