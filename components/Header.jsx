import { FileSize } from "./FileSize";
import styles from "./system.module.scss";

const PROJECT_GITHUB = "https://github.com/wh64dev/wfmirror.git";
const GLOBAL_SITE = "https://wh64.net";

export async function Header() {
	return (
		<nav className={styles.header}>
			<h1>{process.env.SERVICE_NAME}</h1>
			<div className={styles.right}>
				<div className={styles.icons}>
					<Icon url={PROJECT_GITHUB} className="bi bi-github" />
					<Icon url={PROJECT_GITHUB} className="bi bi-globe-asia-australia" />
				</div>
				<FileSize />
			</div>
		</nav>
	);
}	

async function Icon({ className, url }) {
	return (
		<a className={styles.icon} href={url}>
			<i className={className} />
		</a>
	);
}
