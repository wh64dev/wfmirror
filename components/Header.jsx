import styles from "./system.module.scss";

export function Header() {
	return (
		<nav className={styles.header}>
			<h1>{process.env.SERVICE_NAME}</h1>
		</nav>
	);
}
