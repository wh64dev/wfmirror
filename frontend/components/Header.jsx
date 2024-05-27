import style from "./header.module.scss";

const name = process.env.FRONT_TITLE;

export function Header() {
    return (
        <nav className={style.nav}>
            <div>
                <h1>{name}</h1>
            </div>
        </nav>
    )
}
