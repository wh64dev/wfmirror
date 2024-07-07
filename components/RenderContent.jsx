import styles from "./render.module.scss";

export async function RenderContent({ obj }) {
	const split = obj.url.split("/");
	const index = split.length - 1;
	const filename = split[index];
	console.log(obj);

	return (
		<div>
			<h5>{filename}</h5>
		</div>
	);
}
