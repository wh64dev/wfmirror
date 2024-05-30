import style from "@/styles/Header.module.scss";

export function Header({ name }) {
    return (
        <nav className={style.nav}>
            <div>
                <h1>{name}</h1>
            </div>
        </nav>
    )
}
