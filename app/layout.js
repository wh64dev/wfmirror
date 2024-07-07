import "./globals.scss";
import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";
import "bootstrap-icons/font/bootstrap-icons.scss";

export const metadata = {
	title: process.env.SERVICE_NAME,
	description: "Simple personal file mirror service"
};

export default function RootLayout({ children }) {
	return (
		<html lang="en">
			<body>
				<Header />
				{children}
				<Footer />
			</body>
		</html>
	);
}
