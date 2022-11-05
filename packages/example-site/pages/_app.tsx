import { AppProps } from "next/app";
import Head from "next/head";
import Script from "next/script";
import "./styles.css";

function CustomApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <title>Welcome to example-site!</title>
      </Head>
      <Script src="/js/lib.js"></Script>
      <main className="app">
        <Component {...pageProps} />
      </main>
    </>
  );
}

export default CustomApp;
