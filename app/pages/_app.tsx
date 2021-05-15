import '../styles/globals.css'
import Head from 'next/head'
function MyApp({ Component, pageProps }) {
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
  <Component {...pageProps} />
  
  </>
  )
}

export default MyApp
