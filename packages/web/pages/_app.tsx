import { AppProps } from "next/app";
import Head from "next/head";
import { useState } from "react";
import { Hydrate, QueryClient, QueryClientProvider } from "react-query";
import "./styles.scss";

function App({ Component, pageProps }: AppProps) {
  const [queryClient] = useState(() => new QueryClient());

  return (
    <>
      <QueryClientProvider client={queryClient}>
        <Hydrate state={pageProps.dehydratedState}>
          <Head>
            <title>Gimi</title>
          </Head>
          <main className="app">
            <Component {...pageProps} />
          </main>
        </Hydrate>
      </QueryClientProvider>
    </>
  );
}

export default App;
