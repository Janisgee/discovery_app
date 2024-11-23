import { Kaisei_Tokumin } from "next/font/google";
import { Space_Mono } from "next/font/google";
import { Inter } from "next/font/google";

export const kaiseiTokumin = Kaisei_Tokumin({
  subsets: ["latin"],
  variable: "--font-kaisei_tokumin",
  weight: ["400", "700"],
});

export const space_mono = Space_Mono({
  subsets: ["latin"],
  variable: "--font-space_mono",
  weight: ["400", "700"],
});

export const inter = Inter({
  subsets: ["latin"],
  variable: "--font-inter",
});
