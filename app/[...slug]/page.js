import { Render } from "@/components/Render";

export default async function Path({ params }) {
	async function read() {
		"use server";

		let path = "";
		for (const p of params.slug) {
			path += `/${p}`
		}

		console.log(path);

		const res = await fetch(`http://localhost:${process.env.SERVICE_PORT}/path${path}`, {
			mode: "cors",
			method: "GET"
		});

		return res.json();
	}

	const data = await read();

	return <Render url={`http://localhost:${process.env.SERVICE_PORT}`} data={data} back={true} />
}
