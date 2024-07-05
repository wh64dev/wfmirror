import { FileSize } from "./FileSize";
import styles from "./system.module.scss";

export async function Header() {
	return (
		<nav className={styles.header}>
			<h1>{process.env.SERVICE_NAME}</h1>
			<div className={styles.right}>
				<FileSize />
			</div>
		</nav>
	);
}
