import '../styles/globals.css'
import Head from 'next/head'
import type { AppProps } from 'next/app'
import cookie from "cookie";
import axios from 'axios'
import Layout from 'components/Layout';
import ContextWrapper from 'components/ContextWrapper'
import Header from 'components/Header';
import { useEffect, useState } from 'react';


function MyApp({ Component, pageProps, router }: AppProps) {
  
  const {NEXT_PUBLIC_GOOGLE_ANALYTICS} = process.env
 

  return( 
  <>
    <Head>
    <title>موجز | ajf.sa</title>
    <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    <meta charSet="UTF-8" />
    <meta name="description" content="موجز اخبار,موجز اعمال, موجز رياضة,موجز فن,موجز ترفيه, موجز سياحة, موجز تقنية" />
    <meta name="keywords" content="موجز اخبار,موجز اعمال, موجز رياضة,موجز فن,موجز ترفيه, موجز سياحة, موجز تقنية" />
    <script
    async
    src={`https://www.googletagmanager.com/gtag/js?id=${NEXT_PUBLIC_GOOGLE_ANALYTICS}`}
  />
  <script
    dangerouslySetInnerHTML={{
      __html: `
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());
    gtag('config', '${NEXT_PUBLIC_GOOGLE_ANALYTICS}', {
      page_path: window.location.pathname,
    });
  `,
    }}
  />
  </Head>

    <Header />

  <Layout>
  <Component {...pageProps} />
  </Layout>
  
  </>
  )
}


// MyApp.getInitialProps = async ({Component, ctx}) => {
//   let pageProps = {}
  
//   const {APP_URL} = process.env
//   const res = await axios(`${APP_URL}/tags`)
//   const tags = await res.data

//   if (Component.getInitialProps) {
//       pageProps = await Component.getInitialProps(ctx)
//   }


//   return {
     
//       tags
//   }
// }

export default MyApp
