import styles from "./system.module.scss";

export function Footer() {
	return (
		<footer className={styles.footer}>
			<p>
				&copy; Copyright 2024 <a href="https://github.com/wh64dev">WH64</a>. All Rights Reserved. Powered by <a href="https://wh64.net">WSERVER</a>.
			</p>
		</footer>
	);
}
