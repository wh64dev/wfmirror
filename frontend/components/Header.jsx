import style from "./header.module.scss";

export function Header() {
    return (
        <nav className={style.nav}>
            <div>
                <h1>{process.env.FRONT_TITLE}</h1>
            </div>
        </nav>
    );
}
