"use client";

import styles from "./render.module.scss";
import React from "react";

function RenderEntry({ url, data, back }) {
	if (data === null) {
		return <tr className={styles.entry}>
			<td className={styles.file_icon}>
				<i className="bi bi-folder-x" />
			</td>
			<td className={styles.entry_name}>
				<p>Directory Empty</p>
			</td>
			<td className={styles.entry_item}>X</td>
			<td className={styles.entry_item}>X</td>
		</tr>;
	}

	return (
		<>
			{data.map((file, index) => {
				if (file.type === "dir") {
					const resolvedPath = `${location.pathname}/${file.name}`;

					return (
						<tr key={index} className={styles.entry}>
							<td className={styles.file_icon}>
								<i className="bi bi-folder"/>
							</td>
							<td className={styles.entry_name}>
								<a href={back ? resolvedPath : file.url}>
									{file.name}/
								</a>
							</td>
							<td className={styles.entry_item}>-</td>
							<td className={styles.entry_item}>{file.modified}</td>
						</tr>
					);
				}

				return (
					<tr key={index} className={styles.entry}>
						<td className={styles.file_icon}>
							<i className="bi bi-file-earmark"/>
						</td>
						<td className={styles.entry_name}>
							<a href={`${url}/path/${file.url}`}>
								{file.name}
							</a>
						</td>
						<td className={styles.entry_item}>{file.size}</td>
						<td className={styles.entry_item}>{file.modified}</td>
					</tr>
				);
			})}
			</>
	);
}

function click(ev, url) {
	ev.preventDefault();
	location.href = url;
}

/**
 * @author WH64
 * @returns { JSX.Element }
 */
export function Render({ url, data, back = false }) {
	return (
		<div className={styles.container}>
			<div className={styles.explorer}>
				<b>Path: {data.dir !== "" ? data.dir : "/"}</b>
				<button onClick={ev => click(ev, `${url}/path${data.dir}`)}>Raw</button>
			</div>
			<table className={styles.entries}>
				<tbody>
				<tr className={styles.entry_main}>
					<th></th>
					<th className={styles.entry_name}>Name</th>
					<th className={styles.entry_item}>Size</th>
					<th className={styles.entry_item}>Modified</th>
				</tr>
				{
					!back ? <></> : <tr className={styles.entry}>
						<td className={styles.file_icon}></td>
						<td className={styles.entry_name}>
							<a href={`${location.pathname}/../`}>
								../
							</a>
						</td>
						<td className={styles.entry_item}>-</td>
						<td className={styles.entry_item}>-</td>
					</tr>
				}
				<RenderEntry url={url} data={data.data} back={back} />
				</tbody>
			</table>
		</div>
	);
}
