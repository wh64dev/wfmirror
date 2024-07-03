import { Render } from "@/components/Render";
import styles from "@/app/page.module.scss";

export default async function Page({ params }) {
	async function read() {
		"use server";
		const res = await fetch(`${process.env.SERVER_URL}/path/${params.slug}`, {
			mode: "cors"
		});

		return res.json();
	}

	const data = await read();
	return <Render url={process.env.SERVER_URL} data={data} />
}
