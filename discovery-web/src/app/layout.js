import "@/app/_ui/globals.css";

import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";

import { kaiseiTokumin, space_mono, inter } from "@/app/_ui/fonts/fonts.js";

config.autoAddCss = false;

export const metadata = {
  title: {
    default: "Discovery App",
    template: "%s | Discovery App",
  },
  description: "Generated travel ideas for daily life",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body
        className={`${kaiseiTokumin.variable} ${space_mono.variable} ${inter.variable}`}
      >
        {children}
      </body>
    </html>
  );
}
