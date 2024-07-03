"use client";

import styles from "./render.module.scss";

/**
 * @author WH64
 * @param { string } url
 * @param { any } data
 * @returns { JSX.Element }
 */
export function Render({ url, data }) {
	return (
		<>
			<h2>{data.dir !== "" ? data.dir : "/"}</h2>
			{data.data.map((file, index) => {
				if (file.type === "dir") {
					return (
						<a key={index} href={`/${file.url}`}>
							{file.name}/
						</a>
					);
				}

				return (
					<a key={index} href={`${url}/path/${file.url}`}>
						{file.name}
					</a>
				);
			})}
		</>
	);
}
