import { useRouter } from 'next/router'
import {DocsThemeConfig, ThemeSwitch} from 'nextra-theme-docs'
import { useConfig } from 'nextra-theme-docs'
import Link from "next/link";

const config: DocsThemeConfig = {
  project: {
    link: 'https://github.com/plantoncloud/project-planton'
  },
  docsRepositoryBase: 'https://github.com/plantoncloud/project-planton/tree/main/docs',
    logo: (
        <div className="flex flex-row items-center">
            <img src="/images/logo/full-logo-transparent.png" alt="ProjectPlanton Logo" width="300" />
        </div>
    ),
    nextThemes: {
      defaultTheme: 'light'
    },
    head: function useHead() {
        const config = useConfig()
        const {route} = useRouter()
        const isDefault = route === '/' || !config.title
        const image =
            'https://nextra.site/' +
            (isDefault ? 'og.jpeg' : `api/og?title=${config.title}`)

        const description =
            config.frontMatter.description ||
      'Deploy Apps, OpenSource & Cloud Infra'
    const title = config.title + (route === '/' ? '' : ' - Nextra')

    return (
      <>
        <title>{title}</title>
        <meta property="og:title" content={title} />
        <meta name="description" content={description} />
        <meta property="og:description" content={description} />
        <meta property="og:image" content={image} />

        <meta name="msapplication-TileColor" content="#fff" />
        <meta httpEquiv="Content-Language" content="en" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:site:domain" content="nextra.site" />
        <meta name="twitter:url" content="https://nextra.site" />
        <meta name="apple-mobile-web-app-title" content="Nextra" />
        <link rel="icon" href="/images/logo/favicon.ico" type="image/svg+xml" />
        <link rel="icon" href="/images/logo/favicon.ico" type="image/png" />
        <link
          rel="icon"
          href="/images/logo/favicon.ico"
          type="image/svg+xml"
          media="(prefers-color-scheme: dark)"
        />
        <link
          rel="icon"
          href="/images/logo/favicon.ico"
          type="image/svg+xml"
          media="(prefers-color-scheme: dark)"
        />
      </>
    )
  },
  editLink: {
    content: 'Edit this page on GitHub →'
  },
  feedback: {
    content: 'Question? Give us feedback →',
    labels: 'feedback'
  },
  sidebar: {
    defaultMenuCollapseLevel: 1,
    toggleButton: true
  },
  footer: {
    content: (
      <div className="flex w-full flex-col items-center sm:items-start">
        <p className="mt-4 text-xs">
          © {new Date().getFullYear()} ProjectPlanton.
        </p>
      </div>
    )
  }
}

export default config
