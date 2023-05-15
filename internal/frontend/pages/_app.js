import Head from "next/head";
import Link from "next/link";
import React from "react";
import { createClient, fetchExchange, Provider } from "urql";
import "./globals.css";

const client = createClient({
  url: `http://localhost:8383/query`,
  exchanges: [fetchExchange],
});

function MyApp({ Component, pageProps }) {
  return (
    <div>
      <Head></Head>
      <div className="w-full bg-gray-900 text-center py-5 shadow">
        <img src="/gopher.png" className="h-36 mx-auto" />
        <h1 className="text-2xl text-white w-max mx-auto mt-4">Go Workshop</h1>
      </div>
      <div className="w-full bg-gray-800 text-center py-5 shadow flex flex-row gap-3 justify-center text-white">
        <Link
          className="hover:text-gray-400 transition-all duration-200"
          href="/"
        >
          Home
        </Link>
        <Link
          className="hover:text-gray-400 transition-all duration-200"
          href="/level1"
        >
          Level 1
        </Link>
        <Link
          className="hover:text-gray-400 transition-all duration-200"
          href="/level2"
        >
          Level 2
        </Link>
        <Link
          className="hover:text-gray-400 transition-all duration-200"
          href="/level3"
        >
          Level 3
        </Link>
        <Link
          className="hover:text-gray-400 transition-all duration-200"
          href=""
        >
          Level 4
        </Link>
      </div>
      <div className="w-4/5 mx-auto my-10">
        <Provider value={client}>
          <Component {...pageProps} />
        </Provider>
      </div>
    </div>
  );
}

export default MyApp;
