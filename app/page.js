import { Render } from "@/components/Render";

export default async function Home() {
	async function read() {
		"use server";
		const res = await fetch(`http://localhost:${process.env.SERVICE_PORT}/path/`, {
			mode: "cors",
			method: "GET"
		});

		return res.json();
	}

	const data = await read();
	return <Render url={process.env.SERVER_URL} data={data} />
}
