import { Render } from "@/components/Render";

export default async function Home() {
	async function read() {
		"use server";

		const res = await fetch(`${process.env.SERVER_URL}/path`, {
			mode: "cors",
			method: "GET",
			cache: "no-cache"
		});

		return res.json();
	}

	const data = await read();
	return <Render url={process.env.SERVER_URL} data={data} />
}
