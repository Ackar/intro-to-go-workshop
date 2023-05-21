import { UniqueDirectivesPerLocationRule } from "graphql";
import Head from "next/head";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { createClient, fetchExchange, Provider } from "urql";
import "./globals.css";

const client = createClient({
  url: `https://workshop.sycl.dev/query`,
  exchanges: [fetchExchange],
});

function MyApp({ Component, pageProps }) {
  return (
    <div>
      <Head>
        <title>Go workshop</title>
      </Head>
      <div className="w-full bg-gray-900 text-center py-5 shadow">
        <img src="/gopher.png" className="h-36 mx-auto" />
        <h1 className="text-2xl text-white w-max mx-auto mt-4">Go Workshop</h1>
      </div>
      <div className="w-full bg-gray-800 text-center py-2 shadow flex flex-row gap-3 justify-center text-white">
        <MenuItem name="Home" href="/" />
        <MenuItem name="Level 1" href="/level1" />
        <MenuItem name="Level 2" href="/level2" />
        <MenuItem name="Level 3" href="/level3" />
      </div>
      <div className="w-4/5 mx-auto my-10">
        <Provider value={client}>
          <Component {...pageProps} />
        </Provider>
      </div>
    </div>
  );
}

function MenuItem({ href, name }) {
  const router = useRouter();

  return (
    <Link
      className={`hover:text-gray-400 transition-all duration-200 px-2 py-1 rounded ${
        router.pathname === href ? "bg-gray-900" : ""
      }`}
      href={href}
    >
      {name}
    </Link>
  );
}

export default MyApp;
