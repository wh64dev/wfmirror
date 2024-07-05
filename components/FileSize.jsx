import styles from "./filesize.module.scss";

export async function FileSize() {
	const res = await fetch(`${process.env.SERVER_URL}/nodeinfo`);
	const info = await res.json();

	return (
		<div className={styles.indicator}>
			<div className={styles.info}>
				<i className="bi bi-database-check"></i>
				<div className={styles.stats}>
					<div className={styles.stats_row}>
						<p className={styles.stats_name}>Used:</p>
						<p className={styles.stats_value}>{info.drive.used}</p>
					</div>
					<div className={styles.stats_row}>
						<p className={styles.stats_name}>Total:</p>
						<p className={styles.stats_value}>{info.drive.total}</p>
					</div>
				</div>
			</div>
			<progress
				value={parseInt(info.drive.used.split(" ")[0])}
				max={parseInt(info.drive.total.split(" ")[0])}
			/>
		</div>
	);
}
