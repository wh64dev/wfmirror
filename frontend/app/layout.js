import "./globals.scss";
import "bootstrap-icons/font/bootstrap-icons.scss";

export const metadata = {
    title: "WH64's File Mirror",
    description: "Personal file mirror server application",
};

export default function RootLayout({ children }) {
    return (
        <html lang="en">
            <body>{children}</body>
        </html>
    );
}
